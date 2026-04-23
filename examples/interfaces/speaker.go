package interfaces

// Interface

type Speaker interface {
    Speak() string
}

// Structs

type Person struct {
    Name string
}

type Dog struct {}

// Methods

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

func (d Dog) Speak() string {
    return "Woof"
}

// Function using interface

func SaySomething(s Speaker) string {
    return s.Speak()
}