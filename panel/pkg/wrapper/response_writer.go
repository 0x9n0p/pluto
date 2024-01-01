package wrapper

import "net/http"

type ResponseWriter interface {
	// String sends a string response with status code.
	String(code int, s string) error

	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error

	// File sends a response with the content of the file.
	File(file string) error

	// NoContent sends a response with no body and a status code.
	NoContent(code int) error

	Error(HTTPResponseError) error

	// SetCookie adds a `Set-Cookie` header in HTTP response.
	SetCookie(cookie *http.Cookie)
}
