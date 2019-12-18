package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
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
	user.Token = u.GenerateToken()
	user.Remember = u.GenerateToken()
	user, err = u.userRepository.Store(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Update user service
func (u UserService) Update(user model.User, newPassword bool) (model.User, error) {
	dUser, err := u.userRepository.Find(user.ID)
	if err != nil {
		return model.User{}, errors.New(config.LangConfig{}.I18n()["auth_failed"])
	}
	if newPassword && user.Password != "" {
		err = bcrypt.CompareHashAndPassword([]byte(dUser.Password), []byte(user.Password))
		if err != nil {
			return model.User{}, errors.New(config.LangConfig{}.I18n()["auth_failed"])
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.RenewPassword), bcrypt.DefaultCost)
		if err != nil {
			return model.User{}, errors.New(err.Error())
		}
		user.Password = string(hash)
		user.Token = u.GenerateToken()
	} else {
		user.Password = dUser.Password
	}
	if user.Name == "" {
		user.Name = dUser.Name
	}
	if user.Token == "" {
		user.Token = dUser.Token
	}
	if user.Remember == "" {
		user.Remember = dUser.Remember
	}
	user.Email = dUser.Email
	user.Active = 1
	user, err = u.userRepository.Update(user)
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
	dUser.Token = u.GenerateToken()
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
	dUser.Token = u.GenerateToken()
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

// UpdateRecover user service
func (u UserService) UpdateRecover(user model.User) (model.User, error) {
	dUser, err := u.userRepository.Find(user.ID)
	if err != nil {
		return model.User{}, errors.New(config.LangConfig{}.I18n()["auth_failed"])
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, errors.New(err.Error())
	}
	user.Remember = u.GenerateToken()
	user.Password = string(hash)
	user.Token = u.GenerateToken()
	user.Name = dUser.Name
	user.Token = dUser.Token
	user.Email = dUser.Email
	user.Active = 1
	user, err = u.userRepository.Update(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GenerateToken tokens
func (u UserService) GenerateToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return base64.StdEncoding.EncodeToString([]byte(uuid + time.Now().UTC().Format(time.RFC850)))
}
