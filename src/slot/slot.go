// Package slot implements a simple slot object for vehicles.
// It defines a type Slot with following methods.
// SetNumber(), IsValid(), Allocate(), ISFree(), Free(), GetVehicle()
package slot

import (
	"custerr"
	"vehicle"
)

// SlotNumberLowerLimit is slot lower bound defined as a constant.
const SlotNumberLowerLimit = 1

// Slot defines a Number and a Vehicle.
// If vehicle object is allocated then slot is used
type Slot struct {
	Number  uint64
	Vehicle *vehicle.Vehicle
}

// New Slot Object creation function
func New() *Slot {
	return new(Slot).init()
}

// initialise Object with default values
func (sl *Slot) init() *Slot {
	sl.Number = SlotNumberLowerLimit - 1
	sl.Vehicle = nil
	return sl
}

// SetNumber Set slot number to slot object
func (sl *Slot) SetNumber(number uint64) (*Slot, error) {
	if !IsValidSlotNumber(number) {
		return sl, custerr.ErrSlotNumberInvalid
	}
	sl.Number = number
	return sl, nil
}

// GetNumber get slot number from slot object
func (sl *Slot) GetNumber() uint64 {
	return sl.Number
}

// IsValidSlotNumber help to check the slot number is valid or not
func IsValidSlotNumber(Number uint64) bool {
	return (Number >= SlotNumberLowerLimit)
}

//Allocates the vehicle to the slot
func (sl *Slot) Allocate(vehicle *vehicle.Vehicle) (*Slot, error) {
	if !IsValidSlotNumber(sl.Number) {
		return sl, custerr.ErrInvalidSlotAllocate
	}

	if nil != sl.Vehicle {
		return sl, custerr.ErrSlotAlreadyAllocated
	}

	sl.Vehicle = vehicle
	return sl, nil
}

// GetVehicle get vehicle object from allocated slot.
func (sl *Slot) GetVehicle() *vehicle.Vehicle {
	return sl.Vehicle
}

// Free remove vehicle object from slot
func (sl *Slot) Free() *Slot {
	sl.Vehicle = nil
	return sl
}

// IsFree Verifies that slot is free or not, if no vehicle allocated
func (sl *Slot) IsFree() bool {
	return sl.Vehicle == nil
}
