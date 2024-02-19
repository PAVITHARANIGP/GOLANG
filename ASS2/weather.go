package main

import (
	"fmt"
)

func main() {
	var temperature float64
	var humidity float64
	var windSpeed float64
	var addMore string

	for {
		fmt.Println("Please enter the temperature:")
		fmt.Scanln(&temperature)

		fmt.Println("Please enter the humidity:")
		fmt.Scanln(&humidity)

		fmt.Println("Please enter the wind speed:")
		fmt.Scanln(&windSpeed)

		fmt.Println("Weather Report:")
		fmt.Println("Temperature:", temperature, "°C")
		fmt.Println("Humidity:", humidity, "%")
		fmt.Println("Wind Speed:", windSpeed, "m/s")

		// Using switch case to provide a weather description based on temperature
		switch {
		case temperature < 0:
			fmt.Println("It's freezing!")
		case temperature >= 0 && temperature < 15:
			fmt.Println("It's cold.")
		case temperature >= 15 && temperature < 25:
			fmt.Println("It's moderate.")
		case temperature >= 25:
			fmt.Println("It's hot!")
		}

		fmt.Println("Do you want to add more weather data? (yes/no)")
		fmt.Scanln(&addMore)
		if addMore != "yes" {
			break
		}
	}
}
