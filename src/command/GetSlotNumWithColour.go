// package command with 'slot_numbers_for_cars_with_colour'
// command implementation
package command

import (
	"fmt"
	"custerr"
	"strings"
)

// commandGetSlotNumWithColour defined arguments and related methods
type commandGetSlotNumWithColour struct {
	Command
	Color string
}

// NewcommandGetSlotNumWithColour new command instance
func NewcommandGetSlotNumWithColour() *commandGetSlotNumWithColour {
	var command = new(commandGetSlotNumWithColour)
	command.command = "slot_numbers_for_cars_with_colour"
	return command
}

// Help to print command help information
func (cgc *commandGetSlotNumWithColour) Help() string {
	return `🔸  slot_numbers_for_cars_with_colour <colour>
	Fetch all slot numbers of cars with known colour
	Eg: slot_numbers_for_cars_with_colour White`
}

// Parse to parse arguments
func (cgc *commandGetSlotNumWithColour) Parse(argString string) error {
	cgc.Command.Parse(argString)
	cgc.Color = cgc.Args[0]
	return nil
}

// Verify to check the provided parameters are valid or not
func (cgc *commandGetSlotNumWithColour) Verify() error {
	if "" == cgc.Color {
		return custerr.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (cgc *commandGetSlotNumWithColour) Run() (string, error) {
	var outPutList = []string{}
	slots, err := pL.SearchVehicleByColor(cgc.Color)
	if nil == err {
		for _, s := range slots {
			outPutList = append(
				outPutList,
				fmt.Sprintf("%v", s.GetNumber()),
			)
		}
	} else {
		outPutList = []string{err.Error()}
	}
	cgc.OutPut = strings.Join(outPutList, " ,")
	return cgc.OutPut, err
}
