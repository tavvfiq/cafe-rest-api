package apihelpers

import tb "github.com/tavvfiq/cafe-rest-api/interfaces"

// ErrorResponse helpers for returning error response from controllers
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

// UserSuccessResponse helpers for returning success response from controllers
type UserSuccessResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"msg"`
	Data    tb.User `json:"data"`
}
