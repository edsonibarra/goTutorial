package main

import "fmt"

func main() {
	grid := [2][2]int{{1,2}, {3,4}}
	fmt.Println(grid)

	var nums [2][2]int // initialize the array -> only zeroes
	fmt.Println(nums)

	printEachNum(grid)
}

func printEachNum(arr [2][2]int) {
	for i:=0; i<len(arr); i++{
		fmt.Println(arr[i])
		for j:=0; j<len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
}