package handler

import (
	"net/http"

	"github.com/pixambi/simple-server/internal/config"
	"github.com/pixambi/simple-server/internal/service"
)

type TestHandler struct {
	testService *service.TestService
}

func NewTestHandler(cfg *config.Config, testService *service.TestService) *TestHandler {
	return &TestHandler{
		testService: testService,
	}
}

func (h *TestHandler) HandleTest(w http.ResponseWriter, r *http.Request) {
	response := "Handler response"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (h *TestHandler) HandleService(w http.ResponseWriter, r *http.Request) {
	response := h.testService.TestService()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (h *TestHandler) HandleDomain(w http.ResponseWriter, r *http.Request) {
	response := h.testService.TestDomain()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (h *TestHandler) HandleRepo(w http.ResponseWriter, r *http.Request) {
	response := h.testService.TestRepo()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
