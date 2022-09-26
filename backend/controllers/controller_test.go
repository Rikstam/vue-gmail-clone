package controllers

import (
	"context"
	"encoding/json"
	"gmail-clone-backend/ent/enttest"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var userJSON = `{"first_name":"Jon", "last_name": "Snow", "username": "ThePrinceThatWasPromised", "email":"jon@labstack.com"}`

var emailJSON = `{"from": "team@vuemastery.com",
	"subject": "JAFFAA KAIKILLE",
	"body": "The opening keynote of VueConf US this year was Evan You (the creator of Vue), giving his State of the Vuenion address. He walked us through the journey of getting Vue 3 from a prototype to a reality the past year. He also dove into Vue's overall growth in the community.",
	"sentAt": "2020-03-27T18:25:43.511Z",
	"archived": true,
	"read": true}`

type EmailResponse struct {
	Id      int
	From    string
	Subject string
	SentAt  string
}

type UserResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func TestUser(t *testing.T) {
	// Setup
	e := echo.New()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	defer client.Close()
	ctx := context.Background()

	var expectedUser = UserResponse{}
	if err := json.Unmarshal([]byte(emailJSON), &expectedUser); err != nil {
		t.Fatal(err)
	}

	t.Run("create user via API", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/")

		ctrl := &UserController{Client: client}

		var response = UserResponse{}
		// Assertions
		if assert.NoError(t, ctrl.CreateUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			if err := json.Unmarshal([]byte(rec.Body.Bytes()), &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, "ThePrinceThatWasPromised", response.Username)
		}
	})

	t.Run("create user via client", func(t *testing.T) {
		password, err := bcrypt.GenerateFromPassword([]byte("kissa123"), bcrypt.MinCost)
		if err != nil {
			t.Fatal(err)
		}
		user, err := client.User.
			Create().
			SetUsername("Jaffapena").
			SetFirstName("Jaffa").
			SetLastName("Pena").
			SetEmail("jaffa@pena.com").
			SetPassword(string(password)).
			Save(ctx)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(user)
	})
}

func performRequest(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestEmail(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	e := echo.New()

	ctrl := &EmailController{Client: client}

	var expectedEmail = EmailResponse{}
	if err := json.Unmarshal([]byte(emailJSON), &expectedEmail); err != nil {
		t.Fatal(err)
	}

	t.Run("create email", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(emailJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/emails/")

		var response = EmailResponse{}

		// Assertions
		if assert.NoError(t, ctrl.CreateEmail(c)) {
			if err := json.Unmarshal([]byte(rec.Body.Bytes()), &response); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, 1, response.Id)
			assert.Equal(t, expectedEmail.Subject, response.Subject)
			assert.Equal(t, "team@vuemastery.com", response.From)
			assert.Equal(t, "2020-03-27T18:25:43.511Z", response.SentAt)
		}
	})

	t.Run("get all emails", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/emails/")

		if assert.NoError(t, ctrl.GetEmail(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

}
