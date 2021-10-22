package migration

import (
	"github.com/hanifbg/category-product/repository/category"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&category.Category{})
}
