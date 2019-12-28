package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestIDWithOpts Allows developer to send a list of possible requestID headers
func RequestIDWithOpts(headers []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		for _, opt := range headers {
			header := c.Request.Header.Get(opt)
			if header != "" {
				c.Set("X-Request-Id", header)
				return
			}
		}
		uid, _ := uuid.NewV4()
		c.Set("X-Request-Id", uid.String())
	}}
