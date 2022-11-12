package hasher_test

import (
	"testing"

	"github.com/resyahrial/go-user-management/pkg/hasher"
	"github.com/stretchr/testify/suite"
)

type HasherTestSuite struct {
	suite.Suite
	h *hasher.Hasher
}

func TestHasher(t *testing.T) {
	suite.Run(t, new(HasherTestSuite))
}

func (s *HasherTestSuite) SetupTest() {
	s.h = hasher.New(14)
}

func (s *HasherTestSuite) TestHashPassword() {
	hashedPass, err := s.h.HashPassword("secret")
	s.Nil(err)
	s.NotEqual("secret", hashedPass)
}

func (s *HasherTestSuite) TestCheckPasswordHash() {
	hashedPass, err := s.h.HashPassword("secret")
	s.Nil(err)
	isEqual := s.h.CheckPasswordHash("secret", hashedPass)
	s.True(isEqual)
	isEqual = s.h.CheckPasswordHash("notsecret", hashedPass)
	s.False(isEqual)
}
