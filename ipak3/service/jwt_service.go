package service

import (
	"fmt"
	"os"
	"simpel-wasnaker/ipak3/entity"
	"simpel-wasnaker/ipak3/response"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTServices interface {
	GenerateToken(user entity.User) response.CredentialResponse
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secret string
	issuer string
}

type jwtCustomClaim struct {
	UserID   uint64 `json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	RoleID     string `json:"role_id"`
	jwt.StandardClaims
}

func NewJWTService() JWTServices {
	return &jwtService{
		issuer: "simpel-wasnaker-system",
		secret: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET KEY")
	if secretKey == "" {
		secretKey = "supersecret"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(user entity.User) response.CredentialResponse {
	claims := &jwtCustomClaim{}
	claims.UserID = user.ID
	claims.FullName = user.FullName
	claims.Username = user.Username
	claims.Email = user.Contact.Email
	claims.RoleID = fmt.Sprintf("%v", user.RoleID)
	claims.ExpiresAt = time.Now().AddDate(1, 0, 0).Unix()
	claims.Issuer = j.issuer
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		panic(err)
	}
	return response.CredentialResponse{
		Token:     signedToken,
		UserID:    claims.UserID,
		FullName:  claims.FullName,
		Email:     claims.Email,
		Username:  claims.Username,
		RoleID:    claims.RoleID,
		Issuer:    claims.Issuer,
		IssuedAt:  claims.IssuedAt,
		ExpiresAt: claims.ExpiresAt,
	}
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwtString := strings.Split(token, "Bearer ")[1]
	return jwt.Parse(jwtString, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}
