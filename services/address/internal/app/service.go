package app

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/mathaono/freight-simulator/services/address/internal/cache"
	"github.com/mathaono/freight-simulator/services/address/internal/cep"
)

var reDigits = regexp.MustCompile(`\D`)

type Service struct {
	cache    *cache.RedisCache
	provider cep.Provider
}

func NewService(c *cache.RedisCache, p cep.Provider) *Service {
	return &Service{
		cache:    c,
		provider: p,
	}
}

func NormalizeCEP(cep string) (string, error) {
	digits := reDigits.ReplaceAllString(cep, "")
	if len(digits) != 8 {
		return "", errors.New("cep inválido: deve conter 8 dígitos")
	}
	return digits, nil
}

type Address struct {
	CEP       string  `json:"cep"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (s *Service) Lookup(ctx context.Context, rawCep string) (Address, error) {
	cep, err := NormalizeCEP(strings.TrimSpace(rawCep))
	if err != nil {
		return Address{}, err
	}

	// Verifica no cache
	if cached, ok, _ := s.cache.GetCEP(ctx, cep); ok {
		var out Address
		_ = json.Unmarshal([]byte(cached), &out)
	}

	// Busca no provedor
	ctx, cancel := context.WithTimeout(ctx, 1200*time.Millisecond)
	defer cancel()

	data, err := s.provider.Lookup(ctx, cep)
	if err != nil {
		return Address{}, err
	}

	out := Address{
		CEP:       cep,
		City:      data.City,
		State:     data.State,
		Latitude:  data.Lat,
		Longitude: data.Lon,
	}

	b, err := json.Marshal(out)
	if err != nil {
		return Address{}, nil
	}
	_ = s.cache.SetCEP(ctx, cep, string(b))

	return out, nil
}
