package server

import (
	"github.com/pixambi/simple-server/internal/config"
	"github.com/pixambi/simple-server/internal/service"
)

type Services struct {
	Test *service.TestService
}

func NewServices(cfg *config.Config, repos *Repositories) *Services {
	return &Services{
		Test: service.NewTestService(repos.Test),
	}
}
