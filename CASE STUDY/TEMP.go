package main

import (
	"fmt"
)

func main() {
	var choice int
	var temperature float64

	fmt.Println("Temperature Converter")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("3. Celsius to Kelvin")
	fmt.Println("4. Kelvin to Celsius")
	fmt.Println("5. Fahrenheit to Kelvin")
	fmt.Println("6. Kelvin to Fahrenheit")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	fmt.Print("Enter temperature: ")
	fmt.Scanln(&temperature)

	switch choice {
	case 1:
		fmt.Printf("%.2f°C is %.2f°F\n", temperature, celsiusToFahrenheit(temperature))
	case 2:
		fmt.Printf("%.2f°F is %.2f°C\n", temperature, fahrenheitToCelsius(temperature))
	case 3:
		fmt.Printf("%.2f°C is %.2fK\n", temperature, celsiusToKelvin(temperature))
	case 4:
		fmt.Printf("%.2fK is %.2f°C\n", temperature, kelvinToCelsius(temperature))
	case 5:
		fmt.Printf("%.2f°F is %.2fK\n", temperature, fahrenheitToKelvin(temperature))
	case 6:
		fmt.Printf("%.2fK is %.2f°F\n", temperature, kelvinToFahrenheit(temperature))
	default:
		fmt.Println("Invalid choice!")
	}
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9/5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func fahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5/9 + 273.15
}

func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin - 273.15) * 9/5 + 32
}
