package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/benfleuty/goprojects/calcapi/constants"
)

func SetResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Error writing response with data\n%v\n", data)
		http.Error(w, constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
	}
}

type CalculationResult struct {
	Operation string `json:"operation"`
}

func GetAdd(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("%s %s Error reading request body: %v", r.Method, r.URL.Path, err)
		slog.Error(msg)
		http.Error(w, msg, 400)
		return
	}

	slog.Info(fmt.Sprintf("%s %s request with body: %s", r.Method, r.URL.Path, bodyBytes))

	var data []CalculationResult
	obj := CalculationResult{}
	data = append(data, obj)
	SetResponse(w, 200, data)
}

func GetSubtract(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("%s %s Error reading request body: %v", r.Method, r.URL.Path, err)
		slog.Error(msg)
		http.Error(w, msg, 400)
		return
	}

	slog.Info(fmt.Sprintf("%s %s request with body: %s", r.Method, r.URL.Path, bodyBytes))

	var data []CalculationResult
	obj := CalculationResult{}
	data = append(data, obj)
	SetResponse(w, 200, data)
}

func GetMultiply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("%s %s Error reading request body: %v", r.Method, r.URL.Path, err)
		slog.Error(msg)
		http.Error(w, msg, 400)
		return
	}

	slog.Info(fmt.Sprintf("%s %s request with body: %s", r.Method, r.URL.Path, bodyBytes))

	var data []CalculationResult
	obj := CalculationResult{}
	data = append(data, obj)
	SetResponse(w, 200, data)
}

func GetDivide(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("%s %s Error reading request body: %v", r.Method, r.URL.Path, err)
		slog.Error(msg)
		http.Error(w, msg, 400)
		return
	}

	slog.Info(fmt.Sprintf("%s %s request with body: %s", r.Method, r.URL.Path, bodyBytes))

	var data []CalculationResult
	obj := CalculationResult{}
	data = append(data, obj)
	SetResponse(w, 200, data)
}
