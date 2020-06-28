//
// Copyright (C) 2020 - -
//

package routers

import (
	"github.com/couponmanager/controllers"
	"github.com/gorilla/mux"
)

// SetCouponRoutes : coupon routes
func SetCouponRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/coupon/", controllers.GetCoupon).Methods("POST")
	return router
}
