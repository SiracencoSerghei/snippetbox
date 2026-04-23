package interfaces

import (
    "fmt"
    "testing"
)

func TestPersonSpeak(t *testing.T) {
    p := Person{Name: "John"}

    result := SaySomething(p)

    fmt.Println("Person says:", result)

    if result != "Hello, my name is John" {
        t.Fatalf("wrong result: %s", result)
    }
}

func TestDogSpeak(t *testing.T) {
    d := Dog{}

    result := SaySomething(d)

    fmt.Println("Dog says:", result)

    if result != "Woof" {
        t.Fatalf("wrong result: %s", result)
    }
}