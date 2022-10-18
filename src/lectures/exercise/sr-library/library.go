//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Member string
type Title string
type CheckoutInfo struct {
	lentBy   Member
	checkout time.Time
	checkin  time.Time
}

type Library struct {
	members []Member
	books   map[Title]CheckoutInfo
}

func addMember(library *Library, member Member) {
	library.members = append(library.members, member)
}
func addBook(library *Library, bookTitle Title) {
	library.books[bookTitle] = CheckoutInfo{}
}

func printStats(library *Library) {
	fmt.Println("--- Current Library state ---")
	fmt.Println("Members:", library.members)
	fmt.Println("Books:")
	for title, info := range library.books {
		fmt.Print("  ", title)
		if info.lentBy == "" {
			fmt.Print(" is available")
			if !info.checkin.IsZero() {
				fmt.Print(" since", info.checkin)
			}
		} else {
			fmt.Print(" is lent by ", info.lentBy, " since ", info.checkout)
		}
		fmt.Println()
	}
}

func isMember(library *Library, checkMember Member) bool {
	for _, member := range library.members {
		if member == checkMember {
			return true
		}
	}
	return false
}

func checkout(library *Library, bookTitle Title, member Member) {
	checkoutInfo, found := library.books[bookTitle]
	if !found {
		fmt.Println("Cannot checkout", bookTitle, "since it is not in the library")
		return
	}
	if checkoutInfo.lentBy != "" {
		fmt.Println("Cannot checkout", bookTitle, "since it is lent by", checkoutInfo.lentBy)
		return
	}
	if !isMember(library, member) {
		fmt.Println("Cannot checkout", bookTitle, "since", member, "is not a member")
		return
	}

	library.books[bookTitle] = CheckoutInfo{lentBy: member, checkout: time.Now()}
	fmt.Println(member, "checked out", bookTitle)
}

func checkin(library *Library, bookTitle Title) {
	checkoutInfo, found := library.books[bookTitle]
	if !found {
		fmt.Println("Cannot checkin", bookTitle, "since it is not in the library")
		return
	}

	library.books[bookTitle] = CheckoutInfo{lentBy: "", checkin: time.Now()}
	fmt.Println(checkoutInfo.lentBy, "checked in", bookTitle)
}

func main() {
	library := Library{books: make(map[Title]CheckoutInfo)}

	addBook(&library, "Pragmatic_Progammer")
	addBook(&library, "Neverending_Story")
	addBook(&library, "DDD")
	addBook(&library, "MTB_Skills")

	addMember(&library, "Paul")
	addMember(&library, "Emil")
	addMember(&library, "Tom")

	printStats(&library)

	// checkout non existing book
	checkout(&library, "The Bible", "Paul")

	// checkout by non-member
	checkout(&library, "DDD", "Frank")

	// checkout
	checkout(&library, "DDD", "Paul")
	printStats(&library)

	// checkout already checked out book
	checkout(&library, "DDD", "Emil")

	// checkin again
	checkin(&library, "DDD")
	printStats(&library)
}
