package delivery

import (
	"net/http"
	"pluto"
	"pluto/panel/auth"
	"pluto/panel/logview"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func init() {
	pluto.FindHTTPHost("panel").GET("/api/v1/logs/bind/:token", func(c echo.Context) error {
		ws, err := (&websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// TODO: Check the origin
				return true
			},
			// TODO: No need to have read buffer
		}).Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		(&logview.BindApplicationLogs{
			Identifier: pluto.ExternalIdentifier{
				Name: c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string),
				Kind: pluto.KindAdmin,
			},
			Writer: &WSConnWrapper{Conn: ws},
		}).Bind()
		return nil
	},
		echojwt.WithConfig(echojwt.Config{
			SigningKey:  auth.JWTSecretKey,
			TokenLookup: "param:token",
		}),
	)
}

type WSConnWrapper struct {
	Conn *websocket.Conn
}

func (w *WSConnWrapper) Write(msg []byte) (int, error) {
	if err := w.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return 0, err
	}
	return len(msg), nil
}
