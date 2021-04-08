package controllers

import (
	"echolearn/models"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	resp := models.Login(username, password)
	return c.JSON(http.StatusOK, resp)
}
