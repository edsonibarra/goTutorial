package main

import "fmt"

func main() {
	
	basicFor(5)
	getSumInRange(3)
	forLoopWithoutStatements()
	forInsteadWhile()

}

func basicFor(n int) {
	//basic for loop example
	fmt.Println("\n----- Basic For -----\n")
	for i:=0; i<n; i++{
		fmt.Println(i)
	}
}

func getSumInRange(n int) {
	// get the sum of consecutive numbers up to n
	fmt.Println("\n----- get sum in range -----\n")

	sum := 0
	for i:=0; i<n; i++{
		sum += i
	}
	fmt.Println(sum)
}

func forLoopWithoutStatements () {
	fmt.Println("\n----- For loop without statement -----\n")

	sum := 1
	
	for ; sum < 1000 ; {
		sum += sum
	}
	
	fmt.Println("Current Sum:", sum)
}

func forInsteadWhile() int {
	fmt.Println("\n----- for instead of while -----\n")

	// For loop as an alternative to while loop

	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("For loop as While loop sum =", sum)
	return sum
}