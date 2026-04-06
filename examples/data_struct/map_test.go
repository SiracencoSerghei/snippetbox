package examples

import "testing"

func TestBuildMap(t *testing.T) {
	m := BuildMap()

	v, ok := m["Bell Labs"]
	if !ok {
		t.Fatal("key not found")
	}

	if v.Lat != 40.68433 || v.Long != -74.39967 {
		t.Errorf("wrong value: got %+v", v)
	}
}

func TestAddToMap(t *testing.T) {
	m := BuildMap()

	m = AddToMap(m, "Google", Vertex{1.23, 4.56})

	v, ok := m["Google"]
	if !ok {
		t.Fatal("key not added")
	}

	if v.Lat != 1.23 || v.Long != 4.56 {
		t.Errorf("wrong value: got %+v", v)
	}
}