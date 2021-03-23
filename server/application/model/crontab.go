package model

import (
	"github.com/phper-go/frame/db"
	"github.com/phper-go/frame/ext/validator"
)

type Crontab struct {
	Id         uint64
	Name       string
	Cmd        string
	Time       string
	Host       string
	Detail     string
	Status     string
	CreateTime string
	UpdateTime string
}

func (this *Crontab) Rules() validator.Rules {
	return validator.Rules{
		&validator.Required{Fields: "name,cmd,time,host,file"},
	}
}

func (this *Crontab) PrimaryKey() string {
	return "id"
}

func (this *Crontab) Table() string {
	return "crontab"
}

func (this *Crontab) DBCmd(sdb *db.DB) *db.DBCommand {
	return sdb.Cmd().Table(this.Table())
}
