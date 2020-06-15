package common

import (
	"context"
	"math"
	"sync"
)

const DefaultMaxFetchersCount = 1
const DefaultMaxRequestsCountPerSecond = 0

type ListingSettings struct {
	MaxRequestsCountPerSecond int
	StreamBufferLength        int
	MaxFetchersCount          int
	MaxItemsPerRequest        int
}

type Cursor struct {
	Limit  int
	Offset int
}

type ItemsStream chan Item

type Item struct {
	Err           error
	TotalCount    int
	Payload       interface{}
}

func setListingSettingsDefaults(settingsFromInput ListingSettings) ListingSettings {
	if settingsFromInput.MaxRequestsCountPerSecond == 0 {
		settingsFromInput.MaxRequestsCountPerSecond = DefaultMaxRequestsCountPerSecond
	}

	if settingsFromInput.MaxItemsPerRequest == 0 || settingsFromInput.MaxItemsPerRequest > MaxCountPerBulkRequestItem * MaxCountPerBulkRequestItem {
		settingsFromInput.MaxItemsPerRequest = MaxCountPerBulkRequestItem * MaxCountPerBulkRequestItem
	}

	if settingsFromInput.MaxFetchersCount == 0 {
		settingsFromInput.MaxFetchersCount = DefaultMaxFetchersCount
	}

	return settingsFromInput
}

type DataProvider interface {
	Count(ctx context.Context, filters map[string]interface{}) (int, error)
	Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error
}

type Lister struct {
	listingSettings     ListingSettings
	reqThrottler        Throttler
	listingDataProvider DataProvider
}

func NewLister(settings ListingSettings, dataProvider DataProvider, sl Sleeper) *Lister {
	settings = setListingSettingsDefaults(settings)

	thrl := NewSleepThrottler(settings.MaxRequestsCountPerSecond, sl)

	return &Lister{
		listingSettings:     settings,
		reqThrottler:        thrl,
		listingDataProvider: dataProvider,
	}
}

func (p *Lister) Get(ctx context.Context, filters map[string]interface{}) ItemsStream {
	p.reqThrottler.Throttle()

	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	totalCount, err := p.listingDataProvider.Count(ctx, filters)
	if err != nil {
		outputChan := make(ItemsStream, 1)
		defer close(outputChan)

		outputChan <- Item{
			Err:        err,
			TotalCount: totalCount,
			Payload:    nil,
		}
		return outputChan
	}

	cursorsChan := p.getCursors(ctx, totalCount)

	childChans := make([]ItemsStream, 0, p.listingSettings.MaxFetchersCount)
	for i := 0; i < p.listingSettings.MaxFetchersCount; i++ {
		childChan := p.fetchProductsChunk(ctx, cursorsChan, totalCount, filters)
		childChans = append(childChans, childChan)
	}

	return p.mergeChannels(ctx, childChans...)
}

func (p *Lister) fetchProductsChunk(ctx context.Context, cursorChan chan []Cursor, totalCount int, filters map[string]interface{}) ItemsStream {
	prodStream := make(chan Item, p.listingSettings.StreamBufferLength)
	go func() {
		defer close(prodStream)
		for cursors := range cursorChan {
			p.fetchProductsFromAPI(ctx, cursors, totalCount, prodStream, filters)

			select {
			case <-ctx.Done():
				return
			default:
				continue
			}
		}
	}()

	return prodStream
}

func (p *Lister) getCursors(ctx context.Context, totalCount int) chan []Cursor {
	out := make(chan []Cursor, p.listingSettings.MaxFetchersCount)

	leftCount := totalCount

	go func() {
		defer close(out)

		curPage := 1
		if p.listingSettings.MaxItemsPerRequest > MaxCountPerBulkRequestItem*MaxBulkRequestsCount {
			p.listingSettings.MaxItemsPerRequest = MaxCountPerBulkRequestItem*MaxBulkRequestsCount
		}

		for ; leftCount > 0; {
			countToFetchForBulkRequest := leftCount
			if leftCount > p.listingSettings.MaxItemsPerRequest {
				countToFetchForBulkRequest = p.listingSettings.MaxItemsPerRequest
			}

			bulkItemsCount := CeilDivisionInt(countToFetchForBulkRequest, MaxCountPerBulkRequestItem)
			if bulkItemsCount > MaxBulkRequestsCount {
				bulkItemsCount = MaxBulkRequestsCount
			}

			limit := CeilDivisionInt(p.listingSettings.MaxItemsPerRequest, bulkItemsCount)
			if limit > MaxCountPerBulkRequestItem {
				limit = MaxCountPerBulkRequestItem
			}

			cursorsForBulkRequest := make([]Cursor, 0, bulkItemsCount)
			for i := 0; i < bulkItemsCount; i++ {
				cursorsForBulkRequest = append(
					cursorsForBulkRequest,
					Cursor{
						Limit:  limit,
						Offset: curPage,
					},
				)
				curPage++
				leftCount -= limit
			}
			select {
			case out <- cursorsForBulkRequest:
				continue
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func (p *Lister) fetchProductsFromAPI(
	ctx context.Context,
	cursors []Cursor,
	totalCount int,
	outputChan ItemsStream,
	filters map[string]interface{},
) {
	bulkFilters := make([]map[string]interface{}, 0, len(cursors))
	for _, cursor := range cursors {
		bulkFilter := make(map[string]interface{})
		for filterKey, filterValue := range filters {
			bulkFilter[filterKey] = filterValue
		}
		bulkFilter["recordsOnPage"] = cursor.Limit
		bulkFilter["pageNo"] = cursor.Offset
		bulkFilters = append(bulkFilters, bulkFilter)
	}

	p.reqThrottler.Throttle()

	err := p.listingDataProvider.Read(ctx, bulkFilters, func(item interface{}) {
		outputChan <- Item{
			Err:        nil,
			TotalCount: totalCount,
			Payload:    item,
		}
	})

	if err != nil {
		outputChan <- Item{
			Err:        err,
			TotalCount: totalCount,
			Payload:    nil,
		}
		return
	}
}

func (p *Lister) mergeChannels(ctx context.Context, childChans ...ItemsStream) ItemsStream {
	parentChan := make(ItemsStream, p.listingSettings.StreamBufferLength)

	var wg sync.WaitGroup
	wg.Add(len(childChans))

	for _, childChan := range childChans {
		go func(productsChildChan <-chan Item) {
			defer wg.Done()
			for prod := range productsChildChan {
				select {
				case parentChan <- prod:
					continue
				case <-ctx.Done():
					return
				}
			}
		}(childChan)
	}

	go func() {
		wg.Wait()
		close(parentChan)
	}()

	return parentChan
}

func CeilDivisionInt(x, y int) int {
	return int(math.Ceil(float64(x) / float64(y)))
}
