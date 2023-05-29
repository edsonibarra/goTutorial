package main

import "fmt"

func main() {
	cityPopulation := map[string]int{
		"New York": 8468000,
		"Tokio": 13960000,
	}
	fmt.Println(cityPopulation)
	for key, value := range cityPopulation {
		fmt.Printf("City: %s, population: %d\n", key, value)
	}
}