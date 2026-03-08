package main

import (
	"fmt"
	"math/cmplx"
)

func main() { // Function ka naam Capital se shuru karo (Exportable)

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
		age    int        = 25
	)

	// as a type of formatted print statements in go 
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Age is: %T Value : %v\n", age, age)
	fmt.Printf("Age is: %v" , age )

	

}
