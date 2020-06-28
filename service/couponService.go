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

	arrIdsItems := []string{}
	var count float64 = 0
	p := orderMap(&items)
	for _, k := range p {
		if count+float64(k.Value) <= user.Amount {
			count = count + float64(k.Value)
			arrIdsItems = append(arrIdsItems, k.Key)
		}
	}
	Response.Amount = user.Amount
	Response.ItemIds = arrIdsItems
	return Response
}
