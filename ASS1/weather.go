package main

import (
	"fmt"
)

func main() {
	for {
		// Declare a variable to store the temperature
		var temperature float64

		// Input statement to get the temperature from the user
		fmt.Println("Enter the current temperature in Celsius:")
		fmt.Scanln(&temperature)

		// Conditional statements to determine the appropriate weather recommendation
		var recommendation string

		if temperature < 10 {
			recommendation = "Wear a heavy jacket."
		} else if temperature >= 10 && temperature <= 20 {
			recommendation = "Wear a light jacket."
		} else {
			recommendation = "Wear a t-shirt."
		}

		// Output statement to display the weather recommendation
		fmt.Println("Weather recommendation:", recommendation)

		// Ask the user if they want to continue
		var continueInput string
		fmt.Println("Do you want to continue? (yes/no)")
		fmt.Scanln(&continueInput)
		if continueInput != "yes" {
			break
		}
	}
}
