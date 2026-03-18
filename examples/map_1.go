package examples

import "fmt"

// Vertex is a struct type that groups latitude and longitude together
type Vertex struct {
	Lat, Long float64
	// Lat  = latitude (float64 number)
	// Long = longitude (float64 number)
}

var m map[string]Vertex

func BuildMap() map[string]Vertex {
	return map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
	}
}

func Map1Demo() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Printf("%T\n", m["Lat"])
	fmt.Printf("%T\n", m["Bell Labs"].Lat)
	// fmt.Printf("%T\n", m.Lat) // invalid operation: m.Lat (m is of type map[string]Vertex) Because m is a map, not a struct.

}
// {40.68433 -74.39967}
// examples.Vertex
// float64

// Golden rule (memorize)

// Dot (.) works ONLY on structs and methods
// Brackets ([]) work on maps, slices, arrays, strings