package controllers

import (
	"context"
	"gmail-clone-backend/ent"
	"gmail-clone-backend/ent/user"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type JWTCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type AuthController struct {
	Client *ent.Client
}

func (controller *AuthController) CreateJWTtoken(c echo.Context) error {
	return nil
}

func validatePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (controller *AuthController) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	log.Println(password)
	ctx := context.Background()

	userResult, err := controller.Client.User.
		Query().
		Where(user.Username(username)).
		Only(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			c.Logger().Error("Get: ", err)
			return c.String(http.StatusBadRequest, "Get: "+err.Error())
		}
		return c.String(http.StatusNotFound, "Not Found")
	}

	if !validatePassword(userResult.Password, []byte(password)) {
		return echo.ErrUnauthorized
	}

	claims := &JWTCustomClaims{
		username,
		false,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (controller *AuthController) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
