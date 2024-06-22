package model

import (
	Database "sudoku-api/src/databases"
	"time"
)

type User struct {
	ID            uint `omit`
	GoogleIdToken string
	Email         string
	Nickname      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

type Users struct {
	Users []User
}

func FindUserByEmailOrNicknameOrToken(email string, nickname string, googleIdToken string) *User {
	dbConnection := Database.GetOrCreateConnection()

	var user *User

	dbConnection.Limit(1).Where(&User{Email: email}).Or(&User{Nickname: nickname}).Or(&User{GoogleIdToken: googleIdToken}).Find(&user)

	return user
}

func RegisterUser(user User) User {
	dbConnection := Database.GetOrCreateConnection()

	dbConnection.Create(&user)

	return user
}
