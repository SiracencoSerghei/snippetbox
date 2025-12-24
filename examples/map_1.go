package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
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
// main.Vertex
// float64

// Golden rule (memorize)

// Dot (.) works ONLY on structs and methods
// Brackets ([]) work on maps, slices, arrays, strings