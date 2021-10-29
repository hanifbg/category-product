package category_test

import (
	"os"
	"testing"
	"time"

	"github.com/hanifbg/category-product/service/category"
	categoryMock "github.com/hanifbg/category-product/service/category/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	id   = 1
	name = "category"
)

var (
	categoryService category.Service
	categoryRepo    categoryMock.Repository

	categoryData      []category.Category
	emptyCategoryData []category.Category
)

func setup() {
	categoryData = []category.Category{
		{
			ID:        id,
			Name:      name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		{
			ID:        id,
			Name:      name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
	}

	emptyCategoryData = []category.Category{}

	categoryService = category.NewService(&categoryRepo)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetCategory(t *testing.T) {
	t.Run("Expect Get List Category", func(t *testing.T) {
		categoryRepo.On("GetCategory").Return(categoryData, nil).Once()
		category, err := categoryService.GetCategory()
		assert.Nil(t, err)
		assert.EqualValues(t, categoryData, category)
	})
}
