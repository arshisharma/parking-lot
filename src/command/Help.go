// package command with 'help' command implementation
package command

import "fmt"

// commandHelp defined arguments and related methods
type commandHelp struct {
	Command
}

// NewcommandHelp new park command instance
func NewcommandHelp() *commandHelp {
	var command = new(commandHelp)
	command.command = "help"
	return command
}

// Help to print help of get permission command
func (cp *commandHelp) Help() string {
	return `🔸  help
	Shows command's help`
}

// Parse to parse arguments
func (cp *commandHelp) Parse(argString string) error {
	cp.Command.Parse(argString)
	return nil
}

// Verify to check the provided parameters are valid or not
func (cp *commandHelp) Verify() error {
	return nil
}

// Run to execute the command and provide result
func (cp *commandHelp) Run() (string, error) {
	for _, command := range mgrcommand.Commands {
		fmt.Println(command.Help())
	}
	return "", nil
}
