package response

import "simpel-wasnaker/ipak3/entity"

type UserDataResponse struct {
	User       interface{} `json:"user_data"`
	Credential interface{} `json:"credential"`
}
type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	// Password   string             `json:"password"`
	FullName   string          `json:"full_name"`
	EmployeeID string          `json:"employee_id"`
	Contact    ContactResponse `json:"contact"`
	Company    CompanyResponse `json:"company"`
	Role       RoleResponse    `json:"role"`
	// Credential CredentialResponse `json:"credential"`
}

func UserResponseFormatter(user entity.User) UserResponse {
	return UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		FullName:   user.FullName,
		EmployeeID: user.EmployeeID,
		Contact: ContactResponse{
			ID:    user.Contact.ID,
			Email: user.Contact.Email,
			Phone: user.Contact.Phone,
			Fax:   user.Contact.Fax,
		},
		Company: CompanyResponse{
			ID:      user.Company.ID,
			Address: user.Company.Address,
			Contact: ContactResponse{
				ID:    user.Company.Contact.ID,
				Email: user.Company.Contact.Email,
				Phone: user.Company.Contact.Phone,
				Fax:   user.Company.Contact.Fax,
			},
			Business: BusinessResponse{
				ID:            user.Company.Business.ID,
				BusinessType:  user.Company.Business.BusinessType,
				TypeOfCompany: string(user.Company.Business.TypeOfCompany),
				LegalStatus:   string(user.Company.Business.LegalStatus),
			},
		},
		Role: RoleResponse{
			ID:   user.Role.ID,
			Name: user.Role.Name,
		},
	}
}

func UserDataResponseFormatter(user interface{}, credential interface{}) UserDataResponse {
	return UserDataResponse{
		User:       user,
		Credential: credential,
	}
}
