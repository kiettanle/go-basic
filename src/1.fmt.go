package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func getName(name string) string {
	return "I'm " + name
}

func init() {
	fmt.Println(getName("Kiệt"))
	var message string = "there!\n"
	fmt.Printf("Hi, %v", message)
}
