package main

import (
	"fmt"
	"sort"
)

type Vertex struct {
	Lat, Long float64
}

func main() {

	// ----------------------------
	// 1. MAP OF VALUES (COPY TRAP)
	// ----------------------------
	fmt.Println("== Map of VALUES ==")

	mv := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
	}

	// mv["Bell Labs"].Lat = 50 // ❌ compile error

	v := mv["Bell Labs"]
	v.Lat = 50
	mv["Bell Labs"] = v

	fmt.Println(mv["Bell Labs"])

	// ----------------------------
	// 2. MAP OF POINTERS (NO COPY)
	// ----------------------------
	fmt.Println("\n== Map of POINTERS ==")

	mp := map[string]*Vertex{
		"Google": {37.42202, -122.08408},
	}

	mp["Google"].Lat = 100 // ✅ works
	fmt.Println(*mp["Google"])

	// ----------------------------
	// 3. MISSING KEY TRAP
	// ----------------------------
	fmt.Println("\n== Missing key ==")

	x := mv["Unknown"]
	fmt.Println(x) // zero value

	if _, ok := mv["Unknown"]; !ok {
		fmt.Println("key does not exist")
	}
	fmt.Println(mv)
	fmt.Println(x)

	fmt.Println(len(mv))
	_ = mv["Unknown"]
	fmt.Println(len(mv))

	fmt.Println(len(mv))
	mv["Unknown"] = Vertex{1, 2}
	fmt.Println(len(mv))

	// ----------------------------
	// 4. RANDOM ITERATION
	// ----------------------------
	fmt.Println("\n== Random iteration ==")

	for k, v := range mv {
		fmt.Println(k, v)
	}

	// ----------------------------
	// 5. ORDERED ITERATION
	// ----------------------------
	fmt.Println("\n== Ordered iteration ==")

	keys := make([]string, 0, len(mv))
	fmt.Println(keys)
	for k := range mv {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, mv[k])
	}
	fmt.Println(keys)

	// ----------------------------
	// 6. SLICE vs MAP
	// ----------------------------
	fmt.Println("\n== Slice vs Map ==")

	s := []string{"a", "b", "c"}
	m := map[string]bool{"a": true, "b": true}

	fmt.Println("slice lookup:", contains(s, "b"))
	fmt.Println("map lookup:", m["b"])
}

// O(n)
func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}
