package gorr

import "fmt"

// TODO Make this configurable
const (
	unxCd = 10000
	unxMsg = "UnexpectedError"
	unimCd = 10001
	unimMsg = "NotImplemented"
)

type ErrorCode struct {
	Code    int
	Message string
}

type Error struct {
	ErrorCode   ErrorCode
	StatusCode  int    `json:"-"`
	ErrorDetail string `json:"errorDetail,omitempty"`
}

func (err *Error) Error() string {
	return fmt.Sprintf(
		"%d [%s]: %s | Requested Status Code: %d",
		err.ErrorCode.Code,
		err.ErrorCode.Message,
		err.ErrorDetail,
		err.StatusCode,
	)
}

func NewError(
	code ErrorCode,
	statusCode int,
	desc string,
) *Error {
	return &Error{
		StatusCode:  statusCode,
		ErrorCode:   code,
		ErrorDetail: desc,
	}
}

func NewBadRequestError(
	code ErrorCode,
	desc string,
) *Error {
	return &Error{
		StatusCode:  400,
		ErrorCode:   code,
		ErrorDetail: desc,
	}
}

func NewServerError(
	code ErrorCode,
	desc string,
) *Error {
	return &Error{
		StatusCode:  500,
		ErrorCode:   code,
		ErrorDetail: desc,
	}
}

func NewUnexpectedError(e error) *Error {
	return &Error{
		StatusCode: 500,
		ErrorCode: ErrorCode{
			Code:    unxCd,
			Message: unxMsg,
		},
		ErrorDetail: e.Error(),
	}
}

func NewNotImplemented() *Error {
	return &Error{
		StatusCode: 501,
		ErrorCode: ErrorCode{
			Code:    unimCd,
			Message: unimMsg,
		},
	}
}
