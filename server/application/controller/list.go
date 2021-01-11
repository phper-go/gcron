package controller

import (
	"github.com/phper-go/frame/web/core"
)

type ListController struct {
	core.Controller
}

func (this *ListController) IndexAction() {
	
	
}

func init() {
	core.RegisterController("/list", &ListController{})
}
