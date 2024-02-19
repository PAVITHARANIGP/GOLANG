package main

import (
	"fmt"
)

func main() {
	var mathScore, scienceScore, englishScore float64

	fmt.Println("Welcome to the Grade Calculator Program!")
	fmt.Println("Please enter your scores for Math, Science, and English:")

	fmt.Print("Math Score: ")
	fmt.Scanln(&mathScore)

	fmt.Print("Science Score: ")
	fmt.Scanln(&scienceScore)

	fmt.Print("English Score: ")
	fmt.Scanln(&englishScore)

	// Calculate the average score
	averageScore := (mathScore + scienceScore + englishScore) / 3

	// Determine the corresponding letter grade based on the average score
	var letterGrade string
	switch {
	case averageScore >= 90:
		letterGrade = "A"
	case averageScore >= 80:
		letterGrade = "B"
	case averageScore >= 70:
		letterGrade = "C"
	case averageScore >= 60:
		letterGrade = "D"
	default:
		letterGrade = "F"
	}

	fmt.Printf("Your average grade is %.2f, which corresponds to the grade %s.\n", averageScore, letterGrade)
}
