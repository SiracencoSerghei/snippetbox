// Інструкції if у Go схожі на цикли for: вираз не потрібно брати в круглі дужки (`), але фігурні дужки { } — обов’язкові.

package main

import (
	"fmt"
	"math"
	"reflect"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(reflect.TypeOf(sqrt(2)), reflect.TypeOf(sqrt(-4)))
}
