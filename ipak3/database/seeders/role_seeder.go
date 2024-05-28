package seeders

import (
	"simpel-wasnaker/ipak3/entity"
	"time"
)

func seedRoles() {
	roles := []entity.Role{
		{
			Name: "Super Admin",
		}, {
			Name: "Admin",
		}, {
			Name: "Pimpinan",
		}, {
			Name: "Staf TU",
		}, {
			Name: "Pemeriksa K3",
		}, {
			Name: "Staf Pemeriksa K3",
		},
	}
	for _, role := range roles {
		role.CreatedAt = time.Now()
		_, err := roleRepo.Create(role)
		if err != nil {
			panic(err)
		}
	}
}
