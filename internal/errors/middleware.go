package errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	httpCode int
	errorStr string
}

var customError *CustomError

func (customError *CustomError) Error() string {
	return customError.errorStr
}

func (customError *CustomError) GetCode() int {
	return customError.httpCode
}

func New(httpCode int, errorStr string) *CustomError {
	return &CustomError{
		httpCode: httpCode,
		errorStr: errorStr,
	}
}

func ErrorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		for errors.Unwrap(err) != nil {
			err = errors.Unwrap(err)
		}

		if errors.As(err, &customError) {
			cErr := err.(*CustomError)
			c.JSON(cErr.httpCode, gin.H{
				"error_message": cErr.errorStr,
			})
		} else {
			c.Status(http.StatusInternalServerError)
		}

	}
}
