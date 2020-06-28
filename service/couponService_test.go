//
// Copyright (C) 2020 - -
//

package service

import (
	"log"
	"testing"

	"github.com/couponmanager/models"
)

func TestCalculate(t *testing.T) {
	t.Run("test array empty", func(t *testing.T) {
		var a []string
		Coupons := new(models.Coupon)
		Coupons.ItemIds = a
		Coupons.Amount = 100
		log.Println(Coupons)

		repo := &CouponService{Coupons}
		repo.Calculate(Coupons)
	})
}

// func TestCalculateArryFull(t *testing.T) {
// 	t.Run("test array full", func(t *testing.T) {
// 		var a []string
// 		a = append(a, "LM")
// 		Coupons := new(models.Coupon)
// 		Coupons.ItemIds = a
// 		Coupons.Amount = 100
// 		log.Println(Coupons)

// 		repo := &CouponService{Coupons}

// 		repo.Calculate(Coupons)
// 	})
// }
