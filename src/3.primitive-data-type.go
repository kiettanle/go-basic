package main

import (
	"fmt"
)

func main() {
	fmt.Println("Primitive data type: \n1. Boolean, \n2. Numberics, \n3. Text")

	var firstNumber = 10
	var secondNumber = 3

	fmt.Printf("Sum: %v + %v is %v\n", firstNumber, secondNumber, sum(firstNumber, secondNumber))
	fmt.Printf("Sub: %v - %v is %v\n", firstNumber, secondNumber, sub(firstNumber, secondNumber))
	fmt.Printf("Mul: %v * %v is %v\n", firstNumber, secondNumber, mul(firstNumber, secondNumber))
	fmt.Printf("Div: %v / %v is %v\n", firstNumber, secondNumber, div(firstNumber, secondNumber))
	fmt.Printf("Mod: %v mod %v is %v\n", firstNumber, secondNumber, mod(firstNumber, secondNumber))
	fmt.Println()

	// 10 = 1010, 3 = 0011 (2)
	//     1010
	// AND 0011
	//     0010 = 2 (10)
	fmt.Printf("And: %v and %v is %v\n", firstNumber, secondNumber, and(firstNumber, secondNumber))

	// 10 = 1010, 3 = 0011 (2)
	//     1010
	// OR  0011
	//     1011 = 2 (10)
	fmt.Printf("Or: %v or %v is %v\n", firstNumber, secondNumber, or(firstNumber, secondNumber))

	// 10 = 1010, 3 = 0011 (2)
	//     1010
	// XOR 0011
	//     0100 = 8 (108)
	fmt.Printf("Xor: %v XOR %v is %v\n", firstNumber, secondNumber, xor(firstNumber, secondNumber))

	// 10 = 1010, 3 = 0011 (2)
	//     1010
	// NOT 0011
	//     1001 = 9 (10)
	fmt.Printf("Not: %v NOT %v is %v\n", firstNumber, secondNumber, not(firstNumber, secondNumber))

	// Shiffleft
	// 10 = 0000 1010 ==> 0101 0000 = 80 (10)
	fmt.Printf("ShiffLeft: %v << %v is %v\n", firstNumber, 3, firstNumber<<3)

	// Shiff Right
	// 10 = 0000 1010 ==> 0000 0001 = 1 (10)
	fmt.Printf("ShiffRight: %v >> %v is %v\n", firstNumber, 3, firstNumber>>3)

	// Complex Complex64 Complex128

	var f complex64 = 1 + 2i // Float32 + Float32
	fmt.Printf("Complex64 %v %T\n", f, f)
	fmt.Printf("Real %v type %T\n", real(f), real(f))
	fmt.Printf("Imag %v type %T\n", imag(f), imag(f))

	g := complex(5, 10)
	fmt.Printf("Type %T value = %v \n", g, g)

	// STRING
	firstName := "Kiet"
	lastName := "Le"

	fmt.Println(firstName + lastName)

	// BYTE: UTF-8

	fullName := []byte(firstName + lastName)

	fmt.Printf("%T value % v\n", fullName, fullName)

	fmt.Printf("%T value as string is % v\n", fullName[0], string(fullName[0]))

	// RUNE: UTF-32

	text := []rune(firstName)

	fmt.Printf("%T value %v\n", text, text)
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	if b == 0 {
		return 0
	}

	return a / b
}

func mod(a int, b int) int {
	if b == 0 {
		return 0
	}

	return a % b
}

func and(a int, b int) int {
	return a & b
}

func or(a int, b int) int {
	return a | b
}

func not(a int, b int) int {
	return a ^ b
}

func xor(a int, b int) int {
	return a &^ b
}
