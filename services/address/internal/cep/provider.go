package cep

import "context"

type Data struct {
	CEP   string  `json:"cep"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

type Provider interface {
	Lookup(ctx context.Context, cep string) (Data, error)
}

type MockProvider struct{}

func (MockProvider) Lookup(ctx context.Context, cep string) (Data, error) {
	// Mock basico
	return Data{
		CEP:   cep,
		City:  "São Paulo",
		State: "SP",
		Lat:   -23.55052,
		Lon:   -46.633308,
	}, nil
}
