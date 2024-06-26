package entity

import "gorm.io/gorm"

type UserRoles struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Role      string
	RoleName  string
	SectionID uint64
	Sections  Sections `gorm:"foreignKey:SectionID"`
}
