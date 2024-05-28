package request

type AuthLoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=20,alphanum"`
	Password string `json:"password" validate:"required"`
}

type AuthLogoutRequest struct {
	ID uint64 `json:"id" validate:"required"`
}