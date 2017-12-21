package parking

import (
	. "custerr"
	. "strconv"
	"testing"
	"vehicle"
)

func testVehicle(index uint64) *vehicle.Vehicle {
	struint64 := FormatUint(uint64(index), 10)
	return vehicle.New("KL "+struint64, "Color "+struint64)
}

func testAddVehicle(index uint64, pL *ParkingLot, t *testing.T) {
	oSlot, err := pL.AddVehicle(testVehicle(index))
	if nil != err {
		t.Error("FAIL: Expected err should be nil but got err:",err)
	}
	if nil == oSlot {
		t.Error("FAIL: Expected slot to be allocated but got nil")
	}
}

func testLeaveVehicleByNumber(index uint64, pL *ParkingLot, t *testing.T) {
	oSlots, err := pL.RemoveVehicleBySlotNumber(index)

	if nil != err {
		t.Error("FAIL: Expected err should be nil but got err:",err,index)
	}

	if nil == oSlots {
		t.Error("FAIL: Expected slot to be free but got nil")
	}
}

func testAddVehicleWhenParkingFull(index uint64, pL *ParkingLot, t *testing.T) {
	oSlot, err := pL.AddVehicle(testVehicle(index))
	if nil != oSlot {
		t.Error("FAIL: Expected slot not to be allocated but got allocated:")
	}
	if ErrParkingFull != err {
		t.Error("FAIL: Expected err parking full but got :",err)
	}
}

func TestParkingLot_Functional(t *testing.T) {
	var (
		index,capacity uint64 = 1,100
		pL                       = New(capacity)
	)
	for ; uint64(index) <= uint64(capacity); index++ {
		testAddVehicle(uint64(index), pL, t)
	}
	testAddVehicleWhenParkingFull(uint64(index), pL, t)

	testLeaveVehicleByNumber(uint64(10), pL, t)
	testAddVehicle(uint64(10), pL, t)

	for index = 1; uint64(index) < uint64(capacity); index++ {
		testLeaveVehicleByNumber(uint64(index), pL, t)
	}

}

func BenchmarkParkingLotAddVehicle(b *testing.B) {
	var (
		index, capacity uint64 = 0, 10000000000000
		pL                       = New(capacity)
	)
	for ; uint64(index) < uint64(capacity); index++ {
		pL.AddVehicle(testVehicle(uint64(index)))
	}
}

func BenchmarkParkingLot_AddAndRemove(b *testing.B) {
	var (
		index, capacity uint64 = 0, 100000000000
		pL                       = New(capacity)
	)
	for ; uint64(index) < uint64(capacity); index++ {
		pL.AddVehicle(testVehicle(uint64(index)))
	}
	for ; uint64(index) < uint64(capacity); index++ {
		pL.RemoveVehicleBySlotNumber(uint64(index))
	}
}
