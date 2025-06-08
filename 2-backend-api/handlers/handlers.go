package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/benfleuty/goprojects/calcapi/constants"
	"github.com/google/uuid"
)

func SetResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Error writing response with data\n%v\n", data)
		http.Error(w, constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
	}
}

// Deprecated: Use BasicCalculationRequest
type CalculationResult struct {
	Operation string `json:"operation"`
}

type BasicCalculationRequest struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type BasicCalculationResponse struct {
	Result int `json:"result"`
}

func GetAdd(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	sUuid := uuid.NewString()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("%s %s %s Error reading request body: %v", sUuid, r.Method, r.URL.Path, err)
		slog.Error(msg)
		http.Error(w, msg, 500)
		return
	}

	slog.Info(fmt.Sprintf("%s %s %s request with body: %s", sUuid, r.Method, r.URL.Path, bodyBytes))
	body := &BasicCalculationRequest{}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		msg := fmt.Sprintf("%s Error parsing body!", sUuid)
		slog.Error(msg)
		http.Error(w, msg, 400)
	}

	responseBody := &BasicCalculationResponse{}
	responseBody.Result = body.Number1 + body.Number2

	slog.Info(fmt.Sprintf("%s Sending response 200 OK with body: %+v", sUuid, responseBody))

	SetResponse(w, 200, responseBody)
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
