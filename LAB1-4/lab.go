package main

import (
    "fmt"
)

// Item represents a photography prop item in the store.
type Item struct {
    ID       int
    Name     string
    Quantity int
    Price    float64
    
}

// Store represents a photography prop store.
type Store struct {
    inventory map[int]Item // inventory maps item ID to Item
}

// NewStore creates and returns a new instance of Store.
func NewStore() *Store {
    return &Store{
        inventory: make(map[int]Item),
    }
}

// AddItem adds a new item to the store's inventory.
func (s *Store) AddItem(id int, name string, quantity int, price float64) {
    newItem := Item{
        ID:       id,
        Name:     name,
        Quantity: quantity,
        Price:    price,
    }
    s.inventory[id] = newItem
}

// SellItem attempts to sell an item from the store's inventory.
// It returns true if the item was successfully sold, and false otherwise.
func (s *Store) SellItem(id int, quantity int) bool {
    item, exists := s.inventory[id]
    if !exists {
        fmt.Println("Item not found in inventory")
        return false
    }

    if item.Quantity < quantity {
        fmt.Println("Insufficient quantity in inventory")
        return false
    }

    // Update quantity
    item.Quantity -= quantity
    s.inventory[id] = item
    fmt.Printf("Sold %d units of %s\n", quantity, item.Name)
    return true
}

func main() {
    // Create a new store
    store := NewStore()

    // Prompt user to add items
    fmt.Println("Add items to inventory:")
    for {
        var id, quantity int
        var name string
        var price float64

        fmt.Print("Enter item ID (0 to stop): ")
        fmt.Scanln(&id)
        if id == 0 {
            break
        }

        fmt.Print("Enter item name: ")
        fmt.Scanln(&name)

        fmt.Print("Enter item quantity: ")
        fmt.Scanln(&quantity)

        fmt.Print("Enter item price: ")
        fmt.Scanln(&price)

        store.AddItem(id, name, quantity, price)

        // Ask if user wants to add more items
        var addMore string
        fmt.Print("Do you want to add more items? (yes/no): ")
        fmt.Scanln(&addMore)
        if addMore != "yes" {
            break
        }
    }

    // Display initial inventory
    fmt.Println("\nInitial Inventory:")
    for _, item := range store.inventory {
        fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: $%.2f\n", item.ID, item.Name, item.Quantity, item.Price)
    }

    // Prompt user to sell items
    fmt.Println("\nSell items from inventory:")
    for {
        var id, quantity int

        fmt.Print("Enter item ID to sell (0 to stop): ")
        fmt.Scanln(&id)
        if id == 0 {
            break
        }

        fmt.Print("Enter quantity to sell: ")
        fmt.Scanln(&quantity)

        store.SellItem(id, quantity)
    }

    // Display updated inventory
    fmt.Println("\nUpdated Inventory:")
    for _, item := range store.inventory {
        fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: $%.2f\n", item.ID, item.Name, item.Quantity, item.Price)
    }
}

