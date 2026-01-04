package main

import "fmt"

// func main() {
// 	const name = "Saul Goodman"
// 	const openRate = 30.5

// 	msg := fmt.Sprintf("Hi %v, your open rate is %.1f percent\n",name,openRate)

// 	// don't edit below this line

// 	fmt.Print(msg)
// }

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	userLog := fmt.Sprintf(
		"Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s",
		fname, lname, age, messageRate, isSubscribed, message,
	)

	fmt.Println(userLog)
}

