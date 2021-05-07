package main

import "fmt"

type Contact struct {
	name      string
	last_name string
	phone_num int
}

func main() {
	c1 := Contact{"John", "Doe", 9989456789}
	c2 := Contact{"Alex", "Smith", 9989456789}
	c3 := Contact{"Jo", "Jaary", 9989456789}
	c4 := Contact{"Taylor", "Srith", 9989456789}
	fmt.Println("\n", c1, "\n", c2, "\n", c3, "\n", c4)
}
