package main

import "fmt"

type Task struct {
	day      string
	title    string
	duration float32
	hour     string
}

func main() {
	d1 := Task{"Monday", "Housework", 2.00, "hour"}
	d2 := Task{"Tuesday", "Laisure Activity", 1.30, "hour"}
	d3 := Task{"Wednesday", "Go outside City", 4.00, "hour"}
	d4 := Task{"Thursday", "Climbing", 2.00, "hour"}
	d5 := Task{"Friday", "Fishing", 2.00, "hour"}
	d6 := Task{"Saturday", "Free", 3.00, "hour"}
	d7 := Task{"Sunday", "Holiday", 7.00, "hour"}

	fmt.Println("\n", d1, "\n", d2, "\n", d3, "\n", d4, "\n", d5, "\n", d6, "\n", d7)

}
