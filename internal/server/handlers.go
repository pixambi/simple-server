package server

import (
	"github.com/pixambi/simple-server/internal/config"
	"github.com/pixambi/simple-server/internal/handler"
)

type Handlers struct {
	Test *handler.TestHandler
}

func NewHandlers(cfg *config.Config, services *Services) *Handlers {
	return &Handlers{
		Test: handler.NewTestHandler(cfg, services.Test),
	}
}
