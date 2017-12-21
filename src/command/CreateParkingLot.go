// package command with 'create_parking_lot' command implementation
package command

import (
	"fmt"
	"parking"
	"custerr"
	"strconv"
)

// commandCreateParkingLot defined arguments and related methods
type commandCreateParkingLot struct {
	Command
	Capacity uint64
}

const (
	Start =  1
)
// NewcommandCreateParkingLot new command instance
func NewcommandCreateParkingLot() *commandCreateParkingLot {
	var command = new(commandCreateParkingLot)
	command.command = "create_parking_lot"
	return command
}

// Help to print help of 'create_parking_lot' command
func (ccp *commandCreateParkingLot) Help() string {
	return `🔸  create_parking_lot <slots count>
	Create parking lot slots.
	Eg: create_parking_lot 6`
}

// Parse to parse arguments
func (ccp *commandCreateParkingLot) Parse(argString string) error {
	ccp.Command.Parse(argString)
	if "" != ccp.Args[0] {
		val, err := strconv.ParseUint(ccp.Args[0], 0, 64)
		if nil != err {
			return custerr.ErrInvalidParams
		}
		ccp.Capacity = uint64(val)
	}
	return nil
}

// Verify to check the provided parameters are valid or not
func (ccp *commandCreateParkingLot) Verify() error {
	if 1 > ccp.Capacity {
		return custerr.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (ccp *commandCreateParkingLot) Run() (string, error) {
	pL = parking.New(ccp.Capacity)
	if nil != pL {
		ccp.OutPut = fmt.Sprintf(
			"Created a parking lot with %v slots",
			ccp.Capacity,
		)
	} else {
		ccp.OutPut = custerr.ErrCreateParkingCenter.Error()
	}
	return ccp.OutPut, nil
}
