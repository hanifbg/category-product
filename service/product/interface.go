package product

type Service interface {
	GetProduct(CategoryId int, param string) ([]Product, error)
	GetDetail(ProductId int) (Product, error)
}

type Repository interface {
	GetProduct(CategoryId int, param string) ([]Product, error)
	GetDetail(ProductId int) (Product, error)
}
