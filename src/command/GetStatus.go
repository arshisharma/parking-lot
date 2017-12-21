// package command with 'status' command implementation
package command

import (
	"fmt"
	"strings"
)

// commandGetStatus defined arguments and related methods
type commandGetStatus struct {
	Command
}

// NewcommandGetStatus new status command instance
func NewcommandGetStatus() *commandGetStatus {
	var command = new(commandGetStatus)
	command.command = "status"
	return command
}

// Help to print help of status command
func (st *commandGetStatus) Help() string   {
	return `🔸  status
	Fetch all current parked vehicle details and slot numbers
	Eg: status`
}

// Parse to parse arguments
func (st *commandGetStatus) Parse(argString string) error {
	st.Command.Parse(argString)
	return nil
}

// Verify to check the provided parameters are valid or not
func (st *commandGetStatus) Verify() error {
	return nil
}

// Run to execute the command and provide result
func (st *commandGetStatus) Run() (string, error) {
	var outPutList = []string{
		fmt.Sprintf("%-12s%-20s%-10s",
			"Slot No.",
			"Registration No",
			"Colour",
		),
		fmt.Sprintf("%-12v%-20v%-10v",
			"------------",
			"--------------------",
			"----------",
		),
	}
	slots, err := pL.ReportFilledSlots()
	if nil == err {
		for _, s := range slots {
			v := s.GetVehicle()
			outPutList = append(
				outPutList,
				fmt.Sprintf(
					"%-12v%-20v%-10v",
					s.GetNumber(),
					v.GetNumber(),
					v.GetColour(),
				),
			)
		}
	} else {
		outPutList = []string{
			"No Data Found",
		}
	}
	st.OutPut = strings.Join(outPutList, "\n")
	return st.OutPut, nil
}
