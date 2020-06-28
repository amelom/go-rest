//
// Copyright (C) 2020 - -
//

package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/couponmanager/common"
	"github.com/couponmanager/routers"
)

//Entry point of the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
