package json_demo

import (
	"fmt"
	"testing"
)

func TestEncodePerson(t *testing.T) {
	person := Person{
		Name:  "Mary Smith",
		Email: "mary@example.com",
		Dob:   Dob{17, 3, 1990},
	}

	fmt.Println("\n--- ORIGINAL GO STRUCT ---")
	fmt.Printf("%+v\n", person)

	jsonData, err := EncodePerson(person)
	if err != nil {
		t.Fatalf("encode failed: %v", err)
	}

	fmt.Println("\n--- ENCODED JSON ---")
	fmt.Println(string(jsonData))

	if len(jsonData) == 0 {
		t.Fatal("empty JSON result")
	}
}

func TestDecodePerson(t *testing.T) {
	jsonData := []byte(`{
		"Name":"John Smith",
		"Email":"john@example.com",
		"Dob":{"Day":9,"Month":12,"Year":1988}
	}`)

	fmt.Println("\n--- RAW JSON ---")
	fmt.Println(string(jsonData))

	person, err := DecodePerson(jsonData)
	if err != nil {
		t.Fatalf("decode failed: %v", err)
	}

	fmt.Println("\n--- DECODED STRUCT ---")
	fmt.Printf("%+v\n", person)

	if person.Name != "John Smith" {
		t.Fatalf("expected John Smith, got %s", person.Name)
	}
}