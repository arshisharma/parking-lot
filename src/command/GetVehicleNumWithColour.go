// package command with 'registration_numbers_for_cars_with_colour'
// command implementation
package command

import (
	"fmt"
	"custerr"
	"strings"
)

// commandGetRegNumWithColour defined arguments and related methods
type commandGetRegNumWithColour struct {
	Command
	Color string
}

// NewcommandGetRegNumWithColour new command instance
func NewcommandGetRegNumWithColour() *commandGetRegNumWithColour {
	var command = new(commandGetRegNumWithColour)
	command.command = "registration_numbers_for_cars_with_colour"
	return command
}

// Help to print command help information
func (crc *commandGetRegNumWithColour) Help() string {
	return `🔸  registration_numbers_for_cars_with_colour <colour>
	Searching registration number of vehicle by using colour.
	Eg: registration_numbers_for_cars_with_colour White`;
	}

	// Parse to parse arguments
	func (crc *commandGetRegNumWithColour) Parse(argString string) error {
		crc.Command.Parse(argString)
		crc.Color = crc.Args[0]
		return nil
	}

	// Verify to check the provided parameters are valid or not
	func (crc *commandGetRegNumWithColour) Verify() error {
		if "" == crc.Color {
			return custerr.ErrInvalidParams
		}
		return nil
	}

	// Run to execute the command and provide result
	func (crc *commandGetRegNumWithColour) Run() (string, error) {
		var outPutList = []string{}
		slots, err := pL.SearchVehicleByColor(crc.Color)
		if nil == err {
			for _, s := range slots {
				v := s.GetVehicle()
				outPutList = append(
					outPutList,
					fmt.Sprintf("%v", v.GetNumber()),
				)
			}
		} else {
			outPutList = []string{
				err.Error(),
			}
		}
		crc.OutPut = strings.Join(outPutList, " ,")
		return crc.OutPut, err
	}

