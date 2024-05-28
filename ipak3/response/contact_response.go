package response

type ContactResponse struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Fax      string `json:"fax"`
}