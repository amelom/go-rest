//
// Copyright (C) 2020 - -
//

package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes : router initialization
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetCouponRoutes(router)
	return router
}
