package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")
	fmt.Println("Which operation do you want to perform? (add, subtract, multiply, divide)")
	fmt.Scanf("%s", &operation)
	fmt.Print("Enter first number\n")
	fmt.Scanf("%d\n", &number1)
	fmt.Print("Enter second number\n")
	fmt.Scanf("%d\n", &number2)
	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "subtract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
	}
}
