//
// Copyright (C) 2020 - -
//

package routers

// type ClientMock struct{}

// func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
// 	// mock value
// 	mockResponse := `{"price":100,"id":"MLA862633782"}`
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, mockResponse)
// 	}))
// 	defer ts.Close()
// 	return http.Get(ts.URL)
// }

// func TestCreateUser(t *testing.T) {
// 	r := mux.NewRouter().StrictSlash(false)
// 	r.HandleFunc("/coupon/", controllers.GetCoupon).Methods("POST")

// 	userJson := `{
//         "item_ids": ["MLA862633782", "MLA862633781","0"],
//         "amount":100
//     }`

// 	req, err := http.NewRequest(
// 		"POST",
// 		"/coupon/",
// 		strings.NewReader(userJson),
// 	)

// 	service.GetItems("LM11", clientMock)
// 	if err != nil {
// 		return

// 	}

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	if w.Code != 201 {
// 		t.Errorf("HTTP Status expected: 201, got: %d", w.Code)
// 	}
// }
