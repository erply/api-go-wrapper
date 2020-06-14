package common

import (
	"context"
	"math"
	"sync"
	"time"
)

const MaxItemsPerBulkRequest = 10000
const DefaultMaxCountPerRequest = 1000
const MaxFetchersCount = 100

type ListingSettings struct {
	MaxRequestsCountPerSecond uint
	StreamBufferLength        uint
	MaxCountPerRequest        uint
	MaxFetchersCount          uint
}

type Cursor struct {
	Limit  uint
	Offset uint
}

type ItemsStream chan Item

type Item struct {
	Err           error
	TotalCount    uint
	ProgressCount uint
	Payload       interface{}
}

func setListingSettingsDefaults(settingsFromInput ListingSettings) ListingSettings {
	if settingsFromInput.MaxRequestsCountPerSecond == 0 {
		settingsFromInput.MaxRequestsCountPerSecond = 5
	}

	if settingsFromInput.MaxCountPerRequest > MaxItemsPerBulkRequest {
		settingsFromInput.MaxCountPerRequest = MaxItemsPerBulkRequest
	}

	if settingsFromInput.MaxCountPerRequest == 0 {
		settingsFromInput.MaxCountPerRequest = DefaultMaxCountPerRequest
	}

	if settingsFromInput.MaxFetchersCount == 0 {
		settingsFromInput.MaxFetchersCount = MaxFetchersCount
	}

	return settingsFromInput
}

type DataProvider interface {
	Count(ctx context.Context, filters map[string]interface{}) (int, error)
	Read(ctx context.Context, limit, offset uint,  filters map[string]interface{}, callback func(item interface{})) error
}

type Lister struct {
	listingSettings     ListingSettings
	reqThrottler        Throttler
	listingDataProvider DataProvider
}

func NewLister(settings ListingSettings, dataProvider DataProvider) *Lister {
	settings = setListingSettingsDefaults(settings)

	thrl := NewSleepThrottler(settings.MaxRequestsCountPerSecond, func(sleepTime time.Duration) {
		time.Sleep(sleepTime)
	})

	return &Lister{
		listingSettings:     settings,
		reqThrottler:        thrl,
		listingDataProvider: dataProvider,
	}
}

func (p *Lister) Get(ctx context.Context, filters map[string]interface{}) ItemsStream {
	outputChan := make(ItemsStream, p.listingSettings.StreamBufferLength)
	defer close(outputChan)

	p.reqThrottler.Throttle()

	totalCount, err := p.listingDataProvider.Count(ctx, filters)
	if err != nil {
		outputChan <- Item{
			Err:           err,
			TotalCount:    uint(totalCount),
			ProgressCount: 0,
			Payload:       nil,
		}
		return outputChan
	}

	cursorChan := p.getCursors(ctx, uint(totalCount))

	childChans := make([]ItemsStream, p.listingSettings.MaxFetchersCount)
	for i := 0; i < int(p.listingSettings.MaxFetchersCount); i++ {
		childChan := p.fetchProductsChunk(ctx, cursorChan, uint(totalCount), filters)
		childChans = append(childChans, childChan)
	}

	return p.mergeChannels(ctx, outputChan, childChans...)
}

func (p *Lister) fetchProductsChunk(ctx context.Context, cursorChan chan Cursor, totalCount uint, filters map[string]interface{}) ItemsStream {
	prodStream := make(chan Item, p.listingSettings.StreamBufferLength)
	go func() {
		defer close(prodStream)
		for cursor := range cursorChan {
			p.fetchProductsFromAPI(ctx, cursor.Offset, cursor.Limit, totalCount, prodStream, filters)

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

func (p *Lister) getCursors(ctx context.Context, totalCount uint) chan Cursor {
	chunksCount := ceilDivisionResult(totalCount, p.listingSettings.MaxCountPerRequest)

	out := make(chan Cursor, chunksCount)
	defer close(out)

	var i uint
	for i = 0; i < chunksCount; i++ {
		select {
		case out <- Cursor{Limit: p.listingSettings.MaxCountPerRequest, Offset: i + 1}:
		case <-ctx.Done():
			return out
		}
	}

	return out
}

func (p *Lister) fetchProductsFromAPI(
	ctx context.Context,
	offset, limit, totalCount uint,
	outputChan ItemsStream,
	filters map[string]interface{},
) {
	p.reqThrottler.Throttle()

	err := p.listingDataProvider.Read(ctx, offset, limit, filters, func(item interface{}) {
		outputChan <- Item{
			Err:           nil,
			TotalCount:    totalCount,
			ProgressCount: 0,
			Payload:       item,
		}
	})

	if err != nil {
		outputChan <- Item{
			Err:           err,
			TotalCount:    totalCount,
			ProgressCount: 0,
			Payload:       nil,
		}
		return
	}
}

func (p *Lister) mergeChannels(ctx context.Context, parentChan ItemsStream, childChans ...ItemsStream) ItemsStream {
	var wg sync.WaitGroup
	wg.Add(len(childChans))

	for _, childChan := range childChans {
		go func(productsChildChan <-chan Item) {
			defer wg.Done()
			var i uint = 0
			for prod := range productsChildChan {
				i++
				prod.ProgressCount = i
				select {
				case parentChan <- prod:
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

func ceilDivisionResult(x, y uint) uint {
	return uint(math.Ceil(float64(x) / float64(y)))
}
