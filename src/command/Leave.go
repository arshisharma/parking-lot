// package command with 'leave' command implementation
package command

import (
	"fmt"
	"custerr"
	"slot"
	"strconv"
)

// commandLeave defined arguments and related methods
type commandLeave struct {
	Command
	SlotNumber uint64
}

// NewcommandLeave new leave command instance
func NewcommandLeave() *commandLeave {
	var command = new(commandLeave)
	command.command = "leave"
	return command
}

// Help to print help of leave command
func (cl *commandLeave) Help() string {
	return `🔸  leave <slot number>
	Remove vehicle from parking slot, slot will be freed
	Eg: leave 4`;
}

// Parse to parse arguments
func (cl *commandLeave) Parse(argString string) error {
	cl.Command.Parse(argString)
	if "" != cl.Args[0] {
		val, err := strconv.ParseUint(cl.Args[0], 0, 64)
		if nil != err {
			return custerr.ErrInvalidParams
		}
		cl.SlotNumber = uint64(val)
	}
	return nil
}

// Verify to check the provided parameters are valid or not
func (cl *commandLeave) Verify() error {
	if !slot.IsValidSlotNumber(cl.SlotNumber) {
		return custerr.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (cl *commandLeave) Run() (string, error) {
	oSlot, err := pL.RemoveVehicleBySlotNumber(cl.SlotNumber)
	if nil == err {
		cl.OutPut = fmt.Sprintf(
			"Slot number %v is free",
			oSlot.GetNumber(),
		)
	}
	return cl.OutPut, err
}
