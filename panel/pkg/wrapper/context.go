package wrapper

import "github.com/labstack/echo/v4"

type Context struct {
	echo.Context
}

func (c *Context) Error(err HTTPResponseError) error {
	return c.JSON(err.GetHTTPCode(), map[string]any{"message": err.GetMessage()})
}
