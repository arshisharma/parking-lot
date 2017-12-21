
// package command provides methods for base shell and file command processing
package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Shell stores current shell instance information
type Shell struct {
	PS1 string
}

// NewShell create a shell
func NewShell() *Shell {
	return &Shell{
		PS1: ">",
	}
}

// Process method to handle commands
func (sh *Shell) Process() error {
	reader := bufio.NewReader(os.Stdin)
	commandMgr := NewManager()
	sh.prompt()
	for {
		commandInput, _ := reader.ReadString('\n')
		commandInput = strings.TrimRight(commandInput, "\n")
		if "" != commandInput {
			out, err := commandMgr.Run(commandInput)
			processOutput(out, err)
		} else {
			commandMgr.Commands["help"].Run()
		}
		sh.prompt()
	}
	return nil
}

// prompt display command prompt PS1
func (sh *Shell) prompt() {
	fmt.Print(sh.PS1)
}
