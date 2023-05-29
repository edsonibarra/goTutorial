package main

import "fmt"

func main() {
	var users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distributeBitcoins(users)
}

func distributeBitcoins(users []string)  {
	const MAX int = 10
	result := make(map[string]int)
	budget := 50
	for _, user := range users {
		v := valueOfWord(user)
		fmt.Println("The value of",user,"is",v)
		if v > MAX {
			v = 10
		}
		budget -= v
		result[user] = v
	}
	fmt.Println("Result:\n",result,"\nBudget:\n", budget)
}


func valueOfWord (word string)int{
	value := map[rune]int{
		'A':1,
		'a':1,
		'E':1,
		'e':1,
		'I':2,
		'i':2,
		'O':3,
		'o':3,
		'U':4,
		'u':4,
	}
	totalValue := 0
	for _, c := range word {
		_, ok := value[c]
		if ok {
			totalValue += value[c]
		}
	}
	return totalValue

}