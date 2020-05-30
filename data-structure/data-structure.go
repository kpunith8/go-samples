package main

import (
	"fmt"
	"math"
	"strings"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	i, j := 42, 2701

	fmt.Println(":Pointers:")
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	vert1 := Vertex{1, 2}
	vert1.X = 4
	fmt.Println(":Structs:", vert1.X)

	vert2 := Vertex{1, 2}
	p1 := &vert2
	p1.X = 1e9
	fmt.Println(":Pointers to struct:", vert2)

	fmt.Println(":Struct Literals:", v1, v2, v3, p2)

	fmt.Println(":Arrays:")
	var arr1 [2]string
	arr1[0] = "Punith"
	arr1[1] = "K"
	fmt.Println(arr1[0], arr1[1])
	fmt.Println(arr1)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var arr1Slice []int = primes[1:4]
	fmt.Println("Array Slices:", arr1Slice)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)

	fmt.Println("Slice literals:")
	q1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q1)

	r1 := []bool{true, false, true, true, false, true}
	fmt.Println(r1)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)

	fmt.Println("Slice length and capacity:")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	fmt.Println("Create a slice with make():")
	a1 := make([]int, 5)
	printSlice1("a1", a1)

	b1 := make([]int, 0, 5)
	printSlice1("b1", b1)

	c1 := b1[:2]
	printSlice1("c1", c1)

	d1 := c1[2:5]
	printSlice1("d1", d1)

	fmt.Println("Slices of slices:")
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("Appending values to slice:")
	var s2 []int
	printSlice(s2)

	// append works on nil slices.
	s2 = append(s2, 0)
	printSlice(s2)

	// The slice grows as needed.
	s2 = append(s2, 1)
	printSlice(s2)

	// We can add more than one element at a time.
	s2 = append(s2, 2, 3, 4)
	printSlice(s2)

	fmt.Println("range function:")
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("Skip index from range function")
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	fmt.Println(":Maps:")
	map1 = make(map[string]Vertex1)
	map1["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(map1["Bell Labs"])

	fmt.Println("Map Literals with no types:")
	var map2 = map[string]Vertex1{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {40.33, 72.39967},
	}
	fmt.Println(map2)

	fmt.Println(":Function Values:")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("hypot:", hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println(":Function Closures:")
	// Executing 2 different adders to verify the closure behavior
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}

func adder() func(int) int { // func(int) int, is the return type of the func adder()
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Vertex1 Lat, Lang
type Vertex1 struct {
	Lat, Long float64
}

var map1 map[string]Vertex1

// User defined functions
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Vertex type
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p2 = &Vertex{1, 2} // has type *Vertex
)
