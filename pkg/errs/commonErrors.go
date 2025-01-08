package errs

import "github.com/gin-gonic/gin"

func InternalServerError(c *gin.Context, messages ...string) {
	error := CommonError{
		StatusCode: 500,
		Status:     "internal server error",
		Messages:   messages,
	}
	c.Error(&error)
	return
}
