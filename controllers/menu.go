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

// GetMenuHandler handler for getting all menu with pagination
func GetMenuHandler(ctx echo.Context) error {
	m := []tb.Menu{}
	qm := make(map[string]string)
	var err error

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// get query from url
	qm["search"] = ctx.QueryParam("search")
	qm["filter"] = ctx.QueryParam("filter")
	qm["sortby"] = ctx.QueryParam("sortby")
	qm["order"] = ctx.QueryParam("order")
	qm["page"] = ctx.QueryParam("page")
	qm["limit"] = ctx.QueryParam("limit")

	m, err = models.GetMenu(c, qm)

	if err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusNoContent, ah.ErrorResponse{Status: http.StatusNoContent, Error: "No menu"})
	}

	return ctx.JSON(http.StatusOK, ah.MenuResponsePagination{Status: http.StatusOK, Message: "", Menu: m})
}
