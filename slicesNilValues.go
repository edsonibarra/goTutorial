package main

import "fmt"

func main() {
	var cities []string
	fmt.Printf("Cities: %q\n", cities)
	fmt.Println("The length of cities is:", len(cities))
	fmt.Println("The capacity of cities is:",cap(cities))
	if cities == nil {
		fmt.Println("Cities slice is nil")
	}
}