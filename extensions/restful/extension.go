package restful

import (
	"pluto"
)

type Restful struct {
	pluto.ExtensionDescriptor
}

func (r *Restful) Install() error {
	//TODO implement me
	panic("implement me")
}

func (r *Restful) Uninstall() error {
	//TODO implement me
	panic("implement me")
}

func (r *Restful) GetExtensionDescriptor() pluto.ExtensionDescriptor {
	return r.ExtensionDescriptor
}
