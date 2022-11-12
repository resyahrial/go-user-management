package request_test

import (
	"testing"

	request "github.com/resyahrial/go-user-management/internal/api/handlers/requests"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
	"github.com/stretchr/testify/suite"
)

type PaginatedQueryParamsTestSuite struct {
	suite.Suite
}

func TestPaginatedQueryParams(t *testing.T) {
	suite.Run(t, new(PaginatedQueryParamsTestSuite))
}

func (s *PaginatedQueryParamsTestSuite) SetupTest() {
}

func (s *PaginatedQueryParamsTestSuite) TestConvertToPaginatedQueryParamsEntity() {
	testCases := []struct {
		name           string
		input          *request.PaginatedQueryParams
		expectedOutput *entities.PaginatedQueryParams
		expectedError  error
	}{
		{
			name: "should create basic paginated query params",
			input: &request.PaginatedQueryParams{
				Page:  0,
				Limit: 5,
			},
			expectedOutput: &entities.PaginatedQueryParams{
				Page:  0,
				Limit: 5,
			},
		},
		{
			name: "should give default value of limit",
			input: &request.PaginatedQueryParams{
				Page:  0,
				Limit: 0,
			},
			expectedOutput: &entities.PaginatedQueryParams{
				Page:  0,
				Limit: 10,
			},
		},
		{
			name: "should return error when not pass validation",
			input: &request.PaginatedQueryParams{
				Page:  0,
				Limit: -10,
			},
			expectedError: exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(map[string][]string{
				"limit": {
					"Limit must be 0 or greater",
				},
			}),
		},
	}

	for _, tc := range testCases {
		queryParams, err := tc.input.CastToPaginatedQueryParamsEntity()
		s.Run(tc.name, func() {
			if tc.expectedError == nil {
				s.Nil(err)
			} else {
				s.Equal(tc.expectedError.Error(), err.Error())
			}
			s.EqualValues(tc.expectedOutput, queryParams)
		})
	}
}
