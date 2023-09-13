package product

import (
	"challenge/database"
)

type Service struct {
	productsRepo ReadRepositoryInterface
}

func NewService() *Service {
	db := database.Connect()
	repo := NewRepository(db)

	return &Service{productsRepo: repo}
}

func (s *Service) GetProducts() (Products, error) {
	products, err := s.productsRepo.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}
