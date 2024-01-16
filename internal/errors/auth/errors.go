package auth

import (
	"fmt"
	"net/http"
)

type error interface {
	Error() string
	//New(w http.ResponseWriter, errorString string, code int) *error
}

type AuthError struct {
	Writer      http.ResponseWriter
	ErrorString string
	Code        int
}

func (e *AuthError) Error() string {
	str := fmt.Sprintf("AUTH ERROR: %s", e.ErrorString)
	http.Error(e.Writer, str, e.Code)
	return str
}

func New(w http.ResponseWriter, errorString string, code int) *AuthError {
	err := AuthError{
		Writer:      w,
		ErrorString: errorString,
		Code:        code,
	}

	return &err
}
