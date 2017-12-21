// Package parking manages the complete logic of a parking ParkingLot
// Create, Add Vehicle, Remove vehicle, Search, Status Report etc.
package parking

import (
	"custerr"
	"slot"
	"vehicle"
)

// ParkingLot object holds the config of all required properties
type ParkingLot struct {
	Capacity,
	Counter,
	startSlotIndex uint64
	slots []*slot.Slot
}

// New parking ParkingLot instance
func New(capacity uint64) *ParkingLot {
	return new(ParkingLot).init(capacity)
}

// init parking ParkingLot instance
func (pl *ParkingLot) init(capacity uint64) *ParkingLot {
	pl.Capacity = capacity
	pl.startSlotIndex = 1
	pl.slots = make([]*slot.Slot, uint64(capacity))
  pl.Counter = 0
  start := 1
	for idx := range pl.slots {
		pl.slots[idx] = slot.New()
		pl.slots[idx].SetNumber(uint64(start))
		start = start + 1
	}
	return pl
}

// findNextFreeSlot to get next free slot
func (pl *ParkingLot) findNextFreeSlot() (*slot.Slot, error) {
  for _, objSlot := range pl.slots {
		if objSlot.IsFree() {
			return objSlot, nil
		}
	}
	return nil, custerr.ErrParkingFull
}

// getFreeSlot get next serial slot or free in between
func (pl *ParkingLot) getFreeSlot() (*slot.Slot, error) {
		sl,err :=  pl.findNextFreeSlot();if err != nil{
			return nil,custerr.ErrParkingFull
		}
    return sl,nil
}

// AddVehicle add vehicle to parking ParkingLot
func (pl *ParkingLot) AddVehicle(vehicle *vehicle.Vehicle) (*slot.Slot, error) {
	var (
		err     error
		objSlot *slot.Slot
	)
	objSlot, err = pl.getFreeSlot()
	if err == nil && objSlot != nil {
		objSlot, _ = objSlot.Allocate(vehicle)
		if err == nil {
			pl.Counter = pl.Counter + 1
		}
	}
	return objSlot, err
}

// remove remove vehicle from ParkingLot and decrement counter
func (pl *ParkingLot) remove(s *slot.Slot) {
	s.Free()
	pl.Counter = pl.Counter - 1
}

// RemoveVehicleBySlotNumber  remove slot from ParkingLot slot list by slot number
func (pl *ParkingLot) RemoveVehicleBySlotNumber(Number uint64) (*slot.Slot, error) {
	oSlot, err := pl.GetSlot(Number)
	if nil != err {
		return oSlot, err
	}

	pl.remove(oSlot)
	return oSlot, nil
}



// GetSlot from ParkingLot slot list by vehicle number
func (pl *ParkingLot) GetSlot(Number uint64) (*slot.Slot, error) {
	if Number < pl.Capacity && Number >= pl.startSlotIndex {
		return pl.slots[Number-pl.startSlotIndex], nil
	}
	return nil, custerr.ErrInvalidGetSlotNumber
}

// GetSlotsBy vehicle property { number, color }
func (pl *ParkingLot) GetSlotsBy(property, value string) ([]*slot.Slot, error) {
	var arrSlots = make([]*slot.Slot, 0)
	var val1,val2 string
  lindex := 0
	rindex := len(pl.slots) - 1

	for lindex <= rindex {
		vl := pl.slots[lindex].GetVehicle()
    vr := pl.slots[rindex].GetVehicle()
		if nil != vl && nil != vr {
			switch property {
			case "number":
				val1 = vl.GetNumber()
        val2 = vr.GetNumber()
				break
			case "color":
				val1 = vl.GetColour()
        val2 = vr.GetColour()
				break
			}
			if value == val1 {
				arrSlots = append(arrSlots, pl.slots[lindex])
			}
      if value == val2 {
        arrSlots = append(arrSlots, pl.slots[rindex])
      }
		}else if nil != vl{
      switch property {
			case "number":
				val1 = vl.GetNumber()
				break
			case "color":
				val1 = vl.GetColour()
				break
			}
			if value == val1 {
				arrSlots = append(arrSlots, pl.slots[lindex])
			}

    }else if nil != vr{
      switch property {
			case "number":
				val1 = vr.GetNumber()
				break
			case "color":
				val1 = vr.GetColour()
				break
			}
			if value == val1 {
				arrSlots = append(arrSlots, pl.slots[rindex])
			}
    }
		lindex = lindex + 1
		rindex = rindex - 1
	}
	return arrSlots, nil
}
