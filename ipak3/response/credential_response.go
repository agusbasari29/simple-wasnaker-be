package response

type CredentialResponse struct {
	Token     string `json:"token"`
	UserID    uint64 `json:"user_id"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	RoleID    string `json:"role_id"`
	Issuer    string `json:"issuer"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiresAt int64  `json:"expired_at"`
}
