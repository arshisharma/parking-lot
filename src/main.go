package main

import (
	"command"
	"os"
)

func main() {

	// Process the input file and give an interactive shell
	if len(os.Args) > 1 && "" != os.Args[1] {
		command.NewFilecommandProcessor(os.Args[1]).Process()
	}else {
		// Interactive shell will provide to user to operate
		// with file data and console process
		command.NewShell().Process()
	}
}
