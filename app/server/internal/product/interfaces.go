package product

type ReadRepositoryInterface interface {
	GetProducts() (Products, error)
}
