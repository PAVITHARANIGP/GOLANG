package main

import (
	"fmt"
	"math"
)

func addition(x, y float64) float64 {
	return x + y
}

func subtraction(x, y float64) float64 {
	return x - y
}

func multiplication(x, y float64) float64 {
	return x * y
}

func division(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return x / y
}

func exponentiation(x, y float64) float64 {
	return math.Pow(x, y)
}

func squareRoot(x float64) float64 {
	if x < 0 {
		fmt.Println("Error: Cannot calculate square root of a negative number")
		return 0
	}
	return math.Sqrt(x)
}

func main() {
	var choice int
	var x, y float64

	fmt.Println("Scientific Calculator")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Exponentiation (x^y)")
	fmt.Println("6. Square root (√x)")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter two numbers separated by space: ")
		fmt.Scanln(&x, &y)
		fmt.Printf("Result: %.2f\n", addition(x, y))
	case 2:
		fmt.Print("Enter two numbers separated by space: ")
		fmt.Scanln(&x, &y)
		fmt.Printf("Result: %.2f\n", subtraction(x, y))
	case 3:
		fmt.Print("Enter two numbers separated by space: ")
		fmt.Scanln(&x, &y)
		fmt.Printf("Result: %.2f\n", multiplication(x, y))
	case 4:
		fmt.Print("Enter two numbers separated by space: ")
		fmt.Scanln(&x, &y)
		fmt.Printf("Result: %.2f\n", division(x, y))
	case 5:
		fmt.Print("Enter base and exponent separated by space: ")
		fmt.Scanln(&x, &y)
		fmt.Printf("Result: %.2f\n", exponentiation(x, y))
	case 6:
		fmt.Print("Enter a number: ")
		fmt.Scanln(&x)
		fmt.Printf("Result: %.2f\n", squareRoot(x))
	default:
		fmt.Println("Invalid choice")
	}
}
