package validator_test

import (
	"testing"

	"github.com/resyahrial/go-template/pkg/validator"
	"github.com/stretchr/testify/suite"
)

type ValidatorSuite struct {
	suite.Suite
}

func TestValidator(t *testing.T) {
	suite.Run(t, new(ValidatorSuite))
}

type person struct {
	Name                string `json:"name" validate:"required,min=4,max=15"`
	Email               string `json:"email" validate:"required,email"`
	Age                 int    `validate:"required,numeric,min=18"`
	DriverLicenseNumber string `json:"-" validate:"omitempty,len=12,numeric"`
	PhoneNumber         string `json:"phoneNumber" validate:"phoneNumber"`
}

func (s *ValidatorSuite) TestValidate() {
	p := &person{
		Name:                "Joe",
		Email:               "dummyemail",
		Age:                 0,
		DriverLicenseNumber: "",
		PhoneNumber:         "08123451235",
	}

	mapErr := validator.Validate(p)
	s.Len(mapErr, 4)
	s.Len(mapErr["name"], 1)
	s.Len(mapErr["email"], 1)
	s.Len(mapErr["age"], 0)
	s.Len(mapErr["Age"], 1)
	s.Len(mapErr["DriverLicenseNumber"], 0)
	s.Len(mapErr["phoneNumber"], 1)
	s.Equal("PhoneNumber must starts with '62' and have length between 10-14 character", mapErr["phoneNumber"][0])
}
