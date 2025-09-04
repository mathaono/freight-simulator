package cache

import (
	"context"
)

type CEPCache interface {
	GetCEP(ctx context.Context, cep string) (string, bool, error)
	SetCEP(ctx context.Context, cep, json string) error
}
