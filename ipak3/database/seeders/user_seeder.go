package seeders

import (
	"simpel-wasnaker/ipak3/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func insertUsers() {
	users := []entity.User{
		{
			Username:   "superadmin",
			Password:   password("sangatrahasia"),
			FullName:   "Super Admin",
			EmployeeID: "1234567890",
			RoleID:     1,
			CompanyID:  2,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "admin",
			Password:   password("sangatrahasia"),
			FullName:   "Admin",
			EmployeeID: "1234567890",
			RoleID:     2,
			CompanyID:  2,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "pimpinan01",
			Password:   password("sangatrahasia"),
			FullName:   "<NAME>",
			EmployeeID: "1234567890",
			RoleID:     3,
			CompanyID:  1,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "staftu01",
			Password:   password("sangatrahasia"),
			FullName:   "<NAME>",
			EmployeeID: "1234567890",
			RoleID:     4,
			CompanyID:  1,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "pemeriksa01",
			Password:   password("sangatrahasia"),
			FullName:   "<NAME>",
			EmployeeID: "1234567890",
			RoleID:     5,
			CompanyID:  1,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "stafpemeriksa01",
			Password:   password("sangatrahasia"),
			FullName:   "<NAME>",
			EmployeeID: "1234567890",
			RoleID:     6,
			CompanyID:  1,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "pimpinan02",
			Password:   password("sangatrahasia"),
			FullName:   "<NAME>",
			EmployeeID: "1234567890",
			RoleID:     3,
			CompanyID:  2,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "staftu02",
			Password:   password("sangatrahasia"),
			FullName:   "<NAME>",
			EmployeeID: "1234567890",
			RoleID:     4,
			CompanyID:  3,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "pemeriksa02",
			Password:   password("sangatrahasia"),
			FullName:   "<NAME>",
			EmployeeID: "1234567890",
			RoleID:     5,
			CompanyID:  3,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
		{
			Username:   "stafpemeriksa02",
			Password:   password("sangatrahasia"),
			FullName:   "<NAME>",
			EmployeeID: "1234567890",
			RoleID:     6,
			CompanyID:  3,
			Contact: entity.Contact{
				Email: "<EMAIL>",
				Phone: "081234567890",
				Fax:   "081234567890",
			},
		},
	}

	for _, user := range users {
		user.CreatedAt = time.Now()
		_, err := userRepo.Create(user)
		if err != nil {
			panic(err)
		}
	}
}

func password(pass string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(password)
}

// type SeedUsers struct {
// 	Username   string `faker:"username"`
// 	FullName   string `faker:"name"`
// 	EmployeeID string `faker:"word"`
// }

// func seedUsers(CompanyID uint64, RoleID uint64) {
// 	contact := seedContacts()
// 	seeder := SeedUsers{}
// 	err := faker.FakeData(&seeder)
// 	if err != nil {
// 		panic(err)
// 	}
// 	password, _ := bcrypt.GenerateFromPassword([]byte("sangatrahasia"), bcrypt.DefaultCost)
// 	user := entity.User{}
// 	user.Username = seeder.Username
// 	user.Password = string(password)
// 	user.CreatedAt = time.Now()
// 	user.EmployeeID = seeder.EmployeeID
// 	user.FullName = seeder.FullName
// 	user.RoleID = RoleID
// 	user.CompanyID = CompanyID
// 	user.Contact = contact
// 	_, err = userRepo.Create(user)
// 	if err != nil {
// 		panic(err)
// 	}
// }
