package main

import "fmt"


type Point struct {
	X, Y int
}

func main() {
	// this is a way to declare a variable that will hold an address
	var p2 *Point
	
	// this is another way to create a variable that will hold an address using the short assignment operator
	p0 := &Point{}

	p1 := new(Point) // Initializing variables using new
	
	p2 = &Point{} // empty struct literal 
	
	fmt.Println(*p1 == *p2)
}	