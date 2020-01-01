package restserver

import (
	e "github.com/gabrielpsilva/sonicspeed/error"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Context struct {
	//Rest *gin.Context
	*gin.Context
}

func ContextWrapper(ginContext *gin.Context) *Context {
	return &Context{ginContext}
}

func (c *Context) Log() *logrus.Entry {
	logger := logrus.WithFields(logrus.Fields{
		"requestID": c.RequestID(),
	})
	return logger
}


func (c *Context) RequestID() string {
	if val, ok := c.Value("X-Request-Id").(string); ok {
		return val
	}
	return ""
}


func (c *Context) ErrorHTTP(err e.Error){

	statusCode := getHttpCode(err)
	if statusCode > 499 {
		c.Log().Error(err.Description)
	} else if statusCode > 399 {
		c.Log().Warn(err.Description)
	} else {
		c.Log().Info(err.Description)
	}
	c.JSON(statusCode, &err)
}


func getHttpCode(e e.Error) int{
	if e.Code < 999 {
		e.Description = "error throwing anothe	r error (Jezz!)"
		e.Documentation = "Error code can not be > 1000"
		e.Code = 1500
	}
	return e.Code % 1e3
}

/*
func (e *Error) formatDesc(err string) string {
	return err + e.Description
}


func (c *Context) Err() error {
	return c.Err()
}

func (c *Context) Done() <-chan struct{} {
	return c.Done()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.Value(key)
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Deadline()
}
*/
