package models

import (
	"errors"
	"log"
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

func Login(username, password string) (string, error) {

	user := &User{}

	err := GetDB().Table("users").Where("username = ?", username).First(user).Error
	if err != nil {
		log.Println(err)
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("username not found")
		}

		return "", errors.New("connection error. please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return "", errors.New("invalid login credentials. please try again")
	}

	//Worked! Logged In
	tokenString := CreateToken(uint(user.UserId), user.Username)

	return tokenString, nil
}
