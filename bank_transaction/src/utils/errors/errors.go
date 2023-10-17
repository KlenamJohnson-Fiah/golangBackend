package errors

import "net/http"

type Error struct {
	ErrorCode int
	Message   string
}

func ErrorBadRequest(msg string) *Error {

	return &Error{ErrorCode: http.StatusBadGateway, Message: msg}
}

func ErrorInternalServerError(msg string) *Error {
	return &Error{ErrorCode: http.StatusInternalServerError, Message: msg}
}

func ErrorNotFound(msg string) *Error {
	return &Error{ErrorCode: http.StatusNotFound, Message: msg}
}
