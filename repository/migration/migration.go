package migration

import (
	"github.com/hanifbg/category-product/repository/category"
	"github.com/hanifbg/category-product/repository/product"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&category.Category{}, &product.Product{})
}
