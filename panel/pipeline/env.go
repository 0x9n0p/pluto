package pipeline

import (
	"fmt"
	"os"
	"pluto"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

var Env EnvSpec

type EnvSpec struct {
	StoragePath string `envconfig:"PANEL_STORAGE" required:"true"`
}

func init() {
	if err := envconfig.Process(strings.ToUpper(pluto.Name), &Env); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Environment variables: %v\n", err)
		os.Exit(1)
	}
}
