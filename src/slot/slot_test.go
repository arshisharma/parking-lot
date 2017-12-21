package slot

import (
	"custerr"
	"testing"
	"vehicle"
)


var slots = []Slot{
	{0, vehicle.New("ABC", "Red")},
	{1, vehicle.New("DEF", "White")},
	{2, vehicle.New("XYZ", "Red")},
}

func TestSlotInit(t *testing.T) {
	var s *Slot = New()
	if s.Number != 0 {
		t.Error("Slot number not initiazlized should be zero")
	}
}

func TestSlotSetNumber(t *testing.T) {
	var (
		s   *Slot
		err error
	)

	s = New()
	_, err = s.SetNumber(0)
	if err == nil {
		t.Error("Invalid Slot number")
	}
	_, err = s.SetNumber(1)
	if err != nil {
		t.Error("Error when adding valid slot number",err)
	}
}

func TestSlot_Allocate(t *testing.T) {
	var s *Slot

	for _, o := range slots {
		s = New()
		_, err := s.Allocate(o.Vehicle)
		if err != custerr.ErrInvalidSlotAllocate {
			t.Error("Invalid slot")
		}

		_, err = s.SetNumber(o.Number)
		if o.Number < SlotNumberLowerLimit {
			if err != custerr.ErrSlotNumberInvalid {
				t.Error("Expected invalid slot number but got",err)
			}
		}

		_, err = s.Allocate(o.Vehicle)

		if o.Number < 1 {
			if err != custerr.ErrInvalidSlotAllocate {
				t.Error("Invalid slot")
			}
		} else {
			if s.GetVehicle() != o.Vehicle {
				t.Error("Same Vehicle")
			}
		}

		if !s.IsFree() {
			_, err = s.Allocate(o.Vehicle)
			if err != custerr.ErrSlotAlreadyAllocated {
				t.Error("Expected err slot allocated but got",err)
			}
		}

		if !s.IsFree() {
			s.Free()
			if !s.IsFree() {
				t.Error("Slot is free")
			}
		}
	}
}
