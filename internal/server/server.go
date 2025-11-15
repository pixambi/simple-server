package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pixambi/simple-server/internal/config"
)

type Server struct {
	Config *config.Config
	Router *chi.Mux
}

func New(cfg *config.Config) *Server {
	router := chi.NewRouter()

	s := Server{
		Config: cfg,
		Router: router,
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
	log.Printf("Starting server on %s\n", addr)
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
	})

	s.Router.Route("/v2", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Simple Server v1"))
		})
	})
}
