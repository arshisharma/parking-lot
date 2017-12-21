// package command implements the basic shell command execution process
// Includes help, parse, verify, run
package command

import (
	"custerr"
	"strings"
	"parking"
)

// IManager should have behaviour run and a base parse
type IManager interface {
	Parse() error
	Run() (string, error)
}

var pL *parking.ParkingLot
// Manager handles requested command and available command's list
type Manager struct {
	command, argString string
	Commands       map[string]ICommand
}

var mgrcommand = &Manager{
	Commands: make(map[string]ICommand),
}

// NewManager return command manager
func NewManager() *Manager {
	mgrcommand.Register(NewcommandCreateParkingLot())
	mgrcommand.Register(NewcommandPark())
	mgrcommand.Register(NewcommandLeave())
	mgrcommand.Register(NewcommandGetStatus())
	mgrcommand.Register(NewcommandGetRegNumWithColour())
	mgrcommand.Register(NewcommandGetSlotNumWithColour())
	mgrcommand.Register(NewcommandGetSlotNumWithVehicleNum())
	mgrcommand.Register(NewcommandHelp())
	return mgrcommand
}

// Register Command registration with manager
func (cm *Manager) Register(command ICommand) {
	commandName := strings.ToLower(command.GetName())
	cm.Commands[commandName] = command
}

// IsValidCommad verifies the requested command is valid or not
func (cm *Manager) IsValidCommad(commandName string) bool {
	commandName = strings.ToLower(commandName)
	_, ok := cm.Commands[commandName]
	return ok
}

// Parse requested command and arguments
func (cm *Manager) Parse(commandString string) error {
	commandString = strings.Trim(commandString, " \n\t")
	results := strings.SplitN(commandString, " ", 2)
	cm.command = strings.ToLower(results[0])
	if len(results) > 1 {
		cm.argString = results[1]
	}
	if "" == cm.command {
		return custerr.ErrInvalidCommand
	}
	return nil
}

// Run the requested command and provide output
func (cm *Manager) Run(commandString string) (string, error) {
	err := cm.Parse(commandString)
	if nil != err {
		return "", err
	}
	command, ok := cm.Commands[cm.command]
	if ok {
		command.Clear()
		err := command.Parse(cm.argString)
		if nil != err {
			return "", custerr.ErrInvalidParams
		}
		if nil == command.Verify() {
			return command.Run()
		}
		return "", custerr.ErrInvalidParams
	}
	return "", custerr.ErrInvalidCommand
}
