package main

import (
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

const dbFileName = "db.json"

func main() {

	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	e.GET("/emails/", handler.GetEmails)

	e.GET("/emails/:id", handler.GetEmail)

	e.PUT("/emails/:id", handler.UpdateEmail)

	e.Logger.Fatal(e.Start(":1323"))
}
