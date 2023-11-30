package pluto

import (
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

var Env EnvSpec

type EnvSpec struct {
	Debug             bool
	HOST              string `envconfig:"HOST"`
	HTTPServerAddress string `envconfig:"HTTP_SERVER"`
}

func init() {
	if err := envconfig.Process(strings.ToUpper(Name), &Env); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Environment variables: %v\n", err)
		os.Exit(1)
	}
}
