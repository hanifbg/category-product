package product

type Service interface {
	GetProduct(CategoryId int) ([]Product, error)
}

type Repository interface {
	GetProduct(CategoryId int) ([]Product, error)
}
