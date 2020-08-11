package utils

import (
	"encoding/json"
	"net/http"
)

// JSON responde uma requisição em dados json
func JSON(w http.ResponseWriter, status int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}