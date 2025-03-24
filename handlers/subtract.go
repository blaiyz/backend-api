package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)


func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	slog.Info("/subtract hit")

	if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		slog.Error("Invalid method")
        return
    }

	var req BinaryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// Bad request
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Error(fmt.Sprintf("Bad request: %v", r.Body))
		return
	}

	result := req.Num1 - req.Num2
	response := Response{
		Result: result,
	}
	slog.Info(fmt.Sprintf("Computed %d-%d=%d", req.Num1, req.Num2, result))

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}