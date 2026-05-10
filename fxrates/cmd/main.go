package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"example.com/fxrates/internal/providers"
	"example.com/fxrates/internal/service"
	"example.com/fxrates/internal/store"
)

const baseCCy string = "USD"

func main() {
	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, handlerOpts))
	slog.SetDefault(logger)

	s := store.NewRatesCache(baseCCy)
	e := providers.NewRatesEngine(s)
	svc := service.NewService(e)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /rates", svc.GetRates)
	mux.HandleFunc("GET /rates/{base}/{quote}", svc.GetRate)

	handler := contentTypeMiddleware("application/json", mux)

	const port string = ":3000"
	srv := http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	slog.Info("Listening...", "port", port)
	if err := srv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			slog.Info("Server shutting down")
		} else {
			slog.Error("Server shutting down", "error", err)
		}
	}
}

func contentTypeMiddleware(contentType string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		next.ServeHTTP(w, r)
	})
}
