package service

type TestRepository interface {
	TestFunction() (string, error)
}
