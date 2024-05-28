package response

import "time"

type CompanyResponse struct {
	ID              uint64           `json:"id"`
	Name            string           `json:"name"`
	Address         string           `json:"address"`
	Contact         ContactResponse  `json:"contact"`
	Business        BusinessResponse `json:"business"`
	RegisteredDate  time.Time        `json:"registered_date"`
	RegistrationNo  string           `json:"registration_no"`
	NPWP            string           `json:"npwp"`
	EstablishedDate time.Time        `json:"established_date"`
	// ContactID       uint64
	// BusinessID      uint64
}
