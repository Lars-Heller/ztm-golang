//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	a, b int
}

func areaAndPerimeter(r Rectangle) (int, int) {
	return r.a * r.b, 2 * (r.a + r.b)
}

func main() {

	r := Rectangle{2, 2}
	area, perimeter := areaAndPerimeter(r)
	fmt.Println("Rectange", r, "has area", area, "and perimeter", perimeter)

	r = Rectangle{1, 4}
	area, perimeter = areaAndPerimeter(r)
	fmt.Println("Rectange", r, "has area", area, "and perimeter", perimeter)

}
