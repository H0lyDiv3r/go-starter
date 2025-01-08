package errs

import "github.com/gin-gonic/gin"

func BadRequest(c *gin.Context, messages ...string) {
	error := CommonError{
		StatusCode: 400,
		Status:     "bad request",
		Messages:   messages,
	}
	c.Error(&error)
	return
}
func Unauthorized(c *gin.Context, messages ...string) {
	error := CommonError{
		StatusCode: 401,
		Status:     "unauthorized",
		Messages:   messages,
	}
	c.Error(&error)
	return
}
func Forbidden(c *gin.Context, messages ...string) {
	error := CommonError{
		StatusCode: 403,
		Status:     "forbidden",
		Messages:   messages,
	}
	c.Error(&error)
	return
}
func NotFound(c *gin.Context, messages ...string) {
	error := CommonError{
		StatusCode: 403,
		Status:     "not found",
		Messages:   messages,
	}
	c.Error(&error)
	return
}
func Conflict(c *gin.Context, messages ...string) {
	error := CommonError{
		StatusCode: 409,
		Status:     "conflict",
		Messages:   messages,
	}
	c.Error(&error)
	return
}

func InternalServerError(c *gin.Context, messages ...string) {
	error := CommonError{
		StatusCode: 500,
		Status:     "internal server error",
		Messages:   messages,
	}
	c.Error(&error)
	return
}

