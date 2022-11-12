package request

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
	"github.com/resyahrial/go-user-management/pkg/validator"
)

type PaginatedQueryParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit" validate:"gte=0,omitempty"`
}

func (q *PaginatedQueryParams) CastToPaginatedQueryParamsEntity() (queryParams *entities.PaginatedQueryParams, err error) {
	if mapErr := validator.Validate(q); len(mapErr) > 0 {
		err = exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(mapErr)
		return
	}
	if err = mapstructure.Decode(q, &queryParams); err != nil {
		return
	}
	if queryParams.Limit == 0 {
		queryParams.Limit = 10
	}
	return
}
