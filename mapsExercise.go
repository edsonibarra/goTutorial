package main

import (
	"fmt"
	"strings"
)


//My solution
func WordCount(s string) map[string]int {
	// write your solution here
   words := strings.Fields(s)
   counter := make(map[string]int)
   fmt.Printf("%q\n", words)
   for _, word := range words{
	  _, ok := counter[word]
	  if ok {
		 counter[word] = counter[word] + 1
	  }else{
		 counter[word] = 1
	  }
   }
   return counter //returning an empty map, modify the statement to return the correct answer
 }

 //educative solution

 func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, word := range words {
		count[word]++
	}
	return count
 }
 