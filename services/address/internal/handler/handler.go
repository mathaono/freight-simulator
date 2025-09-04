package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	chi "github.com/go-chi/chi/v5"
	"github.com/mathaono/freight-simulator/pkg/logger"
	"github.com/mathaono/freight-simulator/services/address/internal/app"
	"go.uber.org/zap"
)

type Handler struct {
	svc app.Service
}

func NewHandler(svc app.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/cep/{cep}", h.GetCEPHandler)
	return r
}

func (h *Handler) GetCEPHandler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	cep = strings.TrimSpace(cep)

	result, err := h.svc.FindCEP(r.Context(), cep)
	if err != nil {
		logger.L().Error("erro ao buscar CEP", zap.String("cep", cep), zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
