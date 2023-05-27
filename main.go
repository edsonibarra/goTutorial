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
	var maxInt8 int8 = 127
	var minInt8 int8 = -128
	var maxInt16 int16 = 32767
	var minInt16 int16 = -32768
	var maxInt32 int32 = 2147483647
	var minInt32 int32 = -2147483648
	var maxInt64 int64 = 9223372036854775807
	var minInt64 int64 = -9223372036854775808
	var pi float64 = 3.14 // 32-bit floating-point numbers

	fmt.Println("bool       =", isEmployee)
	fmt.Println("string     =", employeeName)
	fmt.Println("uint32 max =", maxUint32)
	fmt.Println("unit32 min =", minUint32)
	fmt.Println("uint64 max =", maxUint64)
	fmt.Println("uint64 min =", minUint64)
	fmt.Println("int        =", year)
	fmt.Println("int8 max   =", maxInt8)
	fmt.Println("int8 min   =", minInt8)
	fmt.Println("int16 max  =", maxInt16)
	fmt.Println("int16 min  =", minInt16)
	fmt.Println("int32 max  =", maxInt32)
	fmt.Println("int32 min  =", minInt32)
	fmt.Println("int64 max  =", maxInt64)
	fmt.Println("int64 min  =", minInt64)
	fmt.Println("float64    =", pi)
}