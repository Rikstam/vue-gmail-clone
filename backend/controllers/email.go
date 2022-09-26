package controllers

import (
	"context"
	"encoding/json"
	"gmail-clone-backend/ent"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmailController struct {
	Client *ent.Client
}

func (controller *EmailController) CreateEmail(c echo.Context) error {
	var email ent.Email

	if err := c.Bind(&email); err != nil {
		c.Logger().Error("Bind: ", err)
		return c.String(http.StatusBadRequest, "Bind: "+err.Error())
	}

	e := controller.Client.Email.
		Create().
		SetFrom(email.From).
		SetSubject(email.Subject).
		SetBody(email.Body).
		SetSentAt(email.SentAt)

	newEmail, err := e.Save(context.Background())
	if err != nil {
		c.Logger().Error("Insert: ", err)
		return c.String(http.StatusBadRequest, "Save: "+err.Error())
	}

	c.Logger().Infof("inserted email: %v", newEmail.ID)
	return c.JSON(http.StatusCreated, &newEmail)
}

func (controller *EmailController) GetEmails(c echo.Context) error {
	emails, err := controller.Client.Email.
		Query().
		All(context.Background())
	if err != nil {
		return c.String(http.StatusBadRequest, "failed querying emails: "+err.Error())
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&emails)
}

func (controller *EmailController) GetEmail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	c.Logger().Info("querying email with")
	emailResult, err := controller.Client.Email.
		Get(context.Background(), id)

	if err != nil {
		if !ent.IsNotFound(err) {
			c.Logger().Error("Get: ", err)
			return c.String(http.StatusBadRequest, "Get: "+err.Error())
		}
		return c.String(http.StatusNotFound, "Not Found")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&emailResult)
}
