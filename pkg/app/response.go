package app

import (
	"log"

	"encoding/json"
	"net/http"
)

type ErrorMessageResponse struct {
	Message string `json:"message"`
}

func SendJSONResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("json.NewEncoder(w).Encode(data) error: %v", err)
	}
}

func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	log.Printf("error: %s", message)
	SendJSONResponse(w, ErrorMessageResponse{Message: message}, status)
}

func SendSuccessResponse(w http.ResponseWriter, message string) {
	SendJSONResponse(w, map[string]string{
		"message": message,
	}, http.StatusOK)
}

func SendTokenResponse(w http.ResponseWriter, token string) {
	SendJSONResponse(w, map[string]string{
		"token": token,
	}, http.StatusOK)
}
