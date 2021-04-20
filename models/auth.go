package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Token JWT claims struct
type Token struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(user_id uint, username string) string {
	tk := &Token{UserID: user_id, Username: username}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString
}

func Login(username, password string) Response {

	user := &User{}

	err := GetDB().Table("users").Where("username = ?", username).First(user).Error
	if err != nil {
		code := 500
		message := "Connection error. Please retry"
		if err == gorm.ErrRecordNotFound {
			code = 404
			message = "Username not found"
		}

		res.Code = code
		res.Message = message
		res.Data = nil
		return res
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		res.Code = 400
		res.Message = "Invalid login credentials. Please try again"
		res.Data = nil
		return res
	}
	//Worked! Logged In

	tokenString := CreateToken(uint(user.UserId), user.Username)

	res.Code = 200
	res.Message = "Logged In"
	res.Data = map[string]interface{}{
		"token": tokenString,
	}
	return res
}
