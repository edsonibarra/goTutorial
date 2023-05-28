package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Point struct {
	X, Y int
}

func CreatePerson(name string, age int) Person {
	var p1 Person
	p1 = Person{Name: name, Age: age}
	return p1
}

func main() {
	var people [3]Person // Initialize an array of length 3 type Person
	
	// Create Person
	var p1 Person = CreatePerson("John", 43)
	var p2 Person = CreatePerson("Jane", 46)
	var p3 Person = CreatePerson("Will", 50)
	
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	// Assign each array index to a person object
	people[0] = p1
	people[1] = p2
	people[2] = p3

	// Print the whole array
	fmt.Println(people)

	// Create 3 variables of type Point 
	var point1 Point = Point{X:10, Y:12}
	var point2 Point = Point{X:11, Y:13}
	var point3 Point = Point{X:12, Y:14}

	points := [3]Point{point1, point2, point3}

	fmt.Println(points)

	// elipsis
	bestNumbers := [...]int{2,4,8,16} // Length of the array is determined automatically in the initialization
	fmt.Println(bestNumbers)
}