//
// Copyright (C) 2020 - -
//

package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	configuration struct {
		APIURL, Server string
	}
)

// DisplayAppError : error handling function
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {

	errTemp := ""
	if handlerError != nil {
		errTemp = handlerError.Error()
		log.Printf("AppError]: %s\n", handlerError)
	}

	errObj := appError{
		Error:      errTemp,
		Message:    message,
		HTTPStatus: code,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration
var urlVal = "config/config.json"

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open(urlVal)
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
