package service

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

func (s *Service) GetRate(writer http.ResponseWriter, request *http.Request) {
	baseCcy := strings.ToUpper(request.PathValue("base"))
	quoteCcy := strings.ToUpper(request.PathValue("quote"))

	slog.Debug("GetRate called", "base", baseCcy, "quote", quoteCcy)

	res, err := s.ratesProvider.Get(baseCcy, quoteCcy)
	if err != nil {
		slog.Error("Failed to marshal response", "error", err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("failed to fetch rates data"))
	}

	body, err := json.Marshal(res)
	if err != nil {
		slog.Error("Failed to marshal response", "error", err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("failed to fetch rates data"))
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(body)
}
