package store

import (
	"fmt"
	"log/slog"
	"time"

	"example.com/fxrates/internal/connector"
)

type RatesCache struct {
	baseCcy     string
	lastUpdated time.Time
	nextUpdated time.Time
	rates       map[string]float64
}

func NewRatesCache(baseCcy string) *RatesCache {
	return &RatesCache{
		baseCcy: baseCcy,
	}
}

func (c *RatesCache) IsValid() bool {
	return c.rates != nil && c.nextUpdated.After(time.Now().UTC())
}

func (c *RatesCache) Renew() error {
	slog.Debug("Updating rates cache...")
	rates, err := connector.GetFxRates(c.baseCcy)
	if err != nil {
		return fmt.Errorf("Failed to fetch FX rates: %w", err)
	}

	// Example format Mon, 04 May 2026 00:02:31 +0000
	lastUpdated, err := time.Parse(time.RFC1123Z, rates.TimeLastUpdatedUtc)
	if err != nil {
		return fmt.Errorf("Failed to parse last updated timestamp")
	}

	nextUpdated, err := time.Parse(time.RFC1123Z, rates.TimeNextUpdatedUtc)
	if err != nil {
		return fmt.Errorf("Failed to parse last updated timestamp")
	}

	c.rates = rates.Rates
	c.lastUpdated = lastUpdated
	c.nextUpdated = nextUpdated

	expiresIn := nextUpdated.Sub(time.Now().UTC())
	slog.Debug("Updated rates cache",
		"last_updated", lastUpdated,
		"expires_in", expiresIn,
	)

	return nil
}

func (c *RatesCache) GetBaseCcy() string {
	return c.baseCcy // TODO rename to reference currency
}

func (c *RatesCache) GetLastUpdated() time.Time {
	return c.lastUpdated
}

func (c *RatesCache) GetNextUpdated() time.Time {
	return c.nextUpdated
}

func (c *RatesCache) GetAll() map[string]float64 {
	return c.rates
}

func (c *RatesCache) Get(ccy string) (float64, bool) {
	r, ok := c.rates[ccy]
	return r, ok
}
