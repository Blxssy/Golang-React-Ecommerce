package service

import (
	"errors"
	"time"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/avatar"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/token"
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterUser(email string, password string) (*models.User, string, string, error)
	LoginUser(email string, password string) (*models.User, string, string, error)

	FindById(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAllUsers() (*[]models.User, error)
	UpdateUser(user models.User) error
	CreateUser(user *models.User) error
}

type userService struct {
	container container.Container
}

func NewUserService(container container.Container) UserService {
	return &userService{
		container: container,
	}
}

func (u *userService) RegisterUser(email string, password string) (*models.User, string, string, error) {
	result, err := u.FindByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", "", err
	}

	if result != nil {
		return nil, "", "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", "", err
	}

	username := faker.FirstName()
	user := &models.User{
		Email:      email,
		Username:   username,
		PassHash:   string(hashedPassword),
		AvatarPath: avatar.GenerateRandomAvatar(username),
		Phone:      faker.Phonenumber(),
	}

	if err := u.CreateUser(user); err != nil {
		return nil, "", "", err
	}

	accessToken, err := token.NewToken(user.ID, time.Hour*24)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := token.NewToken(user.ID, time.Hour*24*7)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (u *userService) LoginUser(email string, password string) (*models.User, string, string, error) {
	user, err := u.FindByEmail(email)
	if err != nil {
		return nil, "", "", err
	}
	//logger := u.container.GetLogger()
	//logger.Info("LoginUser", slog.Any("user", user))
	//logger.Info("LoginUser", slog.Any("pas", password))
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password)); err != nil {
		return nil, "", "", err
	}

	accessToken, err := token.NewToken(user.ID, time.Hour*24)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := token.NewToken(user.ID, time.Hour*24*7)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (u *userService) FindById(id string) (*models.User, error) {
	var user models.User
	if err := u.container.GetRepository().First(&user, id).Error; err != nil {
		return nil, nil
	}
	return &user, nil
}

func (u *userService) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if result := u.container.GetRepository().Where("email = ?", email).First(&user).Error; result != nil {
		// u.container.GetLogger().Info("res", slog.Any("res", result))
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
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

func (u *userService) CreateUser(user *models.User) error {
	s := u.container.GetRepository()
	var err error

	if txerr := s.Transaction(func(tx storage.Storage) error {
		_, err = txCreateUser(tx, user)
		return err
	}); txerr != nil {
		u.container.GetLogger().Error(txerr.Error())
		return nil
	}

	return nil
}

func txCreateUser(txstorage storage.Storage, user *models.User) (*models.User, error) {
	var result *models.User
	var err error

	if result, err = user.Create(txstorage); err != nil {
		return nil, err
	}

	return result, nil
}
