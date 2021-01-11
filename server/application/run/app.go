package run

import (
	_ "github.com/phper-go/frame/ext/doc"
	"github.com/phper-go/frame/web/core"
)

var App = &app{}

type app struct {
	core.App
}
