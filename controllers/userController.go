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
	result, err := models.GetAllUser()
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, u.Respond(500, "Internal Server Error"))
	}
	if len(result) < 1 {
		return c.JSON(http.StatusOK, u.Respond(404, "Data not found"))
	}
	return c.JSON(http.StatusOK, u.RespondWithData(200, "Success", result))
}

func CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusOK, u.Respond(400, err.Error()))
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, u.Respond(500, err.Error()))
	}

	user_id, err := user.StoreUser()
	if err != nil {
		return c.JSON(http.StatusOK, u.Respond(500, err.Error()))
	}

	tokenString := models.CreateToken(uint(user_id), user.Username)
	data := map[string]interface{}{
		"user_id": user_id,
		"token":   tokenString,
	}
	return c.JSON(http.StatusOK, u.RespondWithData(200, "Success", data))
}

func UpdateUser(c echo.Context) error {
	user_id := c.Param("id")
	user := &models.APIUser{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusOK, u.Respond(400, err.Error()))
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, u.Respond(500, err.Error()))
	}

	convertedId, _ := strconv.Atoi(user_id)
	user.UserId = convertedId
	if err := user.UpdateUser(); err != nil { //Update account
		return c.JSON(http.StatusOK, u.Respond(500, err.Error()))
	}
	return c.JSON(http.StatusOK, u.Respond(200, "Success"))
}

func DeleteUser(c echo.Context) error {
	user_id := c.Param("id")
	user := &models.User{}

	convertedId, _ := strconv.Atoi(user_id)
	user.UserId = convertedId
	if err := user.DeleteUser(); err != nil { //Delete account
		return c.JSON(http.StatusOK, u.Respond(500, err.Error()))
	}
	return c.JSON(http.StatusOK, u.Respond(200, "Success"))
}
