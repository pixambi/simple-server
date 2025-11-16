package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pixambi/simple-server/internal/config"
)

type Server struct {
	Config       *config.Config
	Router       *chi.Mux
	Repositories *Repositories
	Services     *Services
	Handlers     *Handlers
}

func New(cfg *config.Config) *Server {
	router := chi.NewRouter()

	//Dependencies
	repositories := NewRepositories(cfg)
	services := NewServices(cfg, repositories)
	handlers := NewHandlers(cfg, services)

	s := Server{
		Config:       cfg,
		Router:       router,
		Repositories: repositories,
		Services:     services,
		Handlers:     handlers,
	}

	//Middleware
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	//Register routes
	s.RegisterRoutes()

	return &s
}

func (s *Server) Start(local bool) error {
	addr := ""
	if local {
		addr = fmt.Sprintf("0.0.0.0:%s", s.Config.Port)
	} else {
		addr = fmt.Sprintf("%s:%s", s.Config.Addr, s.Config.Port)
	}

	server := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}
	slog.Info("Server running", "addr", addr)
	return server.ListenAndServe()
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) RegisterRoutes() {
	s.Router.Route("/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Simple Server v1"))
		})

		r.Route("/health", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
		})
		r.Route("/test", func(r chi.Router) {
			r.Get("/test", s.Handlers.Test.HandleTest)
			r.Get("/testService", s.Handlers.Test.HandleService)
			r.Get("/testDomain", s.Handlers.Test.HandleDomain)
			r.Get("/testRepo", s.Handlers.Test.HandleRepo)
		})
	})

	s.Router.Route("/v2", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Simple Server v1"))
		})
	})
}
