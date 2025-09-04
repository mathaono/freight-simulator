package logger

import (
	"strings"

	"go.uber.org/zap"
)

var log *zap.Logger

// Configura o logger global
func Init(env string) error {
	var err error

	switch strings.ToLower(env) {
	case "prod":
		log, err = zap.NewProduction()

	default:
		log, err = zap.NewDevelopment()
	}

	if err != nil {
		return err
	}

	zap.ReplaceGlobals(log)
	return nil
}

// Limpa os buffers (deve ser chamado no shutdown da aplicação)
func Sync() {
	_ = log.Sync()
}

// Retorna a instância atual do logger
func L() *zap.Logger {
	if log == nil {
		_ = Init("dev")
	}
	return log
}
