package wrapper

import (
	"errors"
	"net/http"
	"pluto"

	"github.com/golang-jwt/jwt/v5"
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

	return w.(interface {
		Get(key string) interface{}
	}).Get("user").(*jwt.Token).Claims.(jwt.MapClaims), nil
}

func WriteError(err error, writer ResponseWriter) error {
	var perr *pluto.Error
	if errors.As(err, &perr) {
		return writer.JSON(perr.HTTPCode, perr)
	}
	return writer.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
}
