package main

import "fmt"

func main() {
	for i := 1; i < 50; i++ {
		st := ""
		if i%3 == 0 {
			st += "Fizz"
		}
		if i%5 == 0 {
			st += "Buzz"
		}
		if i%15 == 0 {
			st += "FizzBuzz"
		}
		if st != "" {
			fmt.Println(st)
		} else {
			fmt.Println(i)
		}

	}
}
