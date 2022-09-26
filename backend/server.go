package main

import (
	"context"
	"gmail-clone-backend/config"
	"gmail-clone-backend/controllers"
	"gmail-clone-backend/ent"
	"gmail-clone-backend/handler"
	"gmail-clone-backend/models"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

type EmailStore interface {
	GetEmails() models.Inbox
	GetEmail(id int) models.Email
}

// Controller is a controller for this application.

func main() {
	connectionString := os.Getenv("DB_CONN") //"host=localhost port=5432 user=riksa dbname=gmail_clone sslmode=disable password="

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	var store = config.InitStore()

	emailHandler := handler.Handler{Store: store}

	emailController := &controllers.EmailController{Client: client}
	userController := controllers.UserController{Client: client}
	authController := controllers.AuthController{Client: client}

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Logger.SetOutput(os.Stderr)

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &controllers.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}

	// Restricted group
	r := e.Group("/restricted")

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", authController.Restricted)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	e.POST("/emails/", emailController.CreateEmail)

	e.GET("/emails/", emailController.GetEmails)

	e.GET("/emails/:id", emailController.GetEmail)

	e.PUT("/emails/:id", emailHandler.UpdateEmail)

	e.POST("/users/", userController.CreateUser)

	e.POST("/auth/token", authController.Login)

	e.Logger.Fatal(e.Start(":1323"))
}
