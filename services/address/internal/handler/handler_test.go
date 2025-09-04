package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mathaono/freight-simulator/services/address/internal/app"
	"github.com/mathaono/freight-simulator/services/address/internal/cep"
	"github.com/mathaono/freight-simulator/services/address/internal/handler"
)

type mockCache struct{}

func (m *mockCache) GetCEP(ctx context.Context, cep string) (string, bool, error) {
	return "", false, nil
}

func (m *mockCache) SetCEP(ctx context.Context, cep, json string) error {
	return nil
}

type mockProvider struct {
	data cep.Data
}

func (m *mockProvider) Lookup(ctx context.Context, cep string) (cep.Data, error) {
	return m.data, nil
}

func TestHandler_Success(t *testing.T) {

	svc := app.NewService(&mockCache{}, &mockProvider{data: cep.Data{
		CEP:   "13100000",
		City:  "Campinas",
		State: "SP",
		Lat:   -22.9,
		Lon:   -47.1,
	}})

	h := handler.NewHandler(*svc)

	req := httptest.NewRequest(http.MethodGet, "/cep/13.100-000", nil)
	w := httptest.NewRecorder()

	h.Routes().ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("esperando status 200, obtido %d", resp.StatusCode)
	}

	var got app.Address
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("erro ao decodificar resposta JSON: %v", err)
	}

	want := app.Address{
		CEP:       "13100000",
		City:      "Campinas",
		State:     "SP",
		Latitude:  -22.9,
		Longitude: -47.1,
	}

	if got != want {
		t.Errorf("resposta inesperada: \nObtido: %+v\nEsperado: %+v", got, want)
	}
}
