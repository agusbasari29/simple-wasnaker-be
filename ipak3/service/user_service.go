package service

import (
	"log"
	"simpel-wasnaker/ipak3/entity"
	"simpel-wasnaker/ipak3/repository"
	"simpel-wasnaker/ipak3/request"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(req request.UserRegisterRequest) (entity.User, error)
	Update(req request.UserUpdateRequest) (entity.User, error)
	Delete(req request.UserDeleteRequest) error
	Find(id uint64) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindAllByCompanyID(companyID uint64) ([]entity.User, error)
	FindAllByRoleID(roleID uint64) ([]entity.User, error)
	FindByPhone(phone string) (entity.User, error)
}

type userService struct {
	userRepo    repository.UserRepository
	contactRepo repository.ContactRepository
}

// FindAll implements UserService.
func (service *userService) FindAll() ([]entity.User, error) {
	panic("unimplemented")
}

// FindAllByCompanyID implements UserService.
func (service *userService) FindAllByCompanyID(companyID uint64) ([]entity.User, error) {
	panic("unimplemented")
}

// FindAllByRoleID implements UserService.
func (service *userService) FindAllByRoleID(roleID uint64) ([]entity.User, error) {
	panic("unimplemented")
}

// FindByEmail implements UserService.
func (service *userService) FindByEmail(email string) (entity.User, error) {
	panic("unimplemented")
}

// FindByPhone implements UserService.
func (service *userService) FindByPhone(phone string) (entity.User, error) {
	panic("unimplemented")
}

func NewUserService(userRepo repository.UserRepository, contactRepo repository.ContactRepository) *userService {
	return &userService{userRepo, contactRepo}
}

func (service *userService) Create(req request.UserRegisterRequest) (entity.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	user := entity.User{
		CompanyID:  req.CompanyID,
		FullName:   req.FullName,
		EmployeeID: req.EmployeeID,
		Username:   req.Username,
		Password:   string(password),
	}

	contact := entity.Contact{
		Email: req.Email,
		Phone: req.Phone,
		Fax:   req.Fax,
	}

	contact, err = service.contactRepo.Create(contact)
	if err != nil {
		log.Fatalf("Error inserting contact: %v", err)
	}
	user.ContactID = contact.ID
	user, err = service.userRepo.Create(user)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
	return user, err
}

func (service *userService) Update(req request.UserUpdateRequest) (entity.User, error) {
	user, err := service.userRepo.GetByID(req.ID)
	if err != nil {
		log.Printf("Error finding user by id: %v", err)
		return entity.User{}, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating password: %v", err)
		return entity.User{}, err
	}
	user.RoleID = req.RoleID
	user.FullName = req.FullName
	user.EmployeeID = req.EmployeeID
	user.Username = req.Username
	user.Password = string(password)
	user.Contact.Email = req.Email
	user.Contact.Phone = req.Phone
	user.Contact.Fax = req.Fax
	user, err = service.userRepo.Update(user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
	}
	return user, err
}

func (service *userService) Delete(req request.UserDeleteRequest) error {
	user, err := service.userRepo.GetByID(req.ID)
	if err != nil {
		log.Printf("Error finding user by id: %v", err)
		return err
	}

	_, err = service.userRepo.Delete(user)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
	}
	return err
}

func (service *userService) Find(id uint64) (entity.User, error) {
	return service.userRepo.GetByID(id)
}

func (service *userService) FindByUsername(username string) (entity.User, error) {
	return service.userRepo.GetByUsername(username)
}
