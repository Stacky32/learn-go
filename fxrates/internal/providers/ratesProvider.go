package providers

import (
	"fmt"
	"time"

	"example.com/fxrates/internal/store"
)

type RatesData struct {
	EffectiveAt string             `json:"effective_at"`
	Rates       map[string]float64 `json:"rates"`
}

type RatesProvider interface {
	GetAll() (*RatesData, error)
	Get(base, quote string) (*RatesData, error)
}

type RatesEngine struct {
	store *store.RatesCache
}

func NewRatesEngine(store *store.RatesCache) *RatesEngine {
	return &RatesEngine{
		store: store,
	}
}

func (e *RatesEngine) GetAll() (*RatesData, error) {
	if err := e.renewStore(); err != nil {
		return nil, err
	}

	baseCcy := e.store.GetBaseCcy()
	effectiveAt := e.store.GetLastUpdated().Format(time.RFC3339)
	storedRates := e.store.GetAll()
	rates := make(map[string]float64, len(storedRates))

	for k, v := range storedRates {
		rates[baseCcy+k] = v
	}

	return &RatesData{
		EffectiveAt: effectiveAt,
		Rates:       rates,
	}, nil
}

func (e *RatesEngine) Get(base, quote string) (*RatesData, error) {
	if err := e.renewStore(); err != nil {
		return nil, err
	}

	referenceCcy := e.store.GetBaseCcy()
	effectiveAt := e.store.GetLastUpdated().Format(time.RFC3339)

	baseFx, ok := e.store.Get(base)
	if !ok {
		return nil, fmt.Errorf("Missing quote for %s_%s", referenceCcy, base)
	}

	quoteFx, ok := e.store.Get(quote)
	if !ok {
		return nil, fmt.Errorf("Missing quote for %s_%s", referenceCcy, quote)
	}

	// E.g. EURGBP = USDGBP / USDEUR
	fx := quoteFx / baseFx
	rates := map[string]float64{
		base + quote: fx,
	}

	return &RatesData{
		EffectiveAt: effectiveAt,
		Rates:       rates,
	}, nil
}

func (e *RatesEngine) renewStore() error {
	if e.store.IsValid() {
		return nil
	}

	return e.store.Renew()
}
