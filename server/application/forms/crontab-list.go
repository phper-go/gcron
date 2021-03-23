package forms

import (
	"gcron/server/application/model"
	"html/template"

	"github.com/phper-go/frame/db"
	"github.com/phper-go/frame/func/conv"
	"github.com/phper-go/frame/func/html"
	"github.com/phper-go/frame/web/form"
	"github.com/phper-go/frame/web/form/element"
)

type CrontabListForm struct {
	form.Search
	Host   interface{}
	Status interface{}
}

func (this *CrontabListForm) Title() []map[string]string {
	return []map[string]string{
		map[string]string{"id": "ID"},
		map[string]string{"name": "脚本名称"},
		map[string]string{"host": "执行机器"},
		map[string]string{"time": "执行时间"},
		map[string]string{"cmd": "脚本命令"},
		map[string]string{"do": "操作"},
	}
}

func (this *CrontabListForm) DBCmd(sdb *db.DB) *db.DBCommand {

	cron := &model.Crontab{}
	dbCmd := sdb.Cmd().Table(cron.Table())
	if this.Host != "" {
		dbCmd.Where("host=?", this.Host)
	}
	if this.Status != "" {
		dbCmd.Where("status=?", this.Status)
	}
	return dbCmd.Order("id desc")
}

func (this *CrontabListForm) List(dbCmd *db.DBCommand) ([]map[string]interface{}, error) {

	list, err := dbCmd.QueryRows()

	for _, one := range list {
		id := conv.String(one["id"])
		one["do"] = template.HTML(html.Link("编辑", "/crontab/edit?id="+id))
	}

	return list, err
}

func (this *CrontabListForm) Builder() *form.Builder {

	builder := &form.Builder{}

	builder.Construct("?", "GET")

	builder.Elements = []element.Interface{

		element.Text("host", "执行机器", "机器hostname或者ip"),
		element.Select("status", "是否启用", map[string]string{"": "全部", "0": "禁用", "1": "启用"}),
		element.Button("submit", "提交"),
	}

	builder.Bind(this)

	return builder
}
