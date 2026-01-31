package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	finalCost = bulkMessageCost

	switch {
	case isPremiumUser:// This must run FIRST so we calculate the discounted price
		finalCost = finalCost * (1 - discountRate)

		fallthrough // continue to the next case WITHOUT re-checking condition

	case accountBalance >= finalCost:
		accountBalance -= finalCost
		fmt.Println(purchaseSuccessMessage)
	default:
		fmt.Println(insufficientFundMessage)
	}

	fmt.Println("Account balance:", accountBalance)
}
