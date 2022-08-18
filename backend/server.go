package main

import (
	"gmail-clone-backend/config"
	"gmail-clone-backend/handler"
	"gmail-clone-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EmailStore interface {
	GetEmails() models.Inbox
	GetEmail(id int) models.Email
}

func main() {

	var store = config.InitStore()

	emailHandler := handler.Handler{Store: store}

	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	e.GET("/emails/", emailHandler.GetEmails)

	e.GET("/emails/:id", emailHandler.GetEmail)

	e.PUT("/emails/:id", emailHandler.UpdateEmail)

	e.Logger.Fatal(e.Start(":1323"))
}
