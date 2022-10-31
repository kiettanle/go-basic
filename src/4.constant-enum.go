package main

import "fmt"

// int mapping
const (
	Summer int = 0
	Autumn     = 1
	Winter     = 2
	Spring     = 3
)

type Season uint8

const (
	Summer1 Season = iota
	Autumn1
	Winter1
	Spring1
)

// string mapping
const (
	Summer2 string = "summer"
	Autumn2        = "autumn"
	Winter2        = "winter"
	Spring2        = "spring"
)

// Implement the String method for the Season
func (s Season) String() string {
	switch s {
	case Summer1:
		return "summer"
	case Autumn1:
		return "autumn"
	case Winter1:
		return "winter"
	case Spring1:
		return "spring"
	}
	return "unknown"
}

func printSeason(s Season) {
	fmt.Println("season: ", s)
}

func main() {
	fmt.Println("Constant...")
	const MAX_LENGTH int = 255

	fmt.Println(MAX_LENGTH)

	// Enum
	printSeason(Winter1)
}
