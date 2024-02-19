package main

import "fmt"

//creating struct-for-product
type product struct {
	id    int
	name  string
	brand string
	price float32
}

func main() {
	//Defining-
	var products []product
	var opt bool = true
	var id int = 100
	var name, s string
	var brand string
	var price float32
	fmt.Println("\t\tProduct Inventory")
	for opt {
		fmt.Print("\nEnter the product name : ")
		fmt.Scan(&name)
		fmt.Print("\nEnter the product brand : ")
		fmt.Scan(&brand)
		fmt.Print("\nEnter the product price : ")
		fmt.Scan(&price)
		products = append(products, addProduct(id, name, brand, price))
		id++
		fmt.Print("\nDo you want to add more?(y/n): ")
		fmt.Scan(&s)
		if s != "y" {
			opt = false
		}
	}
	display(products)
}

func addProduct(id int, name string, brand string, price float32) product {
	var item product
	item.id = id
	item.name = name
	item.brand = brand
	item.price = price
	return item
}
func display(products []product) {
	for enum, item := range products { //for _, item := range
		fmt.Printf("\n========Product-%d==============\n", enum+1)
		fmt.Println("\t\tID :", item.id)
		fmt.Println("\t\tName :", item.name)
		fmt.Println("\t\tBrand :", item.brand)
		fmt.Printf("\t\tPrice : %.2f", item.price)
	}
}