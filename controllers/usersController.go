package controllers

import (
	"echolearn/models"
	"fmt"
	"net/http"
	"strconv"

	u "echolearn/utils"

	"github.com/labstack/echo"
)

func GetAllUser(c echo.Context) error {
	resp := models.GetAllUser()
	return c.JSON(http.StatusOK, resp)
}

func CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusOK, u.Message(400, err.Error()))
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, u.Message(500, err.Error()))
	}
	resp := user.StoreUser() //Create account
	return c.JSON(http.StatusOK, resp)
}

func UpdateUser(c echo.Context) error {
	user_id := c.Param("id")
	user := &models.APIUser{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusOK, u.Message(400, err.Error()))
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, u.Message(500, err.Error()))
	}

	convertedId, _ := strconv.Atoi(user_id)
	user.UserId = convertedId
	resp := user.UpdateUser() //Update account
	return c.JSON(http.StatusOK, resp)
}

func DeleteUser(c echo.Context) error {
	user_id := c.Param("id")
	user := &models.User{}

	convertedId, _ := strconv.Atoi(user_id)
	user.UserId = convertedId
	fmt.Println(user)
	resp := user.DeleteUser() //Delete account
	return c.JSON(http.StatusOK, resp)
}
