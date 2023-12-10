package pluto

import (
	"fmt"
)

type Error struct {
	//Code     int    `json:"code"`
	HTTPCode int    `json:"-"`
	Message  string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) String() string {
	return fmt.Sprintf("http_code: %d, message: %s", e.HTTPCode, e.Message)
}
