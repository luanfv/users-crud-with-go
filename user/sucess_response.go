package user

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Data any `json:"data"`
}

func sendJSONSuccessResponse(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(successResponse{Data: data})
}