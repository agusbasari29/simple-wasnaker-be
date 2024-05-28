package response

type BusinessResponse struct {
	ID              uint64 `json:"id"`
	BusinessType    string	`json:"business_type"`
	TypeOfCompany   string	`json:"type_of_company"`
	LegalStatus     string	`json:"legal_status"`
}