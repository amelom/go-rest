//
// Copyright (C) 2020 - -
//
package service

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type ClientMock struct{}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	// mock value
	mockResponse := `{"price":100,"id":"MLA862633782"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, mockResponse)
	}))
	defer ts.Close()
	return http.Get(ts.URL)
}

func TestGetItems(t *testing.T) {
	t.Run("it should return mock value", func(t *testing.T) {
		clientMock := &ClientMock{}
		got := GetItems("LM11", clientMock)
		want := `{"price":100,"id":"MLA862633782"}`
		if !reflect.DeepEqual(got, want) {
			log.Println(got, want)
		}
	})
}
