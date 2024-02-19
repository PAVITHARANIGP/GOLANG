package main

import (
	"fmt"
)

type Book struct {
	Title     string
	Available bool
}

func main() {
	// Initialize inventory with some books
	inventory := map[string]Book{
		"book1":                             {Title: "The Lord of the Rings", Available: true},
		"TheHitchhiker'sGuidetotheGalaxy": {Title: "The Hitchhiker's Guide to the Galaxy", Available: false},
		"PrideandPrejudice":                  {Title: "Pride and Prejudice", Available: true},
	}

	// Get book title from user
	var bookTitle string
	fmt.Println("Enter book title: ")
	fmt.Scanln(&bookTitle)

	// Check availability
	book, ok := inventory[bookTitle]
	if ok {
		if book.Available {
			fmt.Println(bookTitle, "is available!")
		} else {
			fmt.Println(bookTitle, "is currently unavailable.")
		}
	} else {
		fmt.Println("Book not found in inventory.")
	}
}
