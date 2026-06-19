package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(rw http.ResponseWriter, status int, payload any) {
	
	rw.Header().Set("Content-Type", "Application/json")

	rw.WriteHeader(status)

	_ = json.NewEncoder(rw).Encode(payload)
}
