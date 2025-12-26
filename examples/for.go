package main

import "fmt"

func main() {
	// Цикл for у Go має три компоненти:
	// 	init-інструкція (init statement): виконується перед першою ітерацією
	//  вираз умови (condition expression): обчислюється перед кожною ітерацією
	//  post-інструкція (post statement): виконується наприкінці кожної ітерації

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Init- та post-інструкції є необов’язковими.
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// те, що в C називається while, у Go пишеться як for
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)
}
