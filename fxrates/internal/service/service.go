package service

import (
	"example.com/fxrates/internal/providers"
)

type Service struct {
	ratesProvider providers.RatesProvider
}

func NewService(ratesProvider providers.RatesProvider) *Service {
	return &Service{
		ratesProvider: ratesProvider,
	}
}
