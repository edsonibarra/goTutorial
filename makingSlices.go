package main

import "fmt"

func main() {
	fmt.Println("--- Create slices using make keyword ---")
	// It works by allocating a zeroed array and returning a slice that refers to that array.
	names := make([]string, 3)
	fmt.Printf("%q",names)

	nums := make([]int, 3)
	fmt.Println(nums)

	fmt.Println("--- Create slices using make keyword ---")
	
	
	cities := []string{}
	fmt.Printf("Initial Array:\n%q\n",cities)
	
	city := "San Diego"
	fmt.Println("Appending a new city:", city)
	cities = append(cities, city)
	fmt.Printf("Cities:\n%q",cities)
	
	city = "Los Angeles"
	fmt.Println("Appending a new city:", city)
	cities = append(cities, city)
	fmt.Printf("Cities:\n%q",cities)

	otherCities := []string{"New York", "Seatlle", "San Francisco", "Austin"}
	fmt.Printf("\nPrinting otherCities slice:\n%q\n", otherCities)
	fmt.Println("\nCombining both slices cities + otherCities\nResult")
	cities = append(cities, otherCities...)
	fmt.Printf("%q\n",cities)
}