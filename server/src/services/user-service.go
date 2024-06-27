package services

import (
	"errors"
	Model "sudoku-api/src/models"
	"time"
)

type RegisterData struct {
	Email    string
	Nickname string
}

func RegisterUser(registerData RegisterData) (string, error) {
	foundUser := Model.FindUserByEmailOrNickname(registerData.Email, registerData.Nickname)

	if foundUser != nil {
		return "error", errors.New("user with provided data already exists")
	}

	newUser := Model.User{Email: registerData.Email, Nickname: registerData.Nickname, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}

	Model.RegisterUser(newUser)

	return "user registered", nil
}

func FindUserByEmail(email string) (*Model.User, error) {
	foundUser := Model.FindUserByEmail(email)

	if foundUser == nil {
		return nil, errors.New("user not found")
	}

	return foundUser, nil
}
