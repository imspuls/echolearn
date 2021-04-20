package models

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserId   int    `json:"user_id" gorm:"primaryKey;column:user_id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type APIUser struct {
	UserId   int    `json:"user_id" gorm:"primaryKey;column:user_id"`
	Username string `json:"username" validate:"required"`
}

func GetAllUser() ([]*APIUser, error) {
	users := make([]*APIUser, 0)
	err := GetDB().Table("users").Where("status = ?", 1).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (user *User) StoreUser() (int, error) {
	// check username on db
	temp := &User{}
	GetDB().Table("users").Where("username = ?", user.Username).First(temp)
	if temp.Username != "" {
		return 0, errors.New("username already registered")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	result := GetDB().Create(user)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return 0, errors.New("failed to create account")
	}

	return user.UserId, nil
}

func (user *APIUser) UpdateUser() error {
	// check username on db
	if message, exist := checkUser(user.UserId); !exist {
		// error handling...
		return errors.New(message)
	}

	result := GetDB().Table("users").Where("user_id = ?", user.UserId).Update("username", user.Username)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return errors.New("failed to update account")
	}

	return nil
}

func (user *User) DeleteUser() error {
	// check username on db
	if message, exist := checkUser(user.UserId); !exist {
		// error handling...
		return errors.New(message)
	}

	if err := GetDB().Where("user_id = ?", user.UserId).Delete(user).Error; err != nil {
		fmt.Println(err.Error())
		return errors.New("failed to delete account")
	}

	return nil
}

func checkUser(userId int) (string, bool) {
	temp := &User{}
	if err := GetDB().Where("user_id = ?", userId).First(temp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "User not found", false
		}
		return "Connection error. Please retry", false
	}

	return "", true
}
