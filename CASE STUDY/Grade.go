package main

import (
	"fmt"
)

func main() {
	// Define grading criteria
	gradingMap := map[int]string{
		90: "A",
		80: "B",
		70: "C",
		60: "D",
		0:  "F",
	}

	// Initialize variables
	var numSubjects int
	var studentName string
	var subjectScores map[string]int

	// Get student name
	fmt.Println("Enter student name: ")
	fmt.Scanln(&studentName)

	// Get number of subjects
	fmt.Println("Enter number of subjects: ")
	fmt.Scanln(&numSubjects)

	// Create map to store subject scores
	subjectScores = make(map[string]int)

	// Get score for each subject
	for i := 1; i <= numSubjects; i++ {
		var subjectName string
		var score int

		fmt.Printf("Enter subject %d name: ", i)
		fmt.Scanln(&subjectName)

		fmt.Printf("Enter score for %s: ", subjectName)
		fmt.Scanln(&score)

		subjectScores[subjectName] = score
	}

	// Calculate total score and average
	var totalScore int
	for _, score := range subjectScores {
		totalScore += score
	}
	averageScore := float64(totalScore) / float64(numSubjects)

	// Determine grade based on grading criteria
	var grade string
	for threshold, letterGrade := range gradingMap {
		if averageScore >= float64(threshold) {
			grade = letterGrade
			break
		}
	}

	// Print results
	fmt.Printf("Student Name: %s\n", studentName)
	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Grade: %s\n", grade)
}
