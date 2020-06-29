//
// Copyright (C) 2020 - -
//

package service

import (
	"testing"

	"github.com/couponmanager/models"
)

func TestCalculate(t *testing.T) {
	t.Run("test array empty", func(t *testing.T) {
		var a []string
		Coupons := new(models.Coupon)
		Coupons.ItemIds = a
		Coupons.Amount = 100

		repo := &CouponService{Coupons}
		repo.Calculate(Coupons)
	})
}
