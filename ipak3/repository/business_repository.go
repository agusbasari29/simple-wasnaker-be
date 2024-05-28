package repository

import (
	"simpel-wasnaker/ipak3/entity"

	"gorm.io/gorm"
)

type BusinessRepository interface {
	Create(business entity.Business) (entity.Business, error)
	Update(business entity.Business) (entity.Business, error)
	Delete(business entity.Business) (entity.Business, error)
	GetByID(id uint64) (entity.Business, error)
	GetAll() ([]entity.Business, error)
}

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) *businessRepository {
	return &businessRepository{db}
}

func (repo *businessRepository) Create(business entity.Business) (entity.Business, error) {
	err := repo.db.Create(&business).Error
	return business, err
}

func (repo *businessRepository) Update(business entity.Business) (entity.Business, error) {
	err := repo.db.Save(&business).Error
	return business, err
}

func (repo *businessRepository) Delete(business entity.Business) (entity.Business, error) {
	err := repo.db.Delete(&business).Error
	return business, err
}

func (repo *businessRepository) GetByID(id uint64) (entity.Business, error) {
	var business entity.Business
	err := repo.db.First(&business, id).Error
	return business, err
}

func (repo *businessRepository) GetAll() ([]entity.Business, error) {
	var businesses []entity.Business
	err := repo.db.Find(&businesses).Error
	return businesses, err
}
