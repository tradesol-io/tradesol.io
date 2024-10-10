package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"description,omitempty"`
	Hint        string `json:"hint,omitempty"`
	Example     string `json:"example,omitempty"`
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, errorResp ErrorResponse) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResp)
}
