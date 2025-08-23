package plugins

import "go-test/modules"

type ModuleInterface interface {
	GetStart()
}

func Start() {
	for _, module := range modules.Modules {
		module.(ModuleInterface).GetStart()
	}
}
