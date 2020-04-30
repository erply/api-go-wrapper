package api

import (
	"context"
	"testing"
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
	//works
	t.Run("test GetUserRights", func(t *testing.T) {
		resp, err := cli.GetUserRights(context.Background(), map[string]string{
			"getCurrentUser": "1",
		})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	//works
	t.Run("test GetEmployees", func(t *testing.T) {
		resp, err := cli.GetEmployees(context.Background(), map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	//works
	t.Run("test GetBusinessAreas", func(t *testing.T) {
		resp, err := cli.GetBusinessAreas(context.Background(), map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	//works
	t.Run("test GetCurrencies", func(t *testing.T) {
		resp, err := cli.GetCurrencies(context.Background(), map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	//works
	t.Run("test LogProcessingOfCustomerData", func(t *testing.T) {
		if err := cli.LogProcessingOfCustomerData(context.Background(), map[string]string{}); err != nil {
			t.Error(err)
			return
		}
	})
	//works
	t.Run("test GetCountries", func(t *testing.T) {
		resp, err := cli.GetCountries(context.Background(), map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
