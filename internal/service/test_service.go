package service

import testdomain "github.com/pixambi/simple-server/internal/domain/test_domain"

type TestService struct {
}

func NewTestService() *TestService {
	return &TestService{}
}

func (s *TestService) Ping() string {
	domain_response := testdomain.Ping()
	response := "Service response and domain logic:" + domain_response
	return response
}
