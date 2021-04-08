package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Post struct {
	PostId  int    `json:"post_id" gorm:"primaryKey;column:post_id"`
	UserID  string `json:"user_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content"`
}

func GetAllPost() Response {

	posts := make([]*Post, 0)
	err := GetDB().Table("posts").Where("status = ?", 1).Find(&posts).Error
	if err != nil {
		res.Code = 500
		res.Message = err.Error()
		res.Data = nil
		return res
	}
	if len(posts) < 1 {
		res.Code = 404
		res.Message = "Data not found"
		res.Data = nil
		return res
	}
	res.Code = 200
	res.Message = "Success"
	res.Data = posts
	return res
}

func (user *User) StorePost() Response {
	// check username on db
	temp := &User{}
	GetDB().Table("users").Where("username = ?", user.Username).First(temp)
	if temp.Username != "" {
		res.Code = 400
		res.Message = "Username already registered"
		res.Data = nil
		return res
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	result := GetDB().Create(user)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		res.Code = 500
		res.Message = "Failed to create account"
		res.Data = nil
		return res
		// return u.Message(500, result.Error.Error())
	}

	res.Code = 200
	res.Message = "Account has been created"
	res.Data = map[string]interface{}{
		"user_id": user.UserId,
		"token":   "",
	}
	return res
}

func (user *User) UpdatePost() Response {
	// check username on db
	temp := &User{}
	GetDB().Table("users").First(temp)
	if temp.Username != "" {
		res.Code = 400
		res.Message = "Username already registered"
		res.Data = nil
		return res
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	result := GetDB().Update("name", "hello").Create(user)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		res.Code = 500
		res.Message = "Failed to create account"
		res.Data = nil
		return res
		// return u.Message(500, result.Error.Error())
	}
	res.Code = 200
	res.Message = "Account has been created"
	res.Data = map[string]interface{}{
		"user_id": user.UserId,
		"token":   "",
	}
	return res
}
