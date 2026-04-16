package data_struct

// BASIC STRUCT
type Car struct {
	Brand   string
	Model   string
	Mileage int
}

// METHOD (pointer receiver)
func (c *Car) Drive(km int) {
	c.Mileage += km
}

// NESTED STRUCT
type Engine struct {
	Power int
	Type  string
}

type CarWithEngine struct {
	Brand  string
	Engine Engine
}

// EMBEDDING (composition)
type Battery struct {
	Capacity int
}

type ElectricCar struct {
	Brand string
	Battery
}

// ANONYMOUS STRUCT FACTORY
func NewAnonymousUser(name string, age int) interface{} {
	return struct {
		Name string
		Age  int
	}{
		Name: name,
		Age:  age,
	}
}

// POINTER VS VALUE DEMO
func IncreaseMileageValue(c Car, km int) {
	c.Mileage += km
}

func IncreaseMileagePointer(c *Car, km int) {
	c.Mileage += km
}