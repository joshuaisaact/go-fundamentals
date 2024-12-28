package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]bool

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[8] = "USA"

	qty := len(Countries)

	fmt.Println("Countries saved", qty)

	names := []string{"Mary", "John"}
	names2 := append(names, "Carol")
	println(len(names2))

	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	fmt.Println(wellKnownPorts)
}
