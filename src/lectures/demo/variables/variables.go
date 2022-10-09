package main

import "fmt"

func main() {
	var myName = "rudolfson"
	fmt.Println("my name is", myName, myName)

	var name string = "Kathy"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("sum is", sum)

	part1, other := 1, 5

	fmt.Println("part1 =", part1, ", other =", other)

	part2, other := 2, 0
	fmt.Println("part2 =", part2, ", other =", other)

	sum = part1 + part2
	fmt.Println("sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName and type is", lessonName, lessonType)

	word1, word2, _ := "one", "two", "three"
	fmt.Println("word1 is", word1, "word two is", word2)
}
