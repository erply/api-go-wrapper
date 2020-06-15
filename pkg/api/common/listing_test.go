package common

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"sort"
	"sync"
	"testing"
	"time"
)

var NullSleeper = func(sleepTime time.Duration) {}

type payloadMock struct {
	ID int
}

type DataProviderMock struct {
	countLock           sync.Mutex
	CountContextInput   context.Context
	CountFiltersInput   map[string]interface{}
	CountOutputCount    int
	CountOutputErrorStr string

	readLock         sync.Mutex
	ReadContextInput context.Context
	ProductsToRead   []payloadMock
	ReadErrorStr     string
	ReadBulkFilters  [][]map[string]interface{}
}

func (dpm *DataProviderMock) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	dpm.countLock.Lock()
	defer dpm.countLock.Unlock()

	dpm.CountContextInput = ctx
	dpm.CountFiltersInput = filters

	var err error
	if dpm.CountOutputErrorStr != "" {
		err = errors.New(dpm.CountOutputErrorStr)
	}

	return dpm.CountOutputCount, err
}

func (dpm *DataProviderMock) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	dpm.readLock.Lock()
	defer dpm.readLock.Unlock()

	dpm.ReadContextInput = ctx
	dpm.ReadBulkFilters = append(dpm.ReadBulkFilters, bulkFilters)

	for _, prod := range dpm.ProductsToRead {
		callback(prod)
	}

	var err error
	if dpm.ReadErrorStr != "" {
		err = errors.New(dpm.ReadErrorStr)
	}

	return err
}

func TestReadingSuccess(t *testing.T) {
	testCases := []struct {
		name                     string
		total                    int
		inputProds               []payloadMock
		listingSettings          ListingSettings
		expectedBulkFilterInputs func() [][]map[string]interface{}
		expectedProdIDs          []int
	}{
		{
			name:  "too small request limit",
			total: 10,
			inputProds: []payloadMock{
				{ID: 1},
				{ID: 2},
			},
			expectedProdIDs: []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        0,
				MaxItemsPerRequest:        2,
				MaxFetchersCount:          2,
			},
			expectedBulkFilterInputs: func() (res [][]map[string]interface{}) {
				expectedBulkRequestsCount := 5
				res = make([][]map[string]interface{}, 0, expectedBulkRequestsCount)
				expectedBulkInputCount := 1
				for i := 0; i < expectedBulkRequestsCount; i++ {
					bulkItems := make([]map[string]interface{}, 0, expectedBulkInputCount)
					for y := 0; y < expectedBulkInputCount; y++ {
						bulkItems = append(bulkItems, map[string]interface{}{
							"filterKey":     "filterVal",
							"pageNo":        i + 1,
							"recordsOnPage": 2,
						})
					}
					res = append(res, bulkItems)
				}
				return
			},
		},
		{
			name:  "max request limit",
			total: 10001,
			inputProds: []payloadMock{
				{ID: 3},
			},
			expectedProdIDs: []int{3, 3},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        10,
				MaxItemsPerRequest:        10000,
				MaxFetchersCount:          2,
			},
			expectedBulkFilterInputs: func() (res [][]map[string]interface{}) {
				res = make([][]map[string]interface{}, 0, 2)
				bulkItems := make([]map[string]interface{}, 0, 100)
				for y := 0; y < 100; y++ {
					bulkItems = append(bulkItems, map[string]interface{}{
						"filterKey":     "filterVal",
						"pageNo":        y + 1,
						"recordsOnPage": 100,
					})
				}
				res = append(res, bulkItems)
				res = append(res, []map[string]interface{}{
					{
						"filterKey":     "filterVal",
						"pageNo":        101,
						"recordsOnPage": 100,
					},
				})
				return
			},
		},
		{
			name:  "fetch all in one request",
			total: 1000,
			inputProds: []payloadMock{
				{ID: 4},
			},
			expectedProdIDs: []int{4},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        10,
				MaxItemsPerRequest:        10000,
				MaxFetchersCount:          10,
			},
			expectedBulkFilterInputs: func() (res [][]map[string]interface{}) {
				res = make([][]map[string]interface{}, 0, 2)
				bulkItems := make([]map[string]interface{}, 0, 100)
				for y := 0; y < 10; y++ {
					bulkItems = append(bulkItems, map[string]interface{}{
						"filterKey":     "filterVal",
						"pageNo":        y + 1,
						"recordsOnPage": 100,
					})
				}
				res = append(res, bulkItems)
				return
			},
		},
		{
			name:  "max items per request is impossible",
			total: 100,
			inputProds: []payloadMock{
				{ID: 5},
			},
			expectedProdIDs: []int{5},
			listingSettings: ListingSettings{
				MaxRequestsCountPerSecond: 0,
				StreamBufferLength:        10,
				MaxItemsPerRequest:        10001,
				MaxFetchersCount:          10,
			},
			expectedBulkFilterInputs: func() (res [][]map[string]interface{}) {
				return [][]map[string]interface{}{
					{
						{
							"filterKey":     "filterVal",
							"pageNo":        1,
							"recordsOnPage": 100,
						},
					},
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dp := &DataProviderMock{
				CountOutputCount:  testCase.total,
				ProductsToRead:    testCase.inputProds,
				countLock:         sync.Mutex{},
				CountFiltersInput: map[string]interface{}{},
				readLock:          sync.Mutex{},
				ReadBulkFilters:   [][]map[string]interface{}{},
			}
			lister := NewLister(
				testCase.listingSettings,
				dp,
				NullSleeper,
			)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			prodsChan := lister.Get(ctx, map[string]interface{}{"filterKey": "filterVal"})

			actualProds := collectProdsFromChannel(prodsChan)

			assert.Equal(t, map[string]interface{}{"filterKey": "filterVal", "pageNo": 1, "recordsOnPage": 1}, dp.CountFiltersInput)
			assert.Equal(t, ctx, dp.CountContextInput)
			assert.Equal(t, ctx, dp.ReadContextInput)
			assert.ElementsMatch(t, testCase.expectedBulkFilterInputs(), dp.ReadBulkFilters)

			actualProgressCounts := make([]int, 0, len(actualProds))
			actualProdIDs := make([]int, 0, len(actualProds))
			for _, prod := range actualProds {
				assert.NoError(t, prod.Err)
				assert.Equal(t, testCase.total, prod.TotalCount)
				assert.IsType(t, prod.Payload, payloadMock{})
				actualProdIDs = append(actualProdIDs, prod.Payload.(payloadMock).ID)
			}

			sort.Ints(actualProgressCounts)
			sort.Ints(actualProdIDs)

			assert.Equal(t, testCase.expectedProdIDs, actualProdIDs)
		})
	}
}

func TestReadCountError(t *testing.T) {
	dp := &DataProviderMock{
		CountOutputCount:    1,
		CountOutputErrorStr: "some count error",
	}

	lister := NewLister(ListingSettings{}, dp, NullSleeper)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	prodsChan := lister.Get(ctx, map[string]interface{}{"filterKey": "filterVal"})

	actualProds := collectProdsFromChannel(prodsChan)

	assert.Len(t, actualProds, 1)
	assert.EqualError(t, actualProds[0].Err, "some count error")
}

func TestCancelReading(t *testing.T) {
	dp := &DataProviderMock{
		CountOutputCount: 10,
		ProductsToRead: []payloadMock{
			{ID: 4},
		},
		countLock:         sync.Mutex{},
		CountFiltersInput: map[string]interface{}{},
		readLock:          sync.Mutex{},
		ReadBulkFilters:   [][]map[string]interface{}{},
	}
	lister := NewLister(
		ListingSettings{
			MaxRequestsCountPerSecond: 0,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        100,
			MaxFetchersCount:          10,
		},
		dp,
		NullSleeper,
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	prodsChan := lister.Get(ctx, map[string]interface{}{"filterKey": "filterVal"})

	actualProds := collectProdsFromChannel(prodsChan)
	assert.Len(t, actualProds, 0)
	assert.Len(t, dp.ReadBulkFilters, 0)
}

func TestReadItemsError(t *testing.T) {
	dp := &DataProviderMock{
		CountOutputCount:    1,
		ReadErrorStr: "some read items error",
	}

	lister := NewLister(ListingSettings{}, dp, NullSleeper)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	prodsChan := lister.Get(ctx, map[string]interface{}{})

	actualProds := collectProdsFromChannel(prodsChan)

	assert.Len(t, actualProds, 1)
	assert.EqualError(t, actualProds[0].Err, "some read items error")
}

func collectProdsFromChannel(prodsChan ItemsStream) []Item {
	actualProds := make([]Item, 0)
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for prod := range prodsChan {
			actualProds = append(actualProds, prod)
		}
	}()

mainLoop:
	for {
		select {
		case <-doneChan:
			break mainLoop
		case <-time.After(time.Second):
			break mainLoop
		}
	}

	return actualProds
}
