package repository

import (
	"simpel-wasnaker/ipak3/entity"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(company entity.Company) (entity.Company, error)
	Update(company entity.Company) (entity.Company, error)
	Delete(company entity.Company) (entity.Company, error)
	GetByID(id uint64) (entity.Company, error)
	GetAll() ([]entity.Company, error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *companyRepository {
	return &companyRepository{db}
}

func (repo *companyRepository) Create(company entity.Company) (entity.Company, error) {
	err := repo.db.Create(&company).Error
	return company, err
}

func (repo *companyRepository) Update(company entity.Company) (entity.Company, error) {
	err := repo.db.Save(&company).Error
	return company, err
}

func (repo *companyRepository) Delete(company entity.Company) (entity.Company, error) {
	err := repo.db.Delete(&company).Error
	return company, err
}

func (repo *companyRepository) GetByID(id uint64) (entity.Company, error) {
	var company entity.Company
	err := repo.db.First(&company, id).Error
	return company, err
}

func (repo *companyRepository) GetAll() ([]entity.Company, error) {
	var companies []entity.Company
	err := repo.db.Find(&companies).Error
	return companies, err
}
