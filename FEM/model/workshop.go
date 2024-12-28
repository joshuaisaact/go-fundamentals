package model

import "time"

// Embedding
type Workshop struct {
	Course
	Date time.Time
}

func (c Workshop) SignUp() bool {
	return true
}

// Factory
func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor
	return w
}
