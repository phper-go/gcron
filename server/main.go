package main

import (
	_ "example/application/action"
	_ "example/application/controller"
	"example/application/run"
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
