package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/pkg/exception"
)

type Success struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
}

type Failure struct {
	StatusCode int         `json:"-"`
	Error      interface{} `json:"error"`
}

func HandleSuccess(data interface{}) *Success {
	return &Success{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func HandleError(c *gin.Context, err error) *Failure {
	ginErr := c.Error(err)

	switch typeErr := ginErr.Err.(type) {
	case *exception.Base:
		typeErr.LogError()
		if typeErr.CollectionMessage != nil && len(typeErr.CollectionMessage) > 0 {
			return generateFailure(typeErr.Code, typeErr.CollectionMessage)
		}
		return generateFailure(typeErr.Code, map[string]interface{}{
			"message": typeErr.Error(),
		})
	default:
		return generateFailure(http.StatusInternalServerError, map[string]interface{}{
			"message": typeErr.Error(),
		})
	}
}

func generateFailure(statusCode int, errResponse interface{}) *Failure {
	return &Failure{
		StatusCode: statusCode,
		Error:      errResponse,
	}
}
