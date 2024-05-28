package repository

import (
	"simpel-wasnaker/ipak3/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role entity.Role) (entity.Role, error)
	Update(role entity.Role) (entity.Role, error)
	Delete(role entity.Role) (entity.Role, error)
	GetByID(id uint64) (entity.Role, error)
	GetAll() ([]entity.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (repo *roleRepository) Create(role entity.Role) (entity.Role, error) {
	err := repo.db.Create(&role).Error
	return role, err
}

func (repo *roleRepository) Update(role entity.Role) (entity.Role, error) {
	err := repo.db.Save(&role).Error
	return role, err
}

func (repo *roleRepository) Delete(role entity.Role) (entity.Role, error) {
	err := repo.db.Delete(&role).Error
	return role, err
}

func (repo *roleRepository) GetByID(id uint64) (entity.Role, error) {
	var role entity.Role
	err := repo.db.First(&role, id).Error
	return role, err
}

func (repo *roleRepository) GetAll() ([]entity.Role, error) {
	var roles []entity.Role
	err := repo.db.Find(&roles).Error
	return roles, err
}
