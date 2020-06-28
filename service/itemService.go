//
// Copyright (C) 2020 - -
//

package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/couponmanager/common"
)

type (
	// Item struct
	Item struct {
		id    string
		price float64
	}
	// HTTPClienter interface
	HTTPClienter interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

// GetItems : function to get items from api mercadoLibre
func GetItems(id string, client HTTPClienter) Item {
	var itemResult Item
	request, err := http.NewRequest(http.MethodGet, common.AppConfig.APIURL+"items/"+id, nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	// decode body response into map
	json.NewDecoder(resp.Body).Decode(&result)
	if result["error"] == nil {
		itemResult = Item{
			id:    result["id"].(string),
			price: result["price"].(float64),
		}
	}

	return itemResult
}
