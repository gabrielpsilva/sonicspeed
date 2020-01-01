package error

import (
	"fmt"
)

const (
	NoSolution = "Argh, a solution is catalogued yet!"
)



// Error is the general Error format of this application
type Error struct {
	Code          int
	Description   string
	Documentation string
}

// NewError Returns new Error struct
func New(code int, description, documentation string) Error {
	e := Error{}
	e.Code = code
	e.Description = description
	e.Documentation = documentation
	if len(e.Documentation)  <= 0 {
		e.Documentation = NoSolution
	}
	return e
}

// NewErrorE Returns new Error struct
func NewE(code int, err error, documentation string) Error {
	return New(code, fmt.Sprintf("%s", err), documentation)
}


func (e *Error) Error() string {
	//return fmt.Sprintf("ERROR: \ncode: %d\ndesc:%s\ndoc:%s", e.code, e.description, e.documentation)
	return fmt.Sprintf("code=%d desc=\"%s\" doc=\"%s\"", e.Code, e.Description, e.Documentation)
}


