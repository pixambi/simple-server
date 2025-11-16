package server

import (
	"github.com/pixambi/simple-server/internal/config"
	"github.com/pixambi/simple-server/internal/repository"
	"github.com/pixambi/simple-server/internal/service"
)

type Repositories struct {
	Test service.TestRepository
}

func NewRepositories(cfg *config.Config) *Repositories {
	return &Repositories{
		Test: repository.NewTestRepository(),
	}
}
