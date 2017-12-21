package parking

import (
	"testing"
)

const (
	NoSlot = "Should be one slot free"
	Slot11 = "Slot Number should be 11"

	Color        = "Color 12"
	NoColor12    = "Color 12 should present"
	NoColor12Val = "Colour value should be 'Color 12'"

	NoNumber12    = "Number 12 should present"
	NoNumber12Val = "Number value should be 'KL 12'"
)

var (
	start           uint64    = 1
	index, capacity uint64 = 0, 100
	pC                       = New(capacity)
	tVehicle                 = testVehicle(uint64(12))
)

func AddDataForSearch() {
	for ; uint64(index) < uint64(capacity); index++ {
		pC.AddVehicle(testVehicle(uint64(index)))
	}
}

func TestParkingCenterSearchVehicleByColor(t *testing.T) {
	AddDataForSearch()
	oSlots, _ := pC.SearchVehicleByColor(tVehicle.GetColour())
	if 1 != len(oSlots) {
		t.Error("Incorrect color")
	} else if oSlots[0].Vehicle.Color != tVehicle.GetColour() {
		t.Error("Incorrect color")
	}
}

func TestParkingCenterSearchVehicleByNumber(t *testing.T) {
	AddDataForSearch()
	index := uint64(33)
	oSlots, _ := pC.RemoveVehicleBySlotNumber(index)
	if nil == oSlots {
		t.Error("Incorrect slot number",oSlots)
	} else if oSlots.GetNumber() != index {
		t.Error("Slot numbers dont match")
	}
}
