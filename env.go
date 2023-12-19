package pluto

import (
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

var Env EnvSpec

type EnvSpec struct {
	Debug               bool
	Host                []string `envconfig:"HOST" default:"localhost"`
	TCPServerAddress    string   `envconfig:"TCP_SERVER" default:"0.0.0.0:9630"`
	HTTPServerAddress   string   `envconfig:"HTTP_SERVER" default:"0.0.0.0:443"`
	HTTPCertificatePath string   `envconfig:"HTTP_CERTIFICATE_PATH" default:"ssl/plutoengine.crt"`
	HTTPKeyPath         string   `envconfig:"HTTP_KEY_PATH" default:"ssl/plutoengine.key"`
	RootStoragePath     string   `envconfig:"ROOT_STORAGE_PATH" default:"var/storage/"`
}

func init() {
	if err := envconfig.Process(strings.ToUpper(Name), &Env); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Environment variables: %v\n", err)
		os.Exit(1)
	}
}
