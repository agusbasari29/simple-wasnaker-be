package entity

import "gorm.io/gorm"

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "ipak3_users"
}

type User struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Username   string
	Password   string
	FullName   string
	EmployeeID string
	ContactID  uint64
	Contact    Contact //`gorm:"foreignKey:ContactID"`
	CompanyID  uint64
	Company    Company //`gorm:"foreignKey:CompanyID"`
	RoleID     uint64
	Role       Role //`gorm:"references:ID"`
}
