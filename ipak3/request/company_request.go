package request

import "time"

type CreateCompanyRequest struct {
	Name            string    `json:"name" validate:"required,min=5,max=20,alphanum,unique"`
	Address         string    `json:"address" validate:"required"`
	RegisteredDate  time.Time `json:"registered_date" validate:"required"`
	RegistrationNo  string    `json:"registration_no" validate:"required"`
	NPWP            string    `json:"npwp" validate:"required"`
	EstablishedDate time.Time `json:"established_date" validate:"required"`
	Email           string    `json:"email" validate:"required,email"`
	Phone           string    `json:"phone" validate:"required"`
	Fax             string    `json:"fax"`
	BusinessType    string    `json:"business_type" validate:"required"`
	TypeOfCompany   string    `json:"type_of_company" validate:"required"`
	LegalStatus     string    `json:"legal_status" validate:"required"`
}

type CompanyUpdateRequest struct {
	ID              uint64    `json:"id" validate:"required"`
	Name            string    `json:"name"`
	Address         string    `json:"address"`
	RegisteredDate  time.Time `json:"registered_date"`
	RegistrationNo  string    `json:"registration_no"`
	NPWP            string    `json:"npwp"`
	EstablishedDate time.Time `json:"established_date"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Fax             string    `json:"fax"`
	BusinessType    string    `json:"business_type"`
	TypeOfCompany   string    `json:"type_of_company"`
	LegalStatus     string    `json:"legal_status"`
}

type CompanyDeleteRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type CompanyGetRequest struct {
	ID uint64 `json:"id" validate:"required"`
}