package repository

import (
	"simpel-wasnaker/ipak3/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(user entity.User) (entity.User, error)
	GetByID(id uint64) (entity.User, error)
	GetAll() ([]entity.User, error)
	GetAllByCompanyID(companyID uint64) ([]entity.User, error)
	GetAllByRoleID(roleID uint64) ([]entity.User, error)
	GetAllByPhone(phone string) ([]entity.User, error)
	GetAllByEmail(email string) ([]entity.User, error)
	GetByUsername(username string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// GetAllByCompanyID implements UserRepository.
func (repo *userRepository) GetAllByCompanyID(companyID uint64) ([]entity.User, error) {
	panic("unimplemented")
}

// GetAllByEmail implements UserRepository.
func (repo *userRepository) GetAllByEmail(email string) ([]entity.User, error) {
	panic("unimplemented")
}

// GetAllByPhone implements UserRepository.
func (repo *userRepository) GetAllByPhone(phone string) ([]entity.User, error) {
	panic("unimplemented")
}

// GetAllByRoleID implements UserRepository.
func (repo *userRepository) GetAllByRoleID(roleID uint64) ([]entity.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (repo *userRepository) Create(user entity.User) (entity.User, error) {
	err := repo.db.Create(&user).Error
	return user, err
}

func (repo *userRepository) Update(user entity.User) (entity.User, error) {
	err := repo.db.Save(&user).Error
	return user, err
}

func (repo *userRepository) Delete(user entity.User) (entity.User, error) {
	err := repo.db.Delete(&user).Error
	return user, err
}

func (repo *userRepository) GetByID(id uint64) (entity.User, error) {
	var user entity.User
	err := repo.db.First(&user, id).Error
	// err := repo.db.Mod(&user, id).Error
	return user, err
}

func (repo *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := repo.db.Find(&users).Error
	return users, err
}

func (repo *userRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	err := repo.db.Where("email =?", email).First(&user).Error
	return user, err
}

func (repo *userRepository) GetByUsername(username string) (entity.User, error) {
	var user entity.User
	err := repo.db.Model(&entity.User{}).Preload("Contact").Preload("Company").Preload("Role").Where("username = ?", username).Find(&user).Error
	return user, err
}
