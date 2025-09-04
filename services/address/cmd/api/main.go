package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/mathaono/freight-simulator/pkg/logger"
	"github.com/mathaono/freight-simulator/services/address/internal/app"
	"github.com/mathaono/freight-simulator/services/address/internal/cache"
	"github.com/mathaono/freight-simulator/services/address/internal/cep"
	"github.com/mathaono/freight-simulator/services/address/internal/handler"
	"go.uber.org/zap"
)

func main() {
	// Inicializa o logger baseado na variável de ambiente env
	env := os.Getenv("ENV")
	if err := logger.Init(env); err != nil {
		panic("Não foi possível iniciar o logger: " + err.Error())
	}
	defer logger.Sync()

	// Cria instância Redis a partir das variáveis de ambiente
	redisCache, err := cache.NewRedisFromEnv()
	if err != nil {
		logger.L().Fatal("erro ao conectar redis", zap.Error(err))
	}

	// Inicialmente usa MockProvider como CEP (futuramente será trocado por valores reais)
	provider := cep.MockProvider{}

	// Inicializa o Service com cache + provider
	service := app.NewService(redisCache, provider)

	// Inicializa handler HTTP
	h := handler.NewHandler(*service)

	// Cria e configura o roteador
	r := chi.NewRouter()
	r.Mount("/", h.Routes())

	addr := ":8080"
	logger.L().Info("address-svc iniciado", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, r); err != nil {
		logger.L().Fatal("erro ao iniciar servidor", zap.Error(err))
	}

	// Espera sinal de término
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	logger.L().Info("encerrando serviço address-svc")
}
