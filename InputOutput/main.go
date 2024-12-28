package main

import "fmt"

// Global variable
var url = "https://frontendmasters.com"

func init() {
	fmt.Println("A")
}

func init() {
	fmt.Println("B")
}

// function that returns multiple values (NOT an array)
func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

// the name is internal to the function
func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return
}

func birthday(age int) {
	age++
}

func birthdayWithPointer(pointerAge *int) {
	// Panic CLOSES THE APP. Do not use as error management.
	if *pointerAge > 140 {
		panic("Too old too be true")
	}

	// PrintF prints formatted template strings to console
	fmt.Printf("The pointer is %v and the value is %v \n", pointerAge, *pointerAge)

	*pointerAge++
}

func main() {
	stateTax, _ := calculateTax(100)
	fmt.Println(stateTax)

	// Defer moves to the end of the call stack. Typically used for example when closing a DB connection at the end of the function.
	defer fmt.Println("Bye!!")
	defer fmt.Println("Good ")

	age := 22
	fmt.Println(age)
	// We are incrementing the COPY not the variable. This doesn't increment by default, every value is passed as a copy.
	birthday(age)

	// This is passing the reference
	birthdayWithPointer(&age)

	fmt.Println(age)

	// Print the pointer
	fmt.Println(&age)

	fmt.Print("Hello ")
	// Function-scoped variables
	const maxSpeed byte = 60
	message := "Hello from a module"
	price := 34.4
	fmt.Print(message, price, url, maxSpeed)
	PrintData()
	fmt.Println(name)
}
