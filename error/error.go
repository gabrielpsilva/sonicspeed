package error

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func ThrowHTTPError(c *gin.Context, err Error){
	c.AbortWithError(err.getHttpCode(), &err)
}

func ThrowHTTPErrorE(c *gin.Context, err Error, append error){
	if &err != nil {
		err.Description = fmt.Sprintf("%s: %v", err.Description, append)
	}
	//c.AbortWithError(err.getHttpCode(), &err)
	c.AbortWithStatusJSON(err.getHttpCode(), &err)
}

func (e *Error) getHttpCode() int{
	if e.Code < 999 {
		e.Description = "error throwing another error (Jezz!)"
		e.Documentation = "Error code can not be > 1000"
		e.Code = 1500
	}
	return e.Code % 1e3
}


func (e *Error) formatDesc(err string) string {
	return err + e.Description
}

func (e *Error) Error() string {
	//return fmt.Sprintf("ERROR: \ncode: %d\ndesc:%s\ndoc:%s", e.code, e.description, e.documentation)
	return fmt.Sprintf("code=%d desc=\"%s\" doc=\"%s\"", e.Code, e.Description, e.Documentation)
}


