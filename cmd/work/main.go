package main

import (
	"os"

	"github.com/rwxrob/bonzai/pkg/core/run"

	"github.com/Chaitanyabsprip/workdirs"
)

func main() {
	os.Setenv("SHELL", "bash")
	if len(os.Getenv(`DEBUG`)) > 0 {
		run.AllowPanic = true
	}
	workdirs.Cmd.Run()
}
