package forms

import (
	"github.com/phper-go/frame/ext/validator"
	"github.com/phper-go/frame/web/form"
	"github.com/phper-go/frame/web/form/element"
)

type CronIndex struct {
	validator.Api
	Id    interface{}
	Name  interface{}
	Phone interface{}
}

func (this *CronIndex) Rules() validator.Rules {

	return validator.Rules{
		&validator.Required{Fields: "name,phone"},
	}
}

func (this *CronIndex) Builder() *form.Builder {

	builder := form.BuildForm("?", "POST", this, []element.Interface{

		element.Hidden("id"),
		element.Text("name", "脚本名称", "脚本文件名称"),
		element.Text("phone", "手机号", "18612345678"),
		element.Button("submit", "提交"),
	})

	return builder
}
