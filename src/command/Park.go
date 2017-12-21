// package command with 'park' command implementation
package command

import (
	"fmt"
	"custerr"
	"vehicle"
)

// commandPark defined arguments and related methods
type commandPark struct {
	Command
	Vehicle *vehicle.Vehicle
}

// NewcommandPark new park command instance
func NewcommandPark() *commandPark {
	var command = new(commandPark)
	command.command = "park"
	return command
}

// Help to print help of park command
func (cp *commandPark) Help() string {
	return `🔸  park <vehicle registrion number> <colour>
	Park vehicle in slot number, provided registration number, colour
	Eg: park KA-01-HH-1234 White`
}

// Parse to parse arguments
func (cp *commandPark) Parse(argString string) error {
	cp.Command.Parse(argString)
	return nil
}

// Verify to check the provided parameters are valid or not
func (cp *commandPark) Verify() error {
	if 2 != len(cp.Args) {
		return custerr.ErrInvalidParams
	}

	cp.Vehicle = vehicle.New(cp.Args[0], cp.Args[1])
	return nil
}

// Run to execute the command and provide result
func (cp *commandPark) Run() (string, error) {
	slots, err := pL.ReportVehicleByNumber(cp.Vehicle.Number)
	if len(slots) == 0 {
		oSlot, err1 := pL.AddVehicle(cp.Vehicle)
		if nil == err1 {
			cp.OutPut = fmt.Sprintf("Allocated slot number: %v",oSlot.GetNumber())
		}else {
			cp.OutPut = err1.Error()
			return cp.OutPut, err1
		}
	}else {
		cp.OutPut = custerr.ErrCannotPark.Error()
		return cp.OutPut, err
	}

	return cp.OutPut, nil
}
