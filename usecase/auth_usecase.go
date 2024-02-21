package usecase

import (
	"fmt"
	"enigmacamp.com/enigma-laundry-apps/utils/security"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
}

type authUseCase struct {
	userUc UserUseCase
}

func (a *authUseCase) Login(username string, password string) (string, error) {
	user, err := a.userUc.FindByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("Invalid username or password")
	}
	// setelah login berhasil, maka kita berikan token 
	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "",fmt.Errorf("Failed to Generate Token : %s ",err.Error())
	}
	return token,nil
}

func NewAuthUseCase(userUseCase UserUseCase) AuthUseCase {
	return &authUseCase{
		userUc : userUseCase,
	}
}