package usefulinformations

import "fmt"

// Add is a funcition type
type Add func(a int, b int) int

// FunctionType is
func FunctionType() {
	var result int
	result = process(func(a int, b int) int {
		return a + b
	})

	fmt.Println(result)
}

func process(adder Add) int {
	return adder(1, 2)
}
