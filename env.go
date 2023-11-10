package pluto

import (
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

var Env EnvSpec

// EnvSpec fields are not required because multiple programs going to use them
type EnvSpec struct {
	Debug     bool
	HTTPAdmin string `envconfig:"http_admin"`
}

func init() {
	if err := envconfig.Process(strings.ToUpper(Name), &Env); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Environment variables: %v\n", err)
		os.Exit(1)
	}
}
