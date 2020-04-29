package micro

type Service struct {
	Endpoints []*Endpoint
}

func (s *Service) Start() {
}
func (s *Service) Stop() {
}
func (s *Service) Health() {
}
