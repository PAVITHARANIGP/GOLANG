// Imagine you are developing a simple weather application in Go that takes user input for the current
//temperature in Celsius and provides a weather recommendation based on the following conditions:

// If the temperature is below 10 degrees Celsius, recommend wearing a heavy jacket.
// If the temperature is between 10 and 20 degrees Celsius (inclusive), recommend wearing a light jacket.
// If the temperature is above 20 degrees Celsius, recommend wearing a t-shirt.
// Write a Go program that takes the user input for the current temperature,
//processes it using variables and control flow structures, and prints the appropriate weather recommendation.

// Your program should include the following:

// Declaration of a variable to store the temperature.
// Input statement to get the temperature from the user.
// Conditional statements to determine the appropriate weather recommendation based on the temperature.
// Output statement to display the weather recommendation.

package main

import "fmt"

func main() {
	var temp float32
	fmt.Print("Enter the current temperature : ")
	fmt.Scan(&temp)
	if temp > 20 {
		fmt.Println("Temperature is above 20 degrees Celsius, recommend wearing a t-shirt")
	} else if temp > 10 {
		fmt.Println("Wear a light jacket")
	} else {
		fmt.Println("It's too cold.Wear a heavy jacket")
	}
}
