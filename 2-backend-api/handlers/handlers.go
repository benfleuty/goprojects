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

func parseBasicCalculationRequest(w http.ResponseWriter, r *http.Request, sUuid string) (BasicCalculationRequest, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("%s %s %s Error reading request body: %v", sUuid, r.Method, r.URL.Path, err)
		slog.Error(msg)
		http.Error(w, msg, 500)
		return BasicCalculationRequest{}, err
	}

	slog.Info(fmt.Sprintf("%s %s %s request with body: %s", sUuid, r.Method, r.URL.Path, bodyBytes))
	request := BasicCalculationRequest{}
	if err := json.Unmarshal(bodyBytes, &request); err != nil {
		msg := fmt.Sprintf("%s Error parsing body: %v", sUuid, err)
		slog.Error(msg)
		http.Error(w, msg, 400)
	}

	return request, nil
}

func GetAdd(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	sUuid := uuid.NewString()

	var request BasicCalculationRequest
	if req, err := parseBasicCalculationRequest(w, r, sUuid); err != nil {
		return
	} else {
		request = req
	}

	var response BasicCalculationResponse
	response.Result = request.Number1 + request.Number2

	slog.Info(fmt.Sprintf("%s Sending response 200 OK with body: %+v", sUuid, response))

	SetResponse(w, 200, response)
}

func GetSubtract(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	sUuid := uuid.NewString()

	var request BasicCalculationRequest
	if req, err := parseBasicCalculationRequest(w, r, sUuid); err != nil {
		return
	} else {
		request = req
	}

	var response BasicCalculationResponse
	response.Result = request.Number1 - request.Number2
	SetResponse(w, 200, response)
}

func GetMultiply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	sUuid := uuid.NewString()

	var request BasicCalculationRequest
	if req, err := parseBasicCalculationRequest(w, r, sUuid); err != nil {
		return
	} else {
		request = req
	}

	var response BasicCalculationResponse
	response.Result = request.Number1 * request.Number2
	SetResponse(w, 200, response)
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
