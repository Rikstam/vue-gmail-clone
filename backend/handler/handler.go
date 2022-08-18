package handler

import (
	"encoding/json"
	"gmail-clone-backend/filesystem_store"
	"gmail-clone-backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Store *filesystem_store.FileSystemEmailStore
}

func (h *Handler) GetEmails(c echo.Context) error {
	emails := h.Store.GetEmails()
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&emails)
}

func (h *Handler) GetEmail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	email := h.Store.GetEmail(id)

	if email == nil {
		return echo.NewHTTPError(http.StatusNotFound, "email not found")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&email)
}

func (h *Handler) UpdateEmail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	email := new(models.Email)

	if err := c.Bind(email); err != nil {
		return err
	}

	newEmail := h.Store.UpdateEmail(id, email)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&newEmail)
}
