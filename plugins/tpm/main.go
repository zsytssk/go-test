package tpm

import (
	"fmt"
	"go-test/modules"
)

type tpmPlugin struct{}

func (tpmPlugin) GetStart() {
	fmt.Println("tpm start")
}

func init() {
	fmt.Println("tpm init")
	modules.Modules = append(modules.Modules, tpmPlugin{})
}
