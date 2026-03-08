package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// swapping of string
func swap(x, y string) (string, string) {
	return y, x
}

func swapping(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// variables with initializers
var c, j int = 1, 2

func main() {
	fmt.Println("Hello World")
	fmt.Println("Starting to learn go as a backend")
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	fmt.Println("Sum of 1 and 3 is : ", add(1, 3))

	a, b := swap("Hello World", "Go Programming")
	println(a, b)

	p, q := swapping("hello ", " shubham")
	fmt.Println(p, q)

	fmt.Println(split(17))

	// variables with initializers
	var i, python, java = true, false, "no!"
	fmt.Println(j, c, i, python, java)

	// short variable declarations
	var e, s int = 5, 4
	k := 3
	c, py, javascript := true, false, "yes"
	fmt.Println(e, s, k, c, py, javascript)

}
