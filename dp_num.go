package main

import "fmt"

func main() {
	var a []int = []int{1, 4, 5, 6, 7, 2, 1, 4}

	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && j > i {
				fmt.Println("Duplicated number", element)
			}
		}
	}
}
