//
// Copyright (C) 2020 - -
//

package routers

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestSetCouponRoutes(t *testing.T) {
	r := mux.NewRouter().StrictSlash(false)
	SetCouponRoutes(r)
}
