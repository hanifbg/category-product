package product

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

type CreateCategoryData struct {
	Name string `validate:"required"`
}

func (s *service) GetProduct(CategoryId int, param string) ([]Product, error) {

	category, err := s.repository.GetProduct(CategoryId, param)
	if err != nil {
		return []Product{}, err
	}

	return category, nil
}