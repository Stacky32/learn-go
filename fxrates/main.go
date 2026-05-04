package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	rates, err := GetFxRates("GBP")
	if err != nil {
		slog.Error("Failed to get FX rate", "error", err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(rates, "", "\t")
	if err != nil {
		slog.Error("Marshalling failed", "error", err)
		os.Exit(1)
	}

	fmt.Println(string(b))

	gbpaud := rates.Rates["AUD"] // 1.89
	gbpchf := rates.Rates["CHF"] // 1.06
	audchf := gbpchf / gbpaud    // 1.06/1.89

	slog.Info("inferred rates", "GBPAUD", gbpaud, "GBPCHF", gbpchf, "AUDCHF", audchf)

	gbpusd := rates.Rates["USD"]
	gbpczk := rates.Rates["CZK"]
	usdczk := gbpczk / gbpusd

	slog.Info("inferred rates", "GBPCZK", gbpczk, "USDCZK", usdczk)
}

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
