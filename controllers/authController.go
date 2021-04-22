package controllers

import (
	"echolearn/models"
	u "echolearn/utils"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := models.Login(username, password)

	if err != nil {
		return c.JSON(http.StatusOK, u.Respond(500, err.Error()))
	}

	data := map[string]interface{}{
		"token": token,
	}
	return c.JSON(http.StatusOK, u.RespondWithData(200, "Success", data))
}
