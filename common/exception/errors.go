package exception

import "github.com/zsljava/gokit/common/response"

var (
	// common errors
	ErrSuccess             = response.NewError(0, "ok")
	ErrBadRequest          = response.NewError(400, "Bad Request")
	ErrUnauthorized        = response.NewError(401, "Unauthorized")
	ErrNotFound            = response.NewError(404, "Not Found")
	ErrInternalServerError = response.NewError(500, "Internal Server Error")
)
