package category

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

func (s *service) GetCategory() ([]Category, error) {

	category, err := s.repository.GetCategory()
	if err != nil {
		return []Category{}, err
	}

	return category, nil
}
