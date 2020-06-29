package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetCouponDataBad(t *testing.T) {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/coupon/", GetCoupon).Methods("POST")

	couponJSON := ` `

	req, err := http.NewRequest(
		"POST",
		"/coupon/",
		strings.NewReader(couponJSON),
	)

	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code == 500 {
		return
	}
}
func TestGetCouponDataBadFormat(t *testing.T) {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/coupon/", GetCoupon).Methods("POST")

	couponJSON := `{
        "item_ids": ["MLA862633782", "MLA862633781","0"]
    }`

	req, err := http.NewRequest(
		"POST",
		"/coupon/",
		strings.NewReader(couponJSON),
	)

	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code == 400 {
		return
	}
}
