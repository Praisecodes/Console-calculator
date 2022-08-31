package main

import (
	"fmt"
)

func Add(first int64, second int64) (result int64) {
	result = first + second
	return
}

func substract(first int64, second int64) (result int64) {
	result = first - second
	return
}

func divide(first int64, second int64) (result int64) {
	if second == 0 {
		fmt.Println("Division by zero not allowed")
	} else {
		result = first / second
	}
	return
}

func multiply(first int64, second int64) (result int64) {
	result = first * second
	return
}

func performAction(switchCase string, first int64, second int64) {
	switch switchCase {
	case "1":
		fmt.Println(Add(first, second))

	case "2":
		fmt.Println(substract(first, second))

	case "3":
		fmt.Println(divide(first, second))

	case "4":
		fmt.Println(multiply(first, second))

	default:
		fmt.Println("Invalid choice")
	}
}

func main() {
	var choice string
	var firstDigit, secondDigit int64
	fmt.Println("Hello This is a calculator in Go Lang")
	fmt.Print("Please select an option:\n1. Addition\n2.Substraction\n3. Division\n4. Multiplication\nChoice: ")
	fmt.Scanln(&choice)
	if firstDigit != 0 {
		fmt.Printf("First Digit is %v", firstDigit)
		fmt.Print("Enter your secodd digit: ")
		fmt.Scanln(&secondDigit)
		performAction(choice, firstDigit, secondDigit)
	} else {
		fmt.Print("Enter your first digit: ")
		fmt.Scanln(&firstDigit)
		fmt.Print("Enter your secode digit: ")
		fmt.Scanln(&secondDigit)
		performAction(choice, firstDigit, secondDigit)
	}
}
