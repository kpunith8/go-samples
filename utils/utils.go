package utils

// Adder function closures
func Adder() func(int) int { // func(int) int, is the return type of the func adder()
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
