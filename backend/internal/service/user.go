package service

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(email string, password string) error
	LoginUser(email string, password string) error

	FindById(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAllUsers() (*[]models.User, error)
	UpdateUser(user models.User) error
	CreateUser(user *models.User) (*models.User, error)
}

type userService struct {
	container container.Container
}

func NewUserService(container container.Container) UserService {
	return &userService{
		container: container,
	}
}

func (u *userService) RegisterUser(email string, password string) error {

	return nil
}

func (u *userService) LoginUser(email string, password string) error {
	return nil
}

func (u *userService) FindById(id string) (*models.User, error) {
	return nil, nil
}

func (u *userService) FindByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (u *userService) FindAllUsers() (*[]models.User, error) {
	var users []models.User
	if err := u.container.GetRepository().Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *userService) UpdateUser(user models.User) error {
	return nil
}

func (u *userService) ValidatePassword(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func (u *userService) CreateUser(user *models.User) (*models.User, error) {
	s := u.container.GetRepository()
	var result *models.User
	var err error

	if txerr := s.Transaction(func(tx storage.Storage) error {
		result, err = txCreateUser(tx, user)
		return err
	}); txerr != nil {
		u.container.GetLogger().Error(txerr.Error())
		return nil, txerr
	}

	return result, nil
}

func txCreateUser(txstorage storage.Storage, user *models.User) (*models.User, error) {
	var result *models.User
	var err error

	if result, err = user.Create(txstorage); err != nil {
		return nil, err
	}

	return result, nil
}
