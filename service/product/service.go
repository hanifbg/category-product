package product

import serv "github.com/hanifbg/category-product/service"

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

func (s *service) GetDetail(ProductId int) (Product, error) {

	product, err := s.repository.GetDetail(ProductId)
	if err != nil {
		return Product{}, err
	}

	if product.ID == 0 {
		return Product{}, serv.ErrNotFound
	}

	return product, nil
}
