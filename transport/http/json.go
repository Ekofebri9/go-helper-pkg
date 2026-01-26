package http

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if payload != nil {
		_ = json.NewEncoder(w).Encode(payload)
	}
}
