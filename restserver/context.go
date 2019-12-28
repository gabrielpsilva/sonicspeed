package restserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type Context struct {
	Rest *gin.Context
}

func ContextWrapper(ginContext *gin.Context) *Context {
	return &Context{Rest: ginContext}
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

func (c *Context) Done() <-chan struct{} {
	return c.Rest.Done()
}

func (c *Context) Err() error {
	return c.Rest.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.Rest.Value(key)
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Rest.Deadline()
}