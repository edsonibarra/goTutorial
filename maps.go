package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {

	// Create a map using map literals
	celebs := map[string]int {
		"Nicolas Cage": 50,
		"Selena Gomez": 21,
		"Scarlett Johansson": 29,
	}
	fmt.Printf("%#v\n", celebs)

	// Create a map using make keyword
	myMap := make(map[string]Vertex)
	myMap["Bell Labs"] = Vertex{40.68433, -74.39967}
	
	fmt.Println(myMap)
	fmt.Println(myMap["Bell Labs"])

	myMap["Apple Computers"] = Vertex{12.32, -32.43}
	
	fmt.Println(myMap)
	fmt.Println(myMap["Apple Computers"])

	myMap["Microsoft"] = Vertex{43.3222, -43.3222}
	
	fmt.Println(myMap)
	fmt.Println(myMap["Microsoft"])

	fmt.Println("Deleting Microsoft from myMap")
	delete(myMap, "Microsoft")
	fmt.Println(myMap)

	//Check if a key is present in the dictonary

	name, ok := myMap["Microsoft"]
	fmt.Println(name, ok)

	value, exists := myMap["Apple Computers"]
	fmt.Println(value, exists)
}