// --Summary:
//
//	Create a program that can activate and deactivate security tags
//	on products.
//
// --Requirements:
package main

import "fmt"

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
type Item struct {
	name             string
	securityTagState bool
}

// * Create functions to activate and deactivate security tags using pointers
func activate(item *Item) {
	item.securityTagState = true
}
func deactivate(item *Item) {
	item.securityTagState = false
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	for i, _ := range items {
		deactivate(&items[i])
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	items := []Item{
		{"Shoes", true},
		{"Shirt", true},
		{"Trousers", true},
		{"Sweater", true},
	}
	fmt.Println("Initial", items)
	//  - Deactivate any one security tag in the array/slice
	deactivate(&items[3])
	fmt.Println("Initial", items)
	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	fmt.Println("Initial", items)
}
