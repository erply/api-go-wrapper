package api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//here the "works" indicator will be under each test separately
func TestApiRequests(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("test GetUserRights", func(t *testing.T) {
		resp, err := cli.GetUserRights(context.Background(), map[string]string{
			"getCurrentUser": "1",
		})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test GetEmployees", func(t *testing.T) {
		resp, err := cli.GetEmployees(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test GetBusinessAreas", func(t *testing.T) {
		resp, err := cli.GetBusinessAreas(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test GetCurrencies", func(t *testing.T) {
		resp, err := cli.GetCurrencies(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test LogProcessingOfCustomerData", func(t *testing.T) {
		assert.NoError(t, cli.LogProcessingOfCustomerData(context.Background(), map[string]string{}))
	})

	t.Run("test GetCountries", func(t *testing.T) {
		resp, err := cli.GetCountries(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test SaveEvent", func(t *testing.T) {
		layout := "2006-01-02 15:04:05"
		d := time.Now
		assert.NoError(t, err)
		eventID, err := cli.SaveEvent(context.Background(), map[string]string{
			"startTime": d().Format(layout),
			"endTime":   d().Format(layout),
			"typeID":    "APPOINTMENT",
		})
		assert.NoError(t, err)
		assert.NotEqual(t, 0, eventID)
	})
}
