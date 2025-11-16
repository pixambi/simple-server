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
	response := h.testService.Ping()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
