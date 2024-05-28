package service

import (
	"simpel-wasnaker/ipak3/entity"
	"simpel-wasnaker/ipak3/repository"
	"simpel-wasnaker/ipak3/request"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	// Register(req request.AuthRegisterRequest) (entity.User, error)
	Login(req request.AuthLoginRequest) (entity.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *authService {
	return &authService{userRepo}
}

func (service *authService) Login(req request.AuthLoginRequest) (entity.User, error) {
	user, err := service.userRepo.GetByUsername(req.Username)
	if err != nil {
		return user, err
	}
	if !checkPasswordHash(req.Password, user.Password) {
		return user, err
	}
	return user, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
