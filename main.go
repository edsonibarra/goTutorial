package main

import "fmt"

func main() {
	// Basics Types
	var isEmployee bool = true // True or False
	var employeeName string = "John Doe" // An array of characters
	var maxUint32 uint32 = 4294967295 // uint -> (unsigned integer) data type holds only non-negative values
	var minUint32 uint32 = 0
	var maxUint64 uint64 = 18446744073709551615
	var minUint64 uint64 = 0
	var year int = 2023 // holds 32 or 64 bit values that can be negative
	var pi float64 = 3.14 // 32-bit floating-point numbers

	fmt.Println("bool =", isEmployee)
	fmt.Println("string =", employeeName)
	fmt.Println("uint32 max =", maxUint32)
	fmt.Println("unit32 min =", minUint32)
	fmt.Println("uint64 max =", maxUint64)
	fmt.Println("uint64 min =", minUint64)
	fmt.Println("int =", year)
	fmt.Println("float64 =", pi)
	

}