package apihelpers

import tb "github.com/tavvfiq/cafe-rest-api/interfaces"

// DefaultResponse default response from controller
type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

// ErrorResponse helpers for returning error response from controllers
type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"msg"`
}

// UserSuccessResponse helpers for returning success response from controllers
type UserSuccessResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"msg"`
	Data    tb.User `json:"data"`
}

// MenuResponsePagination Menu response with pagination
type MenuResponsePagination struct {
	Status   int         `json:"status"`
	Message  string      `json:"msg"`
	Menu     []tb.Menu   `json:"menu"`
	PageInfo tb.PageInfo `json:"pageInfo"`
}
