package extension

import (
	"pluto"
)

type Installer struct {
	Extension pluto.Extension `json:"-"`
}

func (i *Installer) Install() error {
	// TODO: Check if the extension is already installed

	if err := i.Extension.Install(); err != nil {
		return err
	}

	// TODO: Save the state that this extension has been installed.

	return nil
}
