package usecase

import (
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterNewUser(payload model.UserCredential)( model.UserCredential, error )
	FindAllUser() ([]model.UserCredential,error)
	FindByUsername(username string) (model.UserCredential,error)
	FindById(id string) (model.UserCredential,error)
	FindByUsernamePassword(username string, password string) (model.UserCredential,error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase)RegisterNewUser(payload model.UserCredential)( model.UserCredential, error ){
	// encrypt password 
	bytes , _ := bcrypt.GenerateFromPassword([]byte(payload.Password),bcrypt.DefaultCost)
	payload.Password = string(bytes)
	err := u.repo.Create(payload)
	if err != nil {
		return model.UserCredential{},fmt.Errorf("Failed to create new user : %s",err.Error())
	}
	user,err := u.FindByUsername(payload.Username)
	if err != nil {
		return model.UserCredential{},fmt.Errorf("Failed get by username : %s",err.Error())
	}
	return user, nil
}

func (u *userUseCase) FindAllUser() ([]model.UserCredential,error) {
	return u.repo.List()
}

func (u *userUseCase) FindByUsername(username string) (model.UserCredential,error) {
	return u.repo.GetByUsername(username)
}

func (u *userUseCase)FindByUsernamePassword(username string, password string) (model.UserCredential,error) {
	return u.repo.GetByUsernamePassword(username,password)
}

func (u *userUseCase) FindById(id string) (model.UserCredential,error) {
	userCred , err := u.repo.GetById(id)
	if err != nil {
		return model.UserCredential{},fmt.Errorf("User Credential is not found : %s",err.Error())
	}
	return userCred,nil
}

func NewUserUseCase (repository repository.UserRepository) UserUseCase {
	return &userUseCase {
		repo : repository,
	}
}