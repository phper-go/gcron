package main

import (
	_ "gcron/server/application/action"
	_ "gcron/server/application/controller"
	"gcron/server/application/run"
	"os"

	"github.com/phper-go/frame"
)

func main() {

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "start")
		// os.Args = append(os.Args, "--config=abc")
	}

	frame.Run(run.App, os.Args)
}
