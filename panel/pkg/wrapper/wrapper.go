package wrapper

import (
	"github.com/labstack/echo/v4"
)

type EmptyRequest struct {
}

type Wrapper[T any] struct {
	Request T
	Exec    func(T, ResponseWriter) error
}

func New[T any](f func(T, ResponseWriter) error) *Wrapper[T] {
	return &Wrapper[T]{
		Exec: f,
	}
}

func (s *Wrapper[T]) Handle() func(echo.Context) error {
	return func(c echo.Context) error {
		var t T
		if err := c.Bind(&t); err != nil {
			return err
		}

		return s.Exec(t, c)
	}
}
