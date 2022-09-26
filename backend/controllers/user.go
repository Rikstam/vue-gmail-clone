package controllers

import (
	"context"
	"gmail-clone-backend/ent"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Client *ent.Client
}

func (controller *UserController) CreateUser(c echo.Context) error {
	var user ent.User

	if err := c.Bind(&user); err != nil {
		c.Logger().Error("Bind: ", err)
		return c.String(http.StatusBadRequest, "Bind: "+err.Error())
	}

	hashedPassword, err := hashAndSalt([]byte(user.Password))

	if err != nil {
		c.Logger().Error("PasswordHash: ", err)
		return c.String(http.StatusBadRequest, "Save: "+err.Error())
	}

	u := controller.Client.User.
		Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetPassword(hashedPassword)

	newUser, err := u.Save(context.Background())
	if err != nil {
		c.Logger().Error("Insert: ", err)
		return c.String(http.StatusBadRequest, "Save: "+err.Error())
	}

	c.Logger().Infof("inserted user: %v", newUser.ID)
	return c.JSON(http.StatusCreated, &newUser)
}

func hashAndSalt(pwd []byte) (string, error) {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}
