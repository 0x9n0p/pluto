package wrapper

import (
	"net/http"
	"pluto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetJWTClaims(w ResponseWriter) (val map[string]any, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = &pluto.Error{
				HTTPCode: http.StatusUnauthorized,
				Message:  "No claims provided",
			}
		}
	}()

	return w.(echo.Context).Get("user").(*jwt.Token).Claims.(jwt.MapClaims), nil
}
