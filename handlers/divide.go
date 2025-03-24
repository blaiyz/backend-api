package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type DivideRequest struct {
	Divident int `json:"divident"`
	Divisor  int `json:"divisor"`
}

func HandleDivide(w http.ResponseWriter, r *http.Request) {
	slog.Info("/divide hit")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		slog.Error("Invalid method")
		return
	}

	var req DivideRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		// Bad request
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Error(fmt.Sprintf("Bad request: %v", r.Body))
		return
	}

	if req.Divisor == 0 {
		http.Error(w, "Cannot divide by 0", http.StatusBadRequest)
		slog.Error("Attempted to divide by 0")
		return
	}

	result := req.Divident / req.Divisor
	response := Response{
		Result: result,
	}
	slog.Info(fmt.Sprintf("Computed %d/%d=%d", req.Divident, req.Divisor, result))

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
