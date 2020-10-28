package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	helper "github.com/tavvfiq/cafe-rest-api/apiHelpers"
	tb "github.com/tavvfiq/cafe-rest-api/database/mysql"
	um "github.com/tavvfiq/cafe-rest-api/models"
)

// RegisterHandler handler for register request
func RegisterHandler(ctx echo.Context) error {
	r := new(tb.RegistReq)
	u := tb.User{}
	var err error
	if err := ctx.Bind(r); err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusNoContent, helper.ErrorResponse{Status: http.StatusNoContent, Message: "Error registering user"})
	}
	if u, err = um.RegisterUser(r); err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusNoContent, helper.ErrorResponse{Status: http.StatusNoContent, Message: "Error registering user"})
	}
	return ctx.JSON(http.StatusOK, u)
}
