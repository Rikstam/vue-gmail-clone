package main

import (
	"encoding/json"
	"gmail-clone-backend/filesystem_store"
	"gmail-clone-backend/models"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type EmailStore interface {
	GetEmails() []models.Email
}

const dbFileName = "db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := filesystem_store.NewFileSystemEmailStore(db)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	e.GET("/emails", func(c echo.Context) error {
		emails := store.GetEmails()
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(&emails)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
