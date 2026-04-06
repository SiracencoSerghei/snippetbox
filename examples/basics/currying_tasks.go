package examples

// TASK 1
// makeMultiplier повертає функцію, яка множить на x
// приклад:
// double := makeMultiplier(2)
// double(5) → 10

func makeMultiplier(x int) func(int) int {
	// TODO: реалізуй
	return func(y int) int {
        return x * y
    }
}

// TASK 2
// makeAdder повертає функцію, яка додає x
// приклад:
// add10 := makeAdder(10)
// add10(5) → 15

func makeAdder(x int) func(int) int {
	// TODO: реалізуй
	return func(y int) int {
        return x + y
    }
}

// TASK 3 (складніше)
// applyTwice приймає функцію і повертає нову,
// яка застосовує її ДВА рази
// приклад:
// double := func(x int) int { return x * 2 }
// f := applyTwice(double)
// f(3) → 12 (3 * 2 * 2)
func applyTwice(f func(int) int) func(int) int {
	// TODO: реалізуй
	return func(x int) int {
		return f(f(x))
	}
}



