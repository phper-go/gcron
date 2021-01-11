package controller

import (
	"example/application/forms"

	"github.com/phper-go/frame/ext/validator"
	"github.com/phper-go/frame/web/core"
)

type IndexController struct {
	baseController
}

func (this *IndexController) IndexAction() {

	form := &forms.IndexForm{}

	this.Ctx().Input.Bind(form)

	if this.Ctx().Input.Server.IsPost {
		validator.Check(form)
	}

	this.Ctx().Theme.Assign("form", form)
	this.Ctx().Theme.Assign("builder", form.Builder())
	this.Ctx().Theme.Render("/layout/form")
}

func init() {
	core.RegisterController("/index", &IndexController{})
}
