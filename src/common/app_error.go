package common

import "net/http"

type AppError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	RootError error  `json:"-"`
	Key       string `json:"key"`
}

func NewAppError(errorCode int, message string, rootError error, key string) *AppError {
	return &AppError{
		ErrorCode: errorCode,
		Message:   message,
		RootError: rootError,
		Key:       key,
	}
}

func ServerError(err error) *AppError {
	return NewAppError(
		http.StatusInternalServerError,
		"something went wrong in the server",
		err,
		"ServerError",
	)
}

func Error(err error, message string) *AppError {
	return &AppError{
		ErrorCode: http.StatusInternalServerError,
		Message:   message,
		RootError: err,
	}
}

func (c *AppError) Throw() {
	panic(c)
}
