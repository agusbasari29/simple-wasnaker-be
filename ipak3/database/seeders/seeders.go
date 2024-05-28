package seeders

import (
	"simpel-wasnaker/ipak3/database"
	"simpel-wasnaker/ipak3/repository"

	"gorm.io/gorm"
)

var (
	db           *gorm.DB = database.SetupDBConnection()
	roleRepo              = repository.NewRoleRepository(db)
	companyRepo           = repository.NewCompanyRepository(db)
	// businessRepo          = repository.NewBusinessRepository(db)
	// contactRepo           = repository.NewContactRepository(db)
	userRepo              = repository.NewUserRepository(db)
)

func Seeders() {
	seedRoles()
	insertCompany()
	insertUsers()
	// seedCompanies()
}
