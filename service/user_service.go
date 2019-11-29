package service

import (
	"encoding/base64"
	"time"

	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserContractRepository
}

func NewUserService(rs repository.UserContractRepository) UserContractService {
	return &UserService{rs}
}

func (u UserService) All() ([]model.User, error) {
	users, err := u.userRepository.All()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserService) Find(id int64) (model.User, error) {
	user, err := u.userRepository.Find(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserService) Store(user model.User) (model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(hash)
	user.Token = generateToken(user.Password)
	user, err = u.userRepository.Store(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserService) Update(user model.User) (model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(hash)
	user.Token = generateToken(user.Password)
	user, err = u.userRepository.Update(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserService) Destroy(id int64) (model.User, error) {
	user, err := u.userRepository.Destroy(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserService) FindBy(field string, value interface{}) (model.User, error) {
	user, err := u.userRepository.FindBy(field, value)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserService) Login(user model.User) (model.User, error) {
	dUser, err := u.userRepository.FindBy("email", user.Email)
	if err != nil {
		return model.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dUser.Password), []byte(user.Password))
	if err != nil {
		return model.User{}, err
	}
	dUser.Token = generateToken(dUser.Password)
	dUser, err = u.userRepository.Update(dUser)
	if err != nil {
		return model.User{}, err
	}
	dUser.Password = ""
	return dUser, nil
}

func (u UserService) Logout(user model.User) (model.User, error) {
	dUser, err := u.userRepository.FindBy("email", user.Email)
	if err != nil {
		return model.User{}, err
	}
	dUser.Token = "NULL"
	dUser, err = u.userRepository.Update(dUser)
	return user, nil
}

func generateToken(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password + time.Now().UTC().Format(time.RFC850)))
}
