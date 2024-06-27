package model

import (
	Database "sudoku-api/src/databases"
	"time"
)

type User struct {
	ID        uint `omit`
	Email     string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Users struct {
	Users []User
}

func FindUserByEmailOrNickname(email string, nickname string) *User {
	dbConnection := Database.GetOrCreateConnection()

	var user *User

	dbConnection.Limit(1).Where(&User{Email: email}).Or(&User{Nickname: nickname}).Find(&user)

	if user.ID == 0 {
		return nil
	}

	return user
}

func RegisterUser(user User) User {
	dbConnection := Database.GetOrCreateConnection()

	dbConnection.Create(&user)

	return user
}

func FindUserByEmail(email string) *User {
	dbConnection := Database.GetOrCreateConnection()

	var user *User

	dbConnection.First(&user, "email = ?", email)

	if user.ID == 0 {
		return nil
	}

	return user
}
