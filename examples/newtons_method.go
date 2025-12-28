package main

import (
	"fmt"
	"math"
)

// Sqrt повертає наближене значення квадратного кореня з x
// за допомогою методу Ньютона.
func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN() // для некоректного вводу
	}

	z := 1.0
	const eps = 1e-10

	for {
		next := z - (z*z-x)/(2*z)

		// перевірка умови виходу
		if math.Abs(next-z) < eps {
			return next
		}

		z = next
	}
}

func main() {
	x := 2.0

	fmt.Println("Newton:", Sqrt(x))
	fmt.Println("math.Sqrt:", math.Sqrt(x))
}

