package product_test

import (
	"os"
	"testing"
	"time"

	serv "github.com/hanifbg/category-product/service"
	"github.com/hanifbg/category-product/service/product"
	productMock "github.com/hanifbg/category-product/service/product/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id         = 1
	categoryId = 1
	name       = "category"
	price      = 100
	stock      = 10
	image      = "images"
	detail     = "detail"
)

var (
	productService product.Service
	productRepo    productMock.Repository

	productData   []product.Product
	productDetail product.Product
)

func setup() {
	productData = []product.Product{
		{
			ID:         id,
			CategoryID: categoryId,
			Name:       name,
			Price:      price,
			Stock:      stock,
			Image:      image,
			Detail:     detail,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			DeletedAt:  nil,
		},
		{
			ID:         id,
			CategoryID: categoryId,
			Name:       name,
			Price:      price,
			Stock:      stock,
			Image:      image,
			Detail:     detail,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			DeletedAt:  nil,
		},
	}

	productDetail = product.Product{
		ID:         id,
		CategoryID: categoryId,
		Name:       name,
		Price:      price,
		Stock:      stock,
		Image:      image,
		Detail:     detail,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
	}

	productService = product.NewService(&productRepo)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetProduct(t *testing.T) {
	t.Run("Expect Get List Product", func(t *testing.T) {
		productRepo.On("GetProduct", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(productData, nil).Once()
		product, err := productService.GetProduct(categoryId, "string")
		assert.Nil(t, err)
		assert.EqualValues(t, productData, product)
	})
}

func TestGetDetail(t *testing.T) {
	t.Run("Expect Get Product Detail", func(t *testing.T) {
		productRepo.On("GetDetail", mock.AnythingOfType("int")).Return(productDetail, nil).Once()
		readDetail, err := productService.GetDetail(id)
		assert.Nil(t, err)
		assert.EqualValues(t, productDetail, readDetail)
	})
	t.Run("Expect Error Not Found", func(t *testing.T) {
		productRepo.On("GetDetail", mock.AnythingOfType("int")).Return(product.Product{}, nil).Once()
		readDetail, err := productService.GetDetail(0)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrNotFound)
		assert.EqualValues(t, product.Product{}, readDetail)
	})
}
