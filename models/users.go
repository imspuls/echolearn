package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   int    `json:"user_id" gorm:"column:user_id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type APIUser struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username" validate:"required"`
}

func GetAllUser() Response {

	users := make([]*APIUser, 0)
	err := GetDB().Table("users").Where("status = ?", 1).Find(&users).Error
	if err != nil {
		res.Code = 500
		res.Message = err.Error()
		res.Data = nil
		return res
	}
	if len(users) < 1 {
		res.Code = 404
		res.Message = "Data not found"
		res.Data = nil
		return res
	}
	res.Code = 200
	res.Message = "Success"
	res.Data = users
	return res
}

func (user *User) StoreUser() Response {
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

	//Create new JWT token for the newly registered account
	// tk := &Token{UserID: account.ID, Username: account.Email}
	// token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	// tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	// account.Token = tokenString

	// response := u.Message(200, "Account has been created")
	// response["data"] = map[string]interface{}{
	// 	"user_id": user.Id,
	// 	"token":   "",
	// }
	res.Code = 200
	res.Message = "Account has been created"
	res.Data = map[string]interface{}{
		"user_id": user.UserId,
		"token":   "",
	}
	return res
}

func (user *APIUser) UpdateUser() Response {
	// check username on db
	if message, exist := checkUser(user.UserId); !exist {
		// error handling...
		res.Code = 404
		res.Message = message
		res.Data = nil
		return res
	}

	result := GetDB().Table("users").Where("user_id = ?", user.UserId).Update("username", user.Username)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		res.Code = 500
		res.Message = "Failed to update account"
		res.Data = nil
		return res
	}

	res.Code = 200
	res.Message = "Account has been updated"
	res.Data = map[string]interface{}{
		"user_id": user.UserId,
	}
	return res
}

func (user *User) DeleteUser() Response {
	// check username on db
	if message, exist := checkUser(user.UserId); !exist {
		// error handling...
		res.Code = 404
		res.Message = message
		res.Data = nil
		return res
	}

	if err := GetDB().Where("user_id = ?", user.UserId).Delete(user).Error; err != nil {
		fmt.Println(err.Error())
		res.Code = 500
		res.Message = "Failed to delete account"
		res.Data = nil
		return res
	}

	res.Code = 200
	res.Message = "Account has been deleted"
	res.Data = map[string]interface{}{
		"user_id": user.UserId,
	}
	return res
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
