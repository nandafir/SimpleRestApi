package apps

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s Service) ViewTest() (string, error) {

	return "return", nil
}
