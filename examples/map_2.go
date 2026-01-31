package main
// Declares the main package.
// Every executable Go program must be in package main.

import (
	"fmt"
	// fmt provides formatted I/O functions like Println

	"sort"
	// sort provides utilities for sorting slices
)

// Vertex is a struct type that groups latitude and longitude together
type Vertex struct {
	Lat, Long float64
	// Lat  = latitude (float64 number)
	// Long = longitude (float64 number)
}

func main() {
	// main is the entry point of the program
	// Code execution starts here

	// ----------------------------
	// 1. MAP OF VALUES (COPY TRAP)
	// ----------------------------
	fmt.Println("== Map of VALUES ==")
	// Prints a section title

	mv := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
	}
	// mv is a map:
	// key   = string
	// value = Vertex (struct value, NOT pointer)
	// "Bell Labs" maps to a Vertex with Lat and Long

	// mv["Bell Labs"].Lat = 50
	// ❌ This would cause a compile error
	// Because map values are copied when accessed,
	// you cannot modify struct fields directly in a map

	v := mv["Bell Labs"]
	// Retrieves a COPY of the Vertex stored in the map

	v.Lat = 50
	// Modifies the copy (not the map value)

	mv["Bell Labs"] = v
	// Writes the modified copy back into the map

	fmt.Println(mv["Bell Labs"])
	// Prints the updated value from the map

	// ----------------------------
	// 2. MAP OF POINTERS (NO COPY)
	// ----------------------------
	fmt.Println("\n== Map of POINTERS ==")
	// \n adds a blank line before the text

	mp := map[string]*Vertex{
		"Google": {37.42202, -122.08408},
	}
	// mp is a map where values are POINTERS to Vertex
	// No copying happens when accessing map elements

	mp["Google"].Lat = 100
	// ✅ Works because we modify the struct via a pointer

	fmt.Println(*mp["Google"])
	// Dereferences the pointer and prints the actual struct

	// ----------------------------
	// 3. MISSING KEY TRAP
	// ----------------------------
	fmt.Println("\n== Missing key ==")

	x := mv["Unknown"]
	// Accessing a non-existing key
	// Returns the ZERO VALUE of Vertex (Lat=0, Long=0)

	fmt.Println(x)
	// Prints {0 0}

	if _, ok := mv["Unknown"]; !ok {
		// The "comma ok" idiom:
		// ok == false means the key does not exist
		fmt.Println("key does not exist")
	}

	fmt.Println(mv)
	// Prints the map (still unchanged)

	fmt.Println(x)
	// x is still the zero-value Vertex

	fmt.Println(len(mv))
	// len() returns number of key-value pairs in the map

	_ = mv["Unknown"]
	// Reads the key but does nothing with it
	// Does NOT add the key to the map

	fmt.Println(len(mv))
	// Length is unchanged

	fmt.Println(len(mv))
	mv["Unknown"] = Vertex{1, 2}
	// Assigning a value ADDS a new key to the map

	fmt.Println(len(mv))
	// Length increases by 1

	// ----------------------------
	// 4. RANDOM ITERATION
	// ----------------------------
	fmt.Println("\n== Random iteration ==")

	for k, v := range mv {
		// range over a map iterates in RANDOM order
		fmt.Println(k, v)
	}
	// Go deliberately randomizes map iteration order

	// ----------------------------
	// 5. ORDERED ITERATION
	// ----------------------------
	fmt.Println("\n== Ordered iteration ==")

	keys := make([]string, 0, len(mv))
	// Creates an empty slice of strings
	// length = 0, capacity = len(mv)

	fmt.Println(keys)
	// Prints empty slice: []

	for k := range mv {
		// Iterates over map keys only
		keys = append(keys, k)
		// Adds each key to the slice
	}

	sort.Strings(keys)
	// Sorts the slice of keys alphabetically

	for _, k := range keys {
		// Iterates over sorted keys
		fmt.Println(k, mv[k])
		// Accesses map values in a predictable order
	}

	fmt.Println(keys)
	// Prints the sorted keys slice

	// ----------------------------
	// 6. SLICE vs MAP
	// ----------------------------
	fmt.Println("\n== Slice vs Map ==")

	s := []string{"a", "b", "c"}
	// Slice of strings (ordered list)

	m := map[string]bool{"a": true, "b": true}
	// Map used as a set for fast lookup

	fmt.Println("slice lookup:", contains(s, "b"))
	// Checks if "b" exists in slice (O(n) time)

	fmt.Println("map lookup:", m["b"])
	// Checks if "b" exists in map (O(1) average time)
}

// contains checks if a value exists in a slice
// Time complexity: O(n)
func contains(s []string, v string) bool {
	for _, x := range s {
		// Iterates through each element in the slice
		if x == v {
			// Found the value
			return true
		}
	}
	// Value not found
	return false
}
