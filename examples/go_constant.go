package main

import "fmt"

/*
Constants can be primitive types like strings, integers, booleans and floats.
They cannot be more complex types like slices, maps and structs.

As the name implies, the value of a constant can't be changed after it has been declared.
*/

func main() {
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
