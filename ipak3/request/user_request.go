package request

type GetUserByIdRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type UserRegisterRequest struct {
	CompanyID  uint64 `json:"company_id" validate:"required"` //TODO: get company id from req
	FullName   string `json:"full_name" validate:"required"`
	EmployeeID string `json:"employee_id" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required"`
	Fax        string `json:"fax"`
	Username   string `json:"username" validate:"required,min=5,max=20,alphanum,unique"`
	Password   string `json:"password" validate:"required"`
	RoleID     uint64 `json:"role_id" validate:"required"`
}

type UserUpdateRequest struct {
	ID         uint64 `json:"id" validate:"required"`
	FullName   string `json:"full_name"`
	EmployeeID string `json:"employee_id"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Fax        string `json:"fax"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleID     uint64 `json:"role_id"`
}

type UserDeleteRequest struct {
	ID uint64 `json:"id" validate:"required"`
}

type UserGetRequest struct {
	ID uint64 `json:"id" validate:"required"`
}
