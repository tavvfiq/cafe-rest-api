package apihelpers

import tb "github.com/tavvfiq/cafe-rest-api/database/mysql"

// ErrorResponse helpers for returning error response from controllers
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

// SuccessResponse helpers for returning success response from controllers
type SuccessResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"msg"`
	Data    tb.User `json:"data"`
}
