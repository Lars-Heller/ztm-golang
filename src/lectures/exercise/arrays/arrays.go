//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func printStats(list [4]Product) {
	var cost, totalItems int
	var lastItem Product

	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.name != "" {
			cost += item.price
			totalItems++
			lastItem = item
		}
	}
	fmt.Println("Last item on the list is", lastItem)
	fmt.Println("cost is", cost)
	fmt.Println(totalItems, "on the list")
}

func main() {
	//   - Using an array, create a shopping list with enough room
	//     for 4 products
	var shoppingList [4]Product
	//   - Products must include the price and the name
	//
	// * Insert 3 products into the array
	shoppingList[0] = Product{"Klopapier", 289}
	shoppingList[1] = Product{"Seife", 700}
	shoppingList[2] = Product{"Äpfel", 349}
	//* Print to the terminal:
	printStats(shoppingList)
	//  - The last item on the list
	//  - The total number of items
	//  - The total cost of the items
	//* Add a fourth product to the list and print out the
	//  information again
	shoppingList[3] = Product{"Essig", 379}
	printStats(shoppingList)
}
