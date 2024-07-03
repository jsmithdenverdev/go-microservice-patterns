package main

import "errors"

var (
	ErrUnhandled = errors.New("request failed - please try again")
)

type UnhandledError struct {
	Message string `json:"message"`
}
