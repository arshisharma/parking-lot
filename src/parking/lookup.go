// Package parking manages the complete logic of a parking ParkingLot
// Search, Status Report etc. are implemented here
package parking

import (
	"custerr"
	"slot"
)

// SearchVehicleByColor Get filled slots with vehicle colour
func (pl *ParkingLot) SearchVehicleByColor(number string) ([]*slot.Slot, error) {
	oSlots, err := pl.GetSlotsBy("color", number)
	if 0 == len(oSlots) {
		err = custerr.ErrNotFound
	}
	return oSlots, err
}

// ReportVehicleByNumber Get filled slots with vehicle number
func (pl *ParkingLot) ReportVehicleByNumber(number string) ([]*slot.Slot, error) {
	oSlots, err := pl.GetSlotsBy("number", number)
	if 0 == len(oSlots) {
		err = custerr.ErrNotFound
	}
	return oSlots, err
}

// ReportFilledSlots Get all filled slots
func (pl *ParkingLot) ReportFilledSlots() ([]*slot.Slot, error) {
	allocSlots := make([]*slot.Slot, 0)
	if uint64(0) == pl.Counter {
		return nil, custerr.ErrNoFilledSlots
	}
	for _, s := range pl.slots {
		if !s.IsFree() {
			allocSlots = append(allocSlots, s)
		}
	}
	return allocSlots, nil
}
