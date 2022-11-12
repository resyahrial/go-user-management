package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-user-management/pkg/exception"
)

const (
	SuccessKey   = "SuccessKey"
	FailureKey   = "FailureKey"
	PaginatedKey = "CountKey"
)

type PaginatedResultValue struct {
	Page  int
	Limit int
	Count int64
}

type success struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
	PageInfo   pageInfo    `json:"page_info,omitempty"`
}

type pageInfo struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
}

type failure struct {
	StatusCode int         `json:"-"`
	Error      interface{} `json:"error"`
}

func (m *Middleware) ResponseWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if val, ok := c.Get(FailureKey); ok {
			if err, ok := val.(error); ok {
				handleError(c, err)
				return
			}
		}

		if data, ok := c.Get(SuccessKey); ok {
			if count, ok := c.Get(PaginatedKey); ok {
				if paginatedResultValue, ok := count.(PaginatedResultValue); ok {
					handleSuccessPaginated(c, data, paginatedResultValue)
				}
			} else {
				handleSuccess(c, data)
			}
			return
		}
	}
}

func handleError(c *gin.Context, err error) {
	var (
		f *failure
	)
	ginErr := c.Error(err)

	switch typeErr := ginErr.Err.(type) {
	case *exception.Base:
		typeErr.LogError()
		if typeErr.CollectionMessage != nil && len(typeErr.CollectionMessage) > 0 {
			f = generateFailure(typeErr.Code, typeErr.CollectionMessage)
		} else {
			f = generateFailure(typeErr.Code, map[string]interface{}{
				"message": typeErr.Error(),
			})
		}
	default:
		f = generateFailure(http.StatusInternalServerError, map[string]interface{}{
			"message": typeErr.Error(),
		})
	}

	c.JSON(f.StatusCode, f)
}

func generateFailure(statusCode int, errResponse interface{}) *failure {
	return &failure{
		StatusCode: statusCode,
		Error:      errResponse,
	}
}

func handleSuccess(c *gin.Context, data interface{}) {
	s := &success{
		StatusCode: http.StatusOK,
		Data:       data,
	}
	c.JSON(s.StatusCode, s)
}

func handleSuccessPaginated(c *gin.Context, data interface{}, paginatedResultValue PaginatedResultValue) {
	totalPage := 1
	if paginatedResultValue.Limit < int(paginatedResultValue.Count) {
		addtional := int(paginatedResultValue.Count) / paginatedResultValue.Limit
		if int(paginatedResultValue.Count)%paginatedResultValue.Limit == 0 {
			addtional -= 1
		}
		totalPage += addtional
	}
	s := &success{
		StatusCode: http.StatusOK,
		Data:       data,
		PageInfo: pageInfo{
			CurrentPage: paginatedResultValue.Page + 1,
			TotalPage:   totalPage,
		},
	}
	c.JSON(s.StatusCode, s)
}
