package repository

import (
	"simpel-wasnaker/ipak3/entity"

	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact entity.Contact) (entity.Contact, error)
	Update(contact entity.Contact) (entity.Contact, error)
	Delete(contact entity.Contact) (entity.Contact, error)
	GetByID(id uint64) (entity.Contact, error)
	GetAll() ([]entity.Contact, error)
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *contactRepository {
	return &contactRepository{db}
}

func (repo *contactRepository) Create(contact entity.Contact) (entity.Contact, error) {
	err := repo.db.Create(&contact).Error
	return contact, err
}

func (repo *contactRepository) Update(contact entity.Contact) (entity.Contact, error) {
	err := repo.db.Save(&contact).Error
	return contact, err
}

func (repo *contactRepository) Delete(contact entity.Contact) (entity.Contact, error) {
	err := repo.db.Delete(&contact).Error
	return contact, err
}

func (repo *contactRepository) GetByID(id uint64) (entity.Contact, error) {
	var contact entity.Contact
	err := repo.db.First(&contact, id).Error
	return contact, err
}

func (repo *contactRepository) GetAll() ([]entity.Contact, error) {
	var contacts []entity.Contact
	err := repo.db.Find(&contacts).Error
	return contacts, err
}