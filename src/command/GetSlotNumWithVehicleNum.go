// package command with 'slot_number_for_registration_number'
// command implementation
package command

import (
	"fmt"
	"custerr"
	"strings"
)

// commandGetSlotNumWithVehicleNum defined arguments and related methods
type commandGetSlotNumWithVehicleNum struct {
	Command
	RegistrationNumber string
}

// NewcommandGetSlotNumWithVehicleNum new command instance
func NewcommandGetSlotNumWithVehicleNum() *commandGetSlotNumWithVehicleNum {
	var command = new(commandGetSlotNumWithVehicleNum)
	command.command = "slot_number_for_registration_number"
	return command
}

// Help to print command help information
func (cgs *commandGetSlotNumWithVehicleNum) Help() string {
	return `🔸  slot_number_for_registration_number <registration number>
	Search for slot number using vehicle registration number
	Eg: slot_number_for_registration_number KA-01-HH-3141`
}

// Parse to parse arguments
func (cgs *commandGetSlotNumWithVehicleNum) Parse(argString string) error {
	cgs.Command.Parse(argString)
	cgs.RegistrationNumber = cgs.Args[0]
	return nil
}

// Verify to check the provided parameters are valid or not
func (cgs *commandGetSlotNumWithVehicleNum) Verify() error {
	if "" == cgs.RegistrationNumber {
		return custerr.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (cgs *commandGetSlotNumWithVehicleNum) Run() (string, error) {
	var outPutList = []string{}
	slots, err := pL.ReportVehicleByNumber(
		cgs.RegistrationNumber,
	)
	if nil == err {
		for _, s := range slots {
			outPutList = append(
				outPutList,
				fmt.Sprintf("%v", s.GetNumber()),
			)
		}
	} else {
		outPutList = []string{
			err.Error(),
		}
	}
	cgs.OutPut = strings.Join(outPutList, " ,")
	return cgs.OutPut, err
}
