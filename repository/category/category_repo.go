package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string `json:"name"  validate:"required"`
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
func (repo *GormRepository) GetCategory() error {

	return nil
}
