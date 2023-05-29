package main

import "fmt"

func main() {
	binaryPows := []int{1,2,4,8,16}
	for i, v := range binaryPows {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}