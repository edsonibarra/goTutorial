package main

import "fmt"


func main() {
	nums := []int{1,2,3,4,5,6,7,8}
	fmt.Println(nums)
	fmt.Println("Slicing...")
	fmt.Println(nums[0:3])

	favoriteNums := []int{2,4,8,16,32}
	fmt.Println(favoriteNums)

	fmt.Println("slicing...")
	fmt.Println(favoriteNums[1:])

	fmt.Println("\nLiked Programming Languages")
	likedProgrammingLanguages := []string{"Go","Python","Java","Scala","C"}
	fmt.Println(likedProgrammingLanguages)

	twoProgLang := likedProgrammingLanguages[:2]
	fmt.Println("Two Programming Languages slice: [:2]")
	fmt.Println(twoProgLang)

	restProLang := likedProgrammingLanguages[2:]
	fmt.Println("The Rest of Programming Languages slice: [2:]")
	fmt.Println(restProLang)

	restProLang[1] = "Rust"
	fmt.Println("Modified restProLang. Now =", restProLang)
	fmt.Println("Whole array was mutated too",likedProgrammingLanguages) // The original array is modified
}