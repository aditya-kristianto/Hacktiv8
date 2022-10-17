package controller

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, payload *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(payload.Status)
	json.NewEncoder(w).Encode(payload)
}
