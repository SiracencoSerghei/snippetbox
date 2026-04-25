package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
Area() float64
}
type Circle struct {
radius float64
}
type Square struct {
length float64
}
type Triangle struct {
base float64
height float64
}
// Circle implements Shape
func (c Circle) Area() float64 {
return math.Pi * c.radius * c.radius
}
// Square implements Shape
func (s Square) Area() float64 {
return s.length * s.length
}
func calculateArea(listOfShapes []Shape) {
for _, shape := range listOfShapes {
fmt.Println("Area: ", shape.Area())
}
}

func (t Triangle) Area() float64 {
return 0.5 * t.base * t.height
}
