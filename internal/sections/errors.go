package sections

import "fmt"

type ErrorNotFound struct {
	Id int
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("section %d not found in database", e.Id)
}

type ErrorConflict struct {
	SectionNumber int
}

func (e *ErrorConflict) Error() string {
	return fmt.Sprintf("a section with number %d already exists", e.SectionNumber)
}
