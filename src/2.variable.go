package main

import "fmt"

// Global scope
var calledTime = 0

func main() {
	var firstName string
	firstName = "Kiệt"
	var middleName string = "Tấn"
	var lastName = "Lê"
	fullName := getFullName(firstName, middleName, lastName)
	fmt.Println(fullName)
}

func getFullName(firstName string, middleName string, lastName string) string {
	fullName := firstName + middleName + lastName

	return fullName
}

func init() {
	calledTime = 1
	fmt.Println(calledTime)
}
