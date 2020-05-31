package main

import (
	"fmt"

	"github.com/kpunith8/gosamples/utils"
	"github.com/kpunith8/gosamples/utils/consts"

	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
)

func main() {
	var val int
	val3 := 3
	integerVal := 42
	floatVal := float64(integerVal)
	uIntVal := uint(floatVal)

	fmt.Println("Random number:", rand.Intn(100))

	// A name is exported if it begins with a capital letter

	fmt.Println("Pi value is:", math.Pi)
	fmt.Println(":Functions:")
	fmt.Println("Sum is:", add(10, 15))

	a, b := swap("K", "Punith")
	fmt.Println("Swapped variables", a, b)

	fmt.Println(":Named return values:")
	fmt.Println(split(35))

	fmt.Println(":Variables:", val, val1, val2, val3)

	fmt.Println(":Basic Types:")
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)

	fmt.Println(":Type conversions:")

	fmt.Println(floatVal, uIntVal)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum using for loop:", sum)

	// The init and post statements are optional.
	// It can also be used as `while`
	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	// for {}
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println("For without post and init statements, can be used as while:", sum1)

	// Using if
	fmt.Println("Square of a number: sqrt(25):", sqrt(25), "sqrt(-3):", sqrt(-3))

	fmt.Println("If with a short statement:", powerOfX(3, 2, 10),
		powerOfX(3, 3, 20))
	fmt.Println("If Else:")
	fmt.Println(pow1(3, 2, 10),
		pow1(3, 3, 20))

	fmt.Println(":Switch:")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch without a condition is the same as switch true.

	defer fmt.Println("Defer")

	// Deferred function calls are pushed onto a stack. When a function returns, its
	// deferred calls are executed in last-in-first-out order
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// defer fmt.Println(i)
	}

	fmt.Println("done")

	fmt.Println("Importing module:")
	add1 := utils.Adder()
	fmt.Println(add1(10))

	fmt.Println("Getname nested package:", consts.Getname())
}

func add(x int, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Variables
// The `var` statement declares a list of variables; as in function argument lists, the `type` is last.
// A var statement can be at package or function level.
var c, python, java bool

// Variable Initializers
var val1, val2 int = 12, 22

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var c1, python1, java1 = true, false, "no!"

// variable declarations may be "factored" into blocks, as with import statements.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Constants

// Pi 3.14
const Pi = 3.14

// Truth true
const Truth = true

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func powerOfX(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

func pow1(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	// can't use `v` here
	return limit
}
