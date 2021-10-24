package category

import (
	"time"

	"github.com/hanifbg/category-product/service/category"
	"gorm.io/gorm"
)

type Category struct {
	ID        uint       `gorm:"id;primaryKey;autoIncrement"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
	Name      string     `json:"name"  validate:"required"`
	IsActive  bool       `gorm:"default:true"`
}

func (col *Category) ToCategory() category.Category {
	var category category.Category

	category.ID = col.ID
	category.CreatedAt = col.CreatedAt
	category.UpdatedAt = col.UpdatedAt
	category.DeletedAt = col.DeletedAt
	category.Name = col.Name
	category.IsActive = col.IsActive

	return category
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
func (repo *GormRepository) GetCategory() ([]category.Category, error) {
	var query []Category

	err := repo.DB.Where("deleted_at IS NULL AND is_active = true").Find(&query).Error
	if err != nil {
		return nil, err
	}

	var category []category.Category
	for _, value := range query {
		category = append(category, value.ToCategory())
	}
	return category, nil
}
