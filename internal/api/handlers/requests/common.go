package request

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
)

//go:generate mockgen -destination=mocks/mock.go -source=common.go GinContextDefaultQueryAdapter
type GinContextDefaultQueryAdapter interface {
	DefaultQuery(key, defaultValue string) string
}

type PaginatedQueryParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (q *PaginatedQueryParams) CastToPaginatedQueryParamsEntity() (queryParams *entities.PaginatedQueryParams, err error) {
	if err = mapstructure.Decode(q, &queryParams); err != nil {
		return
	}
	return
}

func (q *PaginatedQueryParams) ParseQueryParams(c GinContextDefaultQueryAdapter) (err error) {
	if q.Page, err = strconv.Atoi(c.DefaultQuery("page", "0")); err != nil {
		return
	}
	if q.Limit, err = strconv.Atoi(c.DefaultQuery("limit", "10")); err != nil {
		return
	}
	return
}
