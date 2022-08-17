package handler

import (
	"encoding/json"
	"gmail-clone-backend/config"
	"gmail-clone-backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var store = config.InitStore()

func GetEmails(c echo.Context) error {
	emails := store.GetEmails()
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&emails)
}

func GetEmail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	email := store.GetEmail(id)

	if email == nil {
		return echo.NewHTTPError(http.StatusNotFound, "email not found")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&email)
}

func UpdateEmail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	email := new(models.Email)

	if err := c.Bind(email); err != nil {
		return err
	}

	newEmail := store.UpdateEmail(id, email)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&newEmail)
}
