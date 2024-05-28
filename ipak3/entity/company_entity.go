package entity

import (
	"time"

	"gorm.io/gorm"
)

func (Company) TableName() string {
	return "ipak3_companies"
}

type Company struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey;autoIncrement"`
	Name            string
	Address         string
	ContactID       uint64
	Contact         Contact //`gorm:"foreignKey:ContactID"`
	BusinessID      uint64
	Business        Business //`gorm:"foreignKey:BusinessID"`
	RegisteredDate  time.Time
	RegistrationNo  string
	NPWP            string
	EstablishedDate time.Time
}
