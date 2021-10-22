package category

type Service interface {
	GetCategory() ([]Category, error)
}

type Repository interface {
	GetCategory() ([]Category, error)
}
