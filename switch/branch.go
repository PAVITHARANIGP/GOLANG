package main

import (
	"fmt"
)

// Prop struct represents information about a photography prop
type Prop struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Rental struct represents information about a prop rental
type Rental struct {
	PropID int
	Days   int
}

func main() {
	// Creating a map to store prop inventory
	propInventory := make(map[int]Prop)
	// Adding sample props to inventory
	propInventory[1] = Prop{ID: 1, Name: "Vintage Camera", Quantity: 5, Price: 50.0}
	propInventory[2] = Prop{ID: 2, Name: "Film Reel", Quantity: 10, Price: 20.0}
	propInventory[3] = Prop{ID: 3, Name: "Studio Lighting Kit", Quantity: 3, Price: 150.0}

	// Creating slices to manage prop rentals
	rentals := make(map[int][]Rental)

	for {
		fmt.Println("\nPhotography Prop Rental System")
		fmt.Println("1. View Props")
		fmt.Println("2. Rent Prop")
		fmt.Println("3. View Rentals")
		fmt.Println("4. Return Prop")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nAvailable Props:")
			for _, prop := range propInventory {
				fmt.Printf("Prop ID: %d, Name: %s, Quantity: %d, Price: $%.2f\n", prop.ID, prop.Name, prop.Quantity, prop.Price)
			}
		case 2:
			var propID, days int
			fmt.Print("\nEnter Prop ID to rent: ")
			fmt.Scanln(&propID)
			fmt.Print("Enter number of days to rent: ")
			fmt.Scanln(&days)

			if prop, ok := propInventory[propID]; ok {
				if prop.Quantity > 0 {
					rentals[propID] = append(rentals[propID], Rental{PropID: propID, Days: days})
					fmt.Println("Prop rented successfully!")
					// propInventory[propID].Quantity-- // Decrease available quantity
				} else {
					fmt.Println("Prop out of stock.")
				}
			} else {
				fmt.Println("Prop not found in inventory.")
			}
		case 3:
			fmt.Println("\nRented Props:")
			for propID, propRentals := range rentals {
				for _, rental := range propRentals {
					fmt.Println("Prop ID: ", propID)
					fmt.Printf("Rental Days: %d\n", rental.Days)
				}
			}
		case 4:
			var propID int
			fmt.Print("\nEnter Prop ID to return: ")
			fmt.Scanln(&propID)

			if prop, ok := propInventory[propID]; ok {
				if len(rentals[propID]) > 0 {
					// Remove the latest rental
					rentals[propID] = rentals[propID][:len(rentals[propID])-1]
					fmt.Println("Prop returned successfully!")
					// propInventory[propID].Quantity++ // Increase available quantity
				} else {
					fmt.Println("No rentals found for this prop.")
				}
			} else {
				fmt.Println("Prop not found in inventory.")
			}
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
