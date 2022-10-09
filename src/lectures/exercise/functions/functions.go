//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hello", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func echo(message string) string {
	return message
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func sum(a, b, c int) int {
	return a + b + c
}

// - Write a function that returns any number
func echoNumber(number int) int {
	return number
}

// - Write a function that returns any two numbers
func echoNumbers(a, b int) (int, int) {
	return a, b
}

//   - Add three numbers together using any combination of the existing functions.
//   - Print the result
//   - Call every function at least once

func main() {
	greet("Frank")

	fmt.Println("Echoing a message", echo("a message"))

	var a, b = echoNumbers(1, 2)
	var sum = sum(a, b, echoNumber(3))
	fmt.Println("Sum is", sum)

}
