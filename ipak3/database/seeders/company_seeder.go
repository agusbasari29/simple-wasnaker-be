package seeders

import (
	"simpel-wasnaker/ipak3/entity"
	"strconv"
	"time"
)

// type SeedCompany struct {
// 	Name            string `faker:"name"`
// 	Address         string `faker:"sentence"`
// 	RegisteredDate  int64  `faker:"unix_time"`
// 	RegistrationNo  string `faker:"sentence"`
// 	NPWP            string `faker:"sentence"`
// 	EstablishedDate int64  `faker:"unix_time"`
// }

func insertCompany() {
	companies := []entity.Company{
		{
			Name:    "DINAS TENAGA KERJA DAN TRANSMIGRASI PROVINSI JAWA BARAT",
			Address: "Jl. Modar Mandir No. 22, Kota Bandung, Jawa Barat 40111, Indonesia",
			Contact: entity.Contact{
				Email: "balaik3@dinsnaker.jabarprov.go.id",
				Phone: "021-12345678",
				Fax:   "021-12345678",
			},
			Business: entity.Business{
				BusinessType:  "Pelayanan Publik",
				TypeOfCompany: entity.GovermentAgency,
				LegalStatus:   entity.Balai,
			},
			RegisteredDate:  convertTime(1613459200),
			RegistrationNo:  "010000000000000",
			NPWP:            "010000000000000",
			EstablishedDate: convertTime(1613459200),
		},
		{
			Name:    "PT. ASTERIA RIKSA INDONESIA",
			Address: "Jl. Kaliurang No. 10, Kota Bandung, Jawa Barat 40111, Indonesia",
			Contact: entity.Contact{
				Email: "cs@asteria.com",
				Phone: "022-12345678",
				Fax:   "022-12345678",
			},
			Business: entity.Business{
				BusinessType:  "Penyedia jasa teknologi informasi", 
				TypeOfCompany: entity.PrivateCompany,
				LegalStatus:   entity.PT,
			},
			RegisteredDate:  convertTime(1613459200),
			RegistrationNo:  "010000000000000",
			NPWP:            "010000000000000",
			EstablishedDate: convertTime(1613459200),
		},
		{
			Name:    "PT. <NAME>",
			Address: "Jl. <NAME> No. 10, Kota Bandung, Jawa Barat 40111, Indonesia",
			Contact: entity.Contact{
				Email: "cs@company.com",
				Phone: "022-12345678",
				Fax:   "022-12345678",
			},
			Business: entity.Business{
				BusinessType:  "Riksa Uji Alat Industri Swasta",
				TypeOfCompany: entity.PrivateCompany,
				LegalStatus:   entity.PT,
			},
			RegisteredDate:  convertTime(1613459200),
			RegistrationNo:  "010000000000000",
			NPWP:            "010000000000000",
			EstablishedDate: convertTime(1613459200),
		},
	}

	for _, company := range companies {
		_, err := companyRepo.Create(company)
		if err != nil {
			panic(err)
		}
	}
}

// func seedCompanies() {
// 	seederCompany := SeedCompany{}
// 	company := entity.Company{}
// 	j := uint64(1)
// 	for i := 0; i < 2; i++ {
// 		contact := seedContacts()
// 		business := seedBusinesses()
// 		company.Contact = contact
// 		company.Business = business
// 		err := faker.FakeData(&seederCompany)
// 		if err != nil {
// 			panic(err)
// 		}
// 		company.Name = seederCompany.Name
// 		company.Address = seederCompany.Address
// 		company.EstablishedDate = convertTime(seederCompany.EstablishedDate)
// 		company.RegisteredDate = convertTime(seederCompany.RegisteredDate)
// 		company.RegistrationNo = seederCompany.RegistrationNo
// 		company.NPWP = seederCompany.NPWP
// 		company.CreatedAt = time.Now()
// 		company, err = companyRepo.Create(company)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if j < 3 {
// 			for i := 0; i < 3; i++ {
// 				seedUsers(company.ID, j)
// 				j++
// 			}
// 		} else {
// 			for i := 0; i < 5; i++ {
// 				seedUsers(company.ID, j)
// 				j++
// 			}
// 		}
// 	}
// }

func convertTime(unix int64) time.Time {
	i, err := strconv.ParseInt(strconv.Itoa(int(unix)), 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0)
}
