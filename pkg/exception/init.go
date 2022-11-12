package exception

import "net/http"

func NewBaseException(statusCode int) *Base {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	return &Base{
		Code:   statusCode,
		Module: BaseModule,
	}
}

func NewAuthenticationException() *Base {
	return NewBaseException(http.StatusForbidden)
}

func NewAuthorizationException() *Base {
	return NewBaseException(http.StatusUnauthorized)
}

func NewBadRequestException() *Base {
	return NewBaseException(http.StatusBadRequest)
}

func NewNotFoundException() *Base {
	return NewBaseException(http.StatusNotFound)
}
