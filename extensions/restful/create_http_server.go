package restful

import (
	"pluto"
	"sync"

	"github.com/labstack/echo/v4"
)

const ProcessorName_CreateHTTPServer = "CREATE_HTTP_SERVER"

func creator_CreateHTTPServer(args []pluto.Value) (p pluto.Processor, err error) {
	defer pluto.CreatorPanicHandler(ProcessorName_CreateHTTPServer, &err)()
	return &CreateHTTPServer{
		Name:    pluto.Find("name", args...).Get().(string),
		Address: pluto.Find("address", args...).Get().(string),
	}, nil
}

var Servers map[string]*echo.Echo
var ServersMutex *sync.RWMutex

type CreateHTTPServer struct {
	Name    string
	Address string
	//ServeTLS bool
}

func (c *CreateHTTPServer) Process(processable pluto.Processable) (pluto.Processable, bool) {
	e := echo.New()

	ServersMutex.Lock()
	Servers[c.Name] = e
	ServersMutex.Unlock()

	if err := e.Start(c.Address); err != nil {
		pluto.ApplicationLogger.Error(pluto.ApplicationLog{
			Message: "Failed to start http server",
			Extra: map[string]any{
				"name":    c.Name,
				"address": c.Address,
			},
		})
		return nil, false
	}

	return processable, true
}
