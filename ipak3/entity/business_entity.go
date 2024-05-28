package entity

import "gorm.io/gorm"

func (Business) TableName() string {
	return "ipak3_businesses"
}

type TypeOfCompany string

const (
	GovermentAgency TypeOfCompany = "Pemerintah"
	PrivateCompany  TypeOfCompany = "Swasta"
)

type LegalStatus string

const (
	PT    LegalStatus = "PT"
	PP    LegalStatus = "PP"
	PD    LegalStatus = "PD"
	CV    LegalStatus = "CV"
	Firma LegalStatus = "Firma"
	Balai LegalStatus = "Balai"
)

type Business struct {
	gorm.Model
	ID            uint64 `gorm:"primaryKey;autoIncrement"`
	BusinessType  string
	TypeOfCompany TypeOfCompany
	LegalStatus   LegalStatus
}
