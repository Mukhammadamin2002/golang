package main

import "fmt"

func main() {
	fmt.Println("Enter the number: ")
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "is even number")
	} else {
		fmt.Println(n, "is odd number")
	}
}
