package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index >= len(s.values) {
		return 0, fmt.Errorf("no element at index %v", index)
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{[]int{3, 2, 1}}
	value, err := stuff.Get(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}
