package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	ah "github.com/tavvfiq/cafe-rest-api/apiHelpers"
	tb "github.com/tavvfiq/cafe-rest-api/interfaces"
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
		return ctx.JSON(http.StatusNoContent, ah.ErrorResponse{Status: http.StatusNoContent, Error: "Error registering user"})
	}
	u, err = models.RegisterUser(c, r)
	if err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusNoContent, ah.ErrorResponse{Status: http.StatusNoContent, Error: "Error registering user"})
	}
	return ctx.JSON(http.StatusOK, ah.UserSuccessResponse{Status: http.StatusOK, Message: "Register Success", Data: u})
}

// LoginHandler handler for login request
func LoginHandler(ctx echo.Context) error {
	l := new(tb.LoginReq)
	u := tb.User{}

	var err error

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.Bind(l); err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusNoContent, ah.ErrorResponse{Status: http.StatusNoContent, Error: "Login Error"})
	}
	u, err = models.LoginUser(c, l)
	if err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusNoContent, ah.ErrorResponse{Status: http.StatusNoContent, Error: "Login Error"})
	}
	return ctx.JSON(http.StatusOK, ah.UserSuccessResponse{Status: http.StatusOK, Message: "Login Success", Data: u})
}
