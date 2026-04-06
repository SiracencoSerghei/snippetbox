package examples

import "testing"

func TestCarDrive(t *testing.T) {
	car := Car{Brand: "Toyota", Model: "Corolla", Mileage: 100}

	car.Drive(50)

	if car.Mileage != 150 {
		t.Errorf("expected mileage 150, got %d", car.Mileage)
	}
}

func TestNestedStruct(t *testing.T) {
	car := CarWithEngine{
		Brand: "BMW",
		Engine: Engine{
			Power: 200,
			Type:  "diesel",
		},
	}

	if car.Engine.Power != 200 {
		t.Errorf("expected 200, got %d", car.Engine.Power)
	}
}

func TestEmbedding(t *testing.T) {
	car := ElectricCar{
		Brand: "Tesla",
		Battery: Battery{
			Capacity: 75,
		},
	}

	if car.Capacity != 75 {
		t.Errorf("expected 75, got %d", car.Capacity)
	}
}

func TestAnonymousStruct(t *testing.T) {
	user := NewAnonymousUser("John", 30)

	u, ok := user.(struct {
		Name string
		Age  int
	})

	if !ok {
		t.Errorf("type assertion failed")
	}

	if u.Name != "John" || u.Age != 30 {
		t.Errorf("unexpected values")
	}
}

func TestPointerVsValue(t *testing.T) {
	car := Car{Mileage: 100}

	IncreaseMileageValue(car, 50)
	if car.Mileage != 100 {
		t.Errorf("value receiver should NOT modify original")
	}

	IncreaseMileagePointer(&car, 50)
	if car.Mileage != 150 {
		t.Errorf("pointer receiver SHOULD modify original")
	}
}