package middleware

import (
	"errors"

	"github.com/H0lyDiv3r/go-headless-cms/pkg/errs"
	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			var commonError *errs.CommonError
			err := c.Errors[0].Err
			if errors.As(err, &commonError) {
				c.JSON(int(commonError.StatusCode), commonError)
				c.Abort()
			} else {
				c.JSON(500, gin.H{
					"errors": "unknown error",
				})
				c.Abort()
			}
		}
		c.Next()
	}
}
