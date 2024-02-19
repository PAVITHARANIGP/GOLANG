package main

import "fmt"

type student struct {
	id    int
	name  string
	age   int
	grade string
}

func main() {
	var students []student
	var id int
	var name, s string
	var age int
	var grade string
	var opt bool = true
	for opt {
		fmt.Print("\n Enter student ID :")
		fmt.Scan(&id)
		fmt.Print("\n Enter student name :")
		fmt.Scan(&name)
		fmt.Print("\n Enter student age :")
		fmt.Scan(&age)
		fmt.Print("\n Enter student grade :")
		fmt.Scan(&grade)
		students = append(students, addDetails(id, name, age, grade))
		fmt.Print("\n Do you want to add more?(y/n) : ")
		fmt.Scan(&s)
		if s != "y" {
			opt = false
		}
	}
}
func addDetails(id int, name string, age int, grade string) student {
	var item student
	item.id = id
	item.name = name
	item.age = age
	item.grade = grade
	return item
}