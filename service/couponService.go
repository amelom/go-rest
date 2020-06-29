//
// Copyright (C) 2020 - -
//

package service

import (
	"net/http"

	"github.com/couponmanager/models"
)

// CouponService struct
type CouponService struct {
	C *models.Coupon
}

// ArrSort struct
type ArrSort struct {
	sum     float64
	itemArr []string
}

// ArrSortList struct
type ArrSortList struct {
	items []ArrSort
}

// Calculate : function to get items and call sort function
func (r *CouponService) Calculate(user *models.Coupon) models.Coupon {
	var Response models.Coupon
	client := &http.Client{}
	items := make(map[string]float64)
	for _, i := range user.ItemIds {
		if items[i] == 0 {
			n := GetItems(i, client)
			if n.id != "" {
				items[n.id] = n.price
			}
		}
	}

	arrPromo := []string{}
	p := orderMap(&items)
	tempP := p
	var arrsortList ArrSortList
	var temp float64
	// contar elmentos
	for j := range tempP {
		var sortList ArrSort
		arrFilter, sumArray, count := countArr(p, arrPromo, &user.Amount)
		sortList.itemArr = arrFilter
		sortList.sum = sumArray
		arrsortList.items = append(arrsortList.items, sortList)
		p[j].Value = 0
		// si los elementos restantes ya no superan el monto maximo se interrumpe el ciclo
		if count >= len(p) {
			temp = arrsortList.items[0].sum
			arrPromo = arrsortList.items[0].itemArr
			break
		}
	}

	for _, k := range arrsortList.items {
		if k.sum > temp {
			arrPromo = k.itemArr
		}
	}

	Response.Amount = user.Amount
	Response.ItemIds = arrPromo
	return Response
}

func countArr(p PairList, arrPromo []string, amount *float64) ([]string, float64, int) {
	sumArray := 0
	count := 0
	for _, k := range p {
		if sumArray+k.Value <= int(*amount) {
			count++
			sumArray = sumArray + k.Value
			if k.Value != 0 {
				arrPromo = append(arrPromo, k.Key)
			}
		}
	}
	return arrPromo, float64(sumArray), count
}
