package database

import (
	"fmt"
	"os"
	"pluto"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

var Env EnvSpec

type EnvSpec struct {
	DatabasePath string `envconfig:"PANEL_DATABASE" default:"./var/panel.db"`
}

func init() {
	if err := envconfig.Process(strings.ToUpper(pluto.Name), &Env); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Environment variables: %v\n", err)
		os.Exit(1)
	}

	setup()
}
