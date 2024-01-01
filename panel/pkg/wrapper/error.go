package wrapper

type HTTPResponseError interface {
	HTTPCode
	Message
}

type HTTPCode interface {
	GetHTTPCode() int
}

type Message interface {
	GetMessage() string
}
