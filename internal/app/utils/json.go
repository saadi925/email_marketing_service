package utils

import (
	"encoding/json"
	"net/http"
)

type MessageResponse struct {
	Message string `json:"message"`
}

// RespondJSON sends a JSON response with status code and data
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// ErrorResponse represents a JSON error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// RespondError sends a JSON error response with status code and message
func RespondError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errResponse := ErrorResponse{Message: message}
	json.NewEncoder(w).Encode(errResponse)
}

// ParseJSON parses JSON data from a request and binds it to a target struct.
func ParseJSON(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(target)
}
