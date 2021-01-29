package controller

import (
	"gcron/server/application/forms"

	"github.com/phper-go/frame/ext/validator"
	"github.com/phper-go/frame/web/core"
)

type CronController struct {
	baseController
}

func (this *CronController) IndexAction() {

	form := &forms.CronIndex{}

	this.Ctx().Input.Bind(form)

	if this.Ctx().Input.Server.IsPost {
		validator.Check(form)
	}

	this.Ctx().Theme.Assign("form", form)
	this.Ctx().Theme.Assign("builder", form.Builder())
	this.Ctx().Theme.Render("/layout/form")
}

func init() {
	core.RegisterController("/cron", &CronController{})
}
