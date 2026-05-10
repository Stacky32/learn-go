package connector

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl string = "https://open.er-api.com/v6/latest/"

type FxRates struct {
	Result             string             `json:"result"`
	Provider           string             `json:"provider"`
	TimeLastUpdatedUtc string             `json:"time_last_update_utc"`
	TimeNextUpdatedUtc string             `json:"time_next_update_utc"`
	BaseCode           string             `json:"base_code"`
	Rates              map[string]float64 `json:"rates"`
}

func GetFxRates(ccy string) (*FxRates, error) {
	uri := baseUrl + ccy
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create HTTP request: %w", err)
	}

	client := http.Client{Timeout: 2 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to make HTTP request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	var fxRates FxRates
	if err := json.Unmarshal(body, &fxRates); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response: %w", err)
	}

	return &fxRates, nil
}
