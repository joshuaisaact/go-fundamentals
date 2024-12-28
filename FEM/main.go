package main

import (
	"fmt"

	"frontendmasters.com/go/server/model"
)

func main() {
	josh := model.Instructor{Id: 3, LastName: "Tuddenham", Score: 100} // new Instructor()
	josh.FirstName = "Josh"

	kyle := model.NewInstructor("kyle", "miller")

	goCourse := model.Course{Id: 2, Name: "Go Fundamentals", Instructor: josh}

	fmt.Printf("%v", goCourse)

	swiftWS := model.Workshop{}

	print(kyle.Print())
	print(josh.Print())

	var courses [2]model.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}
}
