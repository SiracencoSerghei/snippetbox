package data_struct

import (
	"fmt"
	"testing"
)

func TestSortPersonsByName(t *testing.T) {
	list := []Person{
		{"Mary Smith", "mary@example.com", Dob{17, 3, 1990}},
		{"John Smith", "john@example.com", Dob{9, 12, 1988}},
		{"Janet Doe", "janet@example.com", Dob{1, 12, 1988}},
	}

	fmt.Println("Before sorting:")
	for _, p := range list {
		fmt.Printf("%s | %s | %+v\n", p.Name, p.Email, p.Dob)
	}

	SortPersons(list, "name")

	expected := []string{
		"Janet Doe",
		"John Smith",
		"Mary Smith",
	}

	for i, p := range list {
		if p.Name != expected[i] {
			t.Fatalf("expected %s, got %s", expected[i], p.Name)
		}
	}

	fmt.Println("After sorting:")
	for _, p := range list {
		fmt.Printf("%s | %s | %+v\n", p.Name, p.Email, p.Dob)
	}
}

func TestSortPersonsByDate(t *testing.T) {
	list := []Person{
		{"Mary Smith", "mary@example.com", Dob{17, 3, 1990}},
		{"John Smith", "john@example.com", Dob{9, 12, 1988}},
		{"Janet Doe", "janet@example.com", Dob{1, 12, 1988}},
	}

	fmt.Println("Before sorting:")
	for _, p := range list {
		fmt.Printf("%s | %s | %+v\n", p.Name, p.Email, p.Dob)
	}

	SortPersons(list, "date")

	expected := []string{
		"Janet Doe",
		"John Smith",
		"Mary Smith",
	}

	for i, p := range list {
		if p.Name != expected[i] {
			t.Fatalf("date sort failed: expected %s, got %s", expected[i], p.Name)
		}
	}

	fmt.Println("After sorting:")
	for _, p := range list {
		fmt.Printf("%s | %s | %+v\n", p.Name, p.Email, p.Dob)
	}
}


// go test ./examples/data_struct -v -run "TestSortPersonsBy(Name|Date)"