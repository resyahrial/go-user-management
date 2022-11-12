package request_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	request "github.com/resyahrial/go-user-management/internal/api/handlers/requests"
	request_mock "github.com/resyahrial/go-user-management/internal/api/handlers/requests/mocks"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/stretchr/testify/suite"
)

type PaginatedQueryParamsTestSuite struct {
	suite.Suite
	ginContextAdapter *request_mock.MockGinContextDefaultQueryAdapter
}

func TestPaginatedQueryParams(t *testing.T) {
	suite.Run(t, new(PaginatedQueryParamsTestSuite))
}

func (s *PaginatedQueryParamsTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.ginContextAdapter = request_mock.NewMockGinContextDefaultQueryAdapter(ctrl)
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

func (s *PaginatedQueryParamsTestSuite) TestParseQueryParams() {
	testCases := []struct {
		name           string
		mockPage       string
		mockLimit      string
		expectedOutput *entities.PaginatedQueryParams
		expectedError  error
	}{
		{
			name:      "should success parse query param",
			mockPage:  "0",
			mockLimit: "5",
			expectedOutput: &entities.PaginatedQueryParams{
				Page:  0,
				Limit: 5,
			},
		},
	}

	for _, tc := range testCases {
		s.ginContextAdapter.EXPECT().DefaultQuery("page", "0").Return(tc.mockPage)
		s.ginContextAdapter.EXPECT().DefaultQuery("limit", "10").Return(tc.mockLimit)

		queryParams := &request.PaginatedQueryParams{}
		err := queryParams.ParseQueryParams(s.ginContextAdapter)
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
