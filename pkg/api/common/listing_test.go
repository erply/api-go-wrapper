package common

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var nullSleeper = func(sleepTime time.Duration) {}

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
	inputProds := []payloadMock{
		{ID: 1},
	}
	dp := &DataProviderMock{
		CountOutputCount:  10,
		ProductsToRead:    inputProds,
		countLock:         sync.Mutex{},
		CountFiltersInput: map[string]interface{}{},
		readLock:          sync.Mutex{},
		ReadBulkFilters:   [][]map[string]interface{}{},
	}

	lister := NewLister(
		ListingSettings{
			MaxRequestsCountPerSecond: 0,
			StreamBufferLength:        0,
			MaxItemsPerRequest:        2,
			MaxFetchersCount:          2,
		},
		dp,
		nullSleeper,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	prodsChan := lister.Get(ctx, map[string]interface{}{"filterKey": "filterVal"})

	actualProds := collectProdsFromChannel(prodsChan)

	assert.Equal(t, map[string]interface{}{"filterKey": "filterVal", "pageNo": 1, "recordsOnPage": 1}, dp.CountFiltersInput)
	assert.Equal(t, ctx, dp.CountContextInput)
	assert.Equal(t, ctx, dp.ReadContextInput)

	expectedBulkFilterInputs := [][]map[string]interface{}{
		{
			{
				"filterKey":     "filterVal",
				"pageNo":        1,
				"recordsOnPage": 2,
			},
		},
		{
			{
				"filterKey":     "filterVal",
				"pageNo":        2,
				"recordsOnPage": 2,
			},
		},
		{
			{
				"filterKey":     "filterVal",
				"pageNo":        3,
				"recordsOnPage": 2,
			},
		},
		{
			{
				"filterKey":     "filterVal",
				"pageNo":        4,
				"recordsOnPage": 2,
			},
		},
		{
			{
				"filterKey":     "filterVal",
				"pageNo":        5,
				"recordsOnPage": 2,
			},
		},
	}

	assert.ElementsMatch(t, expectedBulkFilterInputs, dp.ReadBulkFilters)

	for _, prod := range actualProds {
		assert.NoError(t, prod.Err)
		assert.Equal(t, 10, prod.TotalCount)
		assert.Equal(t, payloadMock{ID: 1}, prod.Payload)
	}
}

func TestReadCountError(t *testing.T) {
	dp := &DataProviderMock{
		CountOutputCount:    1,
		CountOutputErrorStr: "some count error",
	}

	lister := NewLister(ListingSettings{}, dp, nullSleeper)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	prodsChan := lister.Get(ctx, map[string]interface{}{"filterKey": "filterVal"})

	actualProds := collectProdsFromChannel(prodsChan)

	assert.Len(t, actualProds, 1)
	assert.EqualError(t, actualProds[0].Err, "some count error")
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
