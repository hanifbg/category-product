package product

import (
	"time"

	"github.com/hanifbg/category-product/service/product"
	"gorm.io/gorm"
)

type Product struct {
	ID         uint       `gorm:"id;primaryKey;autoIncrement"`
	CategoryID uint       `gorm:"category_id"`
	Name       string     `gorm:"name"`
	Price      int        `gorm:"price"`
	Stock      int        `gorm:"stock"`
	Image      string     `gorm:"image"`
	Detail     string     `gorm:"detail"`
	CreatedAt  time.Time  `gorm:"created_at"`
	UpdatedAt  time.Time  `gorm:"updated_at"`
	DeletedAt  *time.Time `gorm:"deleted_at"`
}

func (col *Product) ToProduct() product.Product {
	var product product.Product

	product.ID = col.ID
	product.CategoryID = col.CategoryID
	product.Name = col.Name
	product.Price = col.Price
	product.Stock = col.Stock
	product.Image = col.Image
	product.Detail = col.Detail
	product.CreatedAt = col.CreatedAt
	product.UpdatedAt = col.UpdatedAt
	product.DeletedAt = col.DeletedAt

	return product
}

type GormRepository struct {
	DB *gorm.DB
}

func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//InsertUser Insert new User into storage
func (repo *GormRepository) GetProduct(CategoryId int) ([]product.Product, error) {
	var query []Product

	err := repo.DB.Where("category_id = ?", CategoryId).Find(&query).Error
	if err != nil {
		return nil, err
	}

	var category []product.Product
	for _, value := range query {
		category = append(category, value.ToProduct())
	}
	return category, nil
}
