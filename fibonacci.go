package main

import "fmt"

func fibGen() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}

}

func main() {
	fib := fibGen()
	for i := 0; i < 10; i++ {
		fmt.Println(fib(), "")
	}
}