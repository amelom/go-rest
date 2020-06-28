//
// Copyright (C) 2020 - -
//

package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/couponmanager/common"
	"github.com/couponmanager/models"
	data "github.com/couponmanager/service"
	"github.com/go-playground/validator"
)

// GetCoupon : Handler for HTTP Post - "/coupon/"
// controler for coupon
func GetCoupon(w http.ResponseWriter, r *http.Request) {
	var Coupon *models.Coupon
	// Decode the incoming Coupon json
	err := json.NewDecoder(r.Body).Decode(&Coupon)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	// validate struct json
	v := validator.New()
	err = v.Struct(Coupon)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"malformed dating",
			500,
		)
		return
	}
	user := &Coupon
	repo := &data.CouponService{*user}
	// Calculate the value
	cal := repo.Calculate(*user)
	log.Println(cal.ItemIds)
	if j, err := json.Marshal(cal); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else if len(cal.ItemIds) == 0 {
		common.DisplayAppError(
			w,
			err,
			"404-NOT_FOUND",
			404,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}

}
