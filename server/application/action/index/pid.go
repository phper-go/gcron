package index

import (
	"github.com/phper-go/frame/func/conv"
	"github.com/phper-go/frame/web/core"
)

type IndexPidAction struct {
	core.Controller
	AppName interface{}
}

func (this *IndexPidAction) Run() {

	this.Ctx().Output.Content = conv.Bytes(this.AppName)
}

func init() {
	core.RegisterController("/", &IndexPidAction{})
	core.RegisterController("/index/pid", &IndexPidAction{})
	core.RegisterController("/admin/index/login", &IndexPidAction{})
}
