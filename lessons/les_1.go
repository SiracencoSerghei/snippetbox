package lessons

import "fmt"


var city string = "Brussels"

func Les1Demo() {
	name := "Serghei"
	age := 48
	fmt.Println("Hello", name)
	fmt.Println("Age: ", age)
	fmt.Println("City: ", city)
	if age > 40 {
		fmt.Println("experienced developer")
	}
}

