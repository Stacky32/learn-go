package service

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *Service) GetRates(writer http.ResponseWriter, request *http.Request) {
	slog.Debug("GetRates called")
	res, err := s.ratesProvider.GetAll()
	if err != nil {
		slog.Error("Failed to fetch rates", "error", err)
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
