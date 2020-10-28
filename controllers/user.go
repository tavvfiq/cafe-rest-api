package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	ah "github.com/tavvfiq/cafe-rest-api/apiHelpers"
	tb "github.com/tavvfiq/cafe-rest-api/database/mysql"
	"github.com/tavvfiq/cafe-rest-api/models"
)

// RegisterHandler handler for register request
func RegisterHandler(ctx echo.Context) error {
	r := new(tb.RegistReq)
	u := tb.User{}
	var err error

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.Bind(r); err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusNoContent, ah.ErrorResponse{Status: http.StatusNoContent, Message: "Error registering user"})
	}
	u, err = models.RegisterUser(c, r)
	if err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusNoContent, ah.ErrorResponse{Status: http.StatusNoContent, Message: "Error registering user"})
	}
	return ctx.JSON(http.StatusOK, u)
}
