package main

const (
	ErrorMessageUnhandled string = "request failed - please try again later"
)

type UnhandledError struct {
	Message string `json:"message"`
}
