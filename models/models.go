//
// Copyright (C) 2020 - -
//

package models

type (
	// Coupon struct
	Coupon struct {
		ItemIds []string `json:"item_ids" validate:"required"`
		Amount  int      `json:"amount" validate:"required,numeric,gte=0,lte=999"`
	}
)
