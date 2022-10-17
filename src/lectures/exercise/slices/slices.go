//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(title string, line []Part) {
	fmt.Println("\n---", title, "---")
	for i := 0; i < len(line); i++ {
		item := line[i]
		fmt.Println(item)
	}
}

func main() {
	// * Using a slice, create an assembly line that contains type Part
	assemblyLine := []Part{}
	// * Create a function to print out the contents of the assembly line
	printAssemblyLine("Start", assemblyLine)
	// * Perform the following:
	//   - Create an assembly line having any three parts
	assemblyLine = append(assemblyLine, "Part 1", "Part 2", "Part 3")
	printAssemblyLine("Add 3 parts", assemblyLine)
	//   - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Part 4", "Part 5")
	printAssemblyLine("Add 2 parts", assemblyLine)
	//   - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	printAssemblyLine("Remove first 3 parts", assemblyLine)
	//   - Print out the contents of the assembly line at each step
}
