package controller

import (
	"github.com/phper-go/frame/web/core"
)

type baseController struct {
	core.Controller
}

func (this *baseController) Prepare() bool {

	this.Controller.Prepare()
	var layouts = []string{
		"/layout/main",
		"/layout/header",
		"/layout/menu",
		"/layout/footer",
	}
	this.Ctx().Theme.SetLayouts(layouts)
	return true
}
