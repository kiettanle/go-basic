package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("If else")

	num1 := rand.Float64()
	num2 := rand.Float64()
	num3 := rand.Float64()

	fmt.Println(num1, num2)

	if num1 > num2 {
		fmt.Println("num1 > num2")
	} else if num1 < num2 {
		fmt.Println("num1 < num2")
	} else {
		fmt.Println("num1 = num2")
	}
	switch num3 {
	case 0:
		fmt.Println("num3 = 0")
		fallthrough // Keep going
	case 0.5:
		fmt.Println("num3 = 0.5")
		fallthrough
	default:
		fmt.Println("Nothing...")
	}

	// Switch with dynamic condition
	switch {
	case num3 > 0:
		fmt.Println("num3 > 0")
		fallthrough // Keep going
	case num3 > 0.5:
		fmt.Println("num3 > 0.5")
	default:
		fmt.Println("Nothing...")
	}

}
