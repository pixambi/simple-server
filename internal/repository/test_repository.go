package repository

type TestRepository struct {
}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (r *TestRepository) TestFunction() (string, error) {
	return "Response from repository", nil
}
