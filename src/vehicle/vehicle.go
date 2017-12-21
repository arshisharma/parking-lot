// Package vehicle implements handling for vehicle objects
// A vehicle is the entity that occupies a slot in the parking lot
// Vehicle has  vehicle number and Color properties
// GetColour(), GetNumber() Defined to access properties independently.
package vehicle

type Vehicle struct {
	Number, Color string
}

func New(number, color string) *Vehicle {
	return new(Vehicle).Init(number, color)
}

// Init - Initialise created object
func (v *Vehicle) Init(number, color string) *Vehicle {
	v.Color = color
	v.Number = number
	return v
}

// GetColour -  Get value of  vehicle property colour
func (v *Vehicle) GetColour() string {
	return v.Color
}

// GetNumber - Get value of vehicle property number
func (v *Vehicle) GetNumber() string {
	return v.Number
}