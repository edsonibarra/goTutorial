package main

import "fmt"

func main() {
	totalSum := migrateFromCToGo()
	fmt.Println(totalSum)
}

func migrateFromCToGo() int {
	
	nums := []int{6, 2, 77, 4, 12}
	
	count := 0
	totalSum := 0

	for count < len(nums) {
		totalSum += nums[count]
		count++
	}
	
	return totalSum
}
