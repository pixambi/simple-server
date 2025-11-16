package service

import (
	testdomain "github.com/pixambi/simple-server/internal/domain/test_domain"
)

type TestService struct {
	testRepo TestRepository
}

func NewTestService(testRepo TestRepository) *TestService {
	return &TestService{
		testRepo: testRepo,
	}
}

func (s *TestService) TestService() string {
	response := "Service response"
	return response
}

func (s *TestService) TestDomain() string {
	domain_response := testdomain.Ping()
	response := "Domain response:" + domain_response
	return response
}

func (s *TestService) TestRepo() string {
	repo_response, err := s.testRepo.TestFunction()
	if err != nil {
		return "Error: " + err.Error()
	}
	response := "Repository response: " + repo_response
	return response
}
