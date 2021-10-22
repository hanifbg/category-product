package migration

import (
	"github.com/hanifbg/category-product/repository/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
