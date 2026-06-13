package repository

import (
	"gorm.io/gorm"
	"order-service/internal/entity"
)

type CityRepository struct {
	db *gorm.DB
}

func NewCityRepository(
	db *gorm.DB,
) *CityRepository {

	return &CityRepository{
		db: db,
	}
}

func (r *CityRepository) GetByID(
	id int,
) (*entity.City, error) {

	var city entity.City

	err := r.db.
		Where("city_id = ?", id).
		First(&city).Error

	return &city, err
}