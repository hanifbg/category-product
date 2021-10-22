package category

type Service interface {
	GetCategory() error
}

type Repository interface {
	GetCategory() error
}
