package entity

import "gorm.io/gorm"

func (Contact) TableName() string {
	return "ipak3_contacts"
}

type Contact struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey;autoIncrement"`
	Email string
	Phone string
	Fax   string
}
