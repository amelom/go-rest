//
// Copyright (C) 2020 - -
//

package models

type (
	// Coupon struct
	Coupon struct {
		ItemIds []string `json:"item_ids" validate:"required"`
		Amount  float64  `json:"amount" validate:"required,numeric"`
	}
)
