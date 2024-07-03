package handlers

import "errors"

var (
	errUnhandled = errors.New("request failed - please try again")
)

type unhandledError struct {
	Message string `json:"message"`
}
