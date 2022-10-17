package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk removed, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	_, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("we need to buy", totalItems, "items")
}
