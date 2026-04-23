package json_demo

import "encoding/json"

// EncodePerson = Go struct → JSON
func EncodePerson(p Person) ([]byte, error) {
	return json.Marshal(p)
}

// DecodePerson = JSON → Go struct
func DecodePerson(data []byte) (Person, error) {
	var p Person
	err := json.Unmarshal(data, &p)
	return p, err
}