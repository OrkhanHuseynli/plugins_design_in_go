package app

import "context"

type Plugin interface {
	Name() string
	Initialize(ctx context.Context) error
	Stop() error
}
