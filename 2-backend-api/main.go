package main

import (
	"net/http"

	"github.com/benfleuty/goprojects/calcapi/constants"
	"github.com/benfleuty/goprojects/calcapi/handlers"
)

func main() {
	defineApiRoutes()
}

func defineApiRoutes() {
	http.HandleFunc("/add", endpointAdd)
	http.HandleFunc("/subtract", endpointSubtract)
	http.HandleFunc("/multiply", endpointMultiply)
	http.HandleFunc("/divide", endpointDivide)

	http.ListenAndServe(":42069", nil)
}

func endpointAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetAdd(w, r)
	default:
		http.Error(w, constants.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}

func endpointSubtract(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// case http.MethodGet:
	// 	handlers.GetSubtract(w, r)
	default:
		http.Error(w, constants.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}

func endpointMultiply(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// case http.MethodGet:
	// 	handlers.GetMultiply(w, r)
	default:
		http.Error(w, constants.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}

func endpointDivide(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// case http.MethodGet:
	// 	handlers.GetDivide(w, r)
	default:
		http.Error(w, constants.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}
