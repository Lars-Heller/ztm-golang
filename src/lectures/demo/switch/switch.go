package main

import "fmt"

func price() int {
	return 21
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("cheap")
	case p < 10:
		fmt.Println("moderate")
	default:
		fmt.Println("expensive")
	}

	ticket := FirstClass
	switch ticket {
	case Economy:
		fmt.Println("economy")
	case Business:
		fmt.Println("business")
	case FirstClass:
		fmt.Println("first class")
	default:
		fmt.Println("other")
	}

}
