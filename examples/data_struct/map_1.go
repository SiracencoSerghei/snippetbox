package data_struct

import (
	"fmt"
)

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

// ==================================

// MAP OF STRUCTS
// http://localhost:4000/demo/map

// ==================================

func MapOfStructs() string {
	type dob struct {
		day int
		month int
		year int
	}
	type people struct {
		name string
		email string
		dob dob
	}

	members := make(map[int]people)

	members[1] = people{
		name: "Mary Smith",
		email: "marysmith@example.com",
		dob: dob{
		day:
		17,
		month: 3,
		year: 1990,
		},
	}
	members[2] = people{
		name: "John Smith",
		email: "johnsmith@example.com",
		dob: dob{
		day:
		9,
		month: 12,
		year: 1988,
		},
	}
		members[3] = people{
		name: "Janet Doe",
		email: "janetdoe@example.com",
		dob: dob{
		day:
		1,
		month: 12,
		year: 1988,
		},
	}
		members[4] = people{
		name: "Adam Jones",
		email: "adamjones@example.com",
		dob: dob{
		day:
		19,
		month: 8,
		year: 2001,
		},
	}

	var out string

	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
		out += fmt.Sprintf(
			"%d: %s, %s, %d/%d/%d\n",
			k,
			v.name,
			v.email,
			v.dob.day,
			v.dob.month,
			v.dob.year,
		)
	}

	return out
}