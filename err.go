package gorr

import "fmt"

// TODO Make this configurable
const unxCd, unxMsg, unimCd, unimMsg = 10000, "UnexpectedError", 10001, "NotImplemented"

type Error struct {
	StatusCode   int    `json:"-"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ErrorDetail  string `json:"errorDetail,omitempty"`
}

func (err *Error) Error() string {
	return fmt.Sprintf(
		"%d [%s]: %s | Requested Status Code: %d",
		err.ErrorCode,
		err.ErrorMessage,
		err.ErrorDetail,
		err.StatusCode,
	)
}

func UnexpectedError(e error) Error {
	return Error{
		StatusCode:   500,
		ErrorCode:    unxCd,
		ErrorMessage: unxMsg,
		ErrorDetail:  e.Error(),
	}
}

func NotImplemented() Error {
	return Error{
		StatusCode:   501,
		ErrorCode:    unimCd,
		ErrorMessage: unimMsg,
	}
}
