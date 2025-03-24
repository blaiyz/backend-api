package main

import (
	h "backend-api/handlers"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/lmittmann/tint"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("Got request from: %s", r.RemoteAddr))
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	})

	http.HandleFunc("/add", h.HandleAdd)
	http.HandleFunc("/subtract", h.HandleSubtract)
	http.HandleFunc("/multiply", h.HandleMultiply)
	http.HandleFunc("/divide", h.HandleDivide)
	http.HandleFunc("/sum", h.HandleSum)

	slog.Info("Server is starting")
	slog.Error(http.ListenAndServe(":3000", nil).Error())
}
