package service

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/andrewesteves/tapagguapi/config"
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService struct
type UserService struct {
	userRepository repository.UserContractRepository
}

// NewUserService new user service
func NewUserService(rs repository.UserContractRepository) UserContractService {
	return &UserService{rs}
}

// All users service
func (u UserService) All() ([]model.User, error) {
	users, err := u.userRepository.All()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Find user service
func (u UserService) Find(id int64) (model.User, error) {
	user, err := u.userRepository.Find(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Store user service
func (u UserService) Store(user model.User) (model.User, error) {
	uFind, _ := u.userRepository.FindBy("email", user.Email)
	if uFind.Name != "" {
		return model.User{}, errors.New(config.LangConfig{}.I18n()["email_taken"])
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(hash)
	user.Token = generateToken(user.Password)
	user.Remember = generateToken(time.Now().UTC().Format(time.RFC850))
	user, err = u.userRepository.Store(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Update user service
func (u UserService) Update(user model.User, newPassword bool) (model.User, error) {
	if newPassword {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err.Error())
		}
		user.Password = string(hash)
		user.Token = generateToken(user.Password)
	}
	user, err := u.userRepository.Update(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Destroy user service
func (u UserService) Destroy(id int64) (model.User, error) {
	user, err := u.userRepository.Destroy(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// FindBy user service
func (u UserService) FindBy(field string, value interface{}) (model.User, error) {
	user, err := u.userRepository.FindBy(field, value)
	if err != nil {
		return model.User{}, errors.New(config.LangConfig{}.I18n()["email_invalid"])
	}
	return user, nil
}

// Login user service
func (u UserService) Login(user model.User) (model.User, error) {
	dUser, err := u.userRepository.FindBy("email", user.Email)
	if err != nil {
		return model.User{}, errors.New(config.LangConfig{}.I18n()["auth_failed"])
	}
	err = bcrypt.CompareHashAndPassword([]byte(dUser.Password), []byte(user.Password))
	if err != nil {
		return model.User{}, errors.New(config.LangConfig{}.I18n()["auth_failed"])
	}
	dUser.Token = generateToken(dUser.Password)
	dUser, err = u.userRepository.Update(dUser)
	if err != nil {
		return model.User{}, err
	}
	return dUser, nil
}

// Logout user service
func (u UserService) Logout(user model.User) (model.User, error) {
	dUser, err := u.userRepository.FindBy("email", user.Email)
	if err != nil {
		return model.User{}, err
	}
	dUser.Token = generateToken(user.Email)
	dUser, err = u.userRepository.Update(dUser)
	return user, nil
}

// FindByArgs args
func (u UserService) FindByArgs(args map[string]interface{}) (model.User, error) {
	dUser, err := u.userRepository.FindByArgs(args)
	if err != nil {
		return model.User{}, errors.New(config.LangConfig{}.I18n()["confirmation_invalid"])
	}
	return dUser, nil
}

func generateToken(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value + time.Now().UTC().Format(time.RFC850)))
}
