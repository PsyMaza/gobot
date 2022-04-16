package product

type Service struct {
	Products map[string]Product
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}
