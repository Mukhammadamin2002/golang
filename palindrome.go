package main

import (
	"fmt"
)

func plm(word string) string {

	for i := 0; i < len(word)/2; i++ {

		if word[i] != word[len(word)-1-i] {
			return "This is not"
		}

	}

	return "This word is Palindrome"
}

func main() {
	result := plm("anna")
	fmt.Println(result)
}
