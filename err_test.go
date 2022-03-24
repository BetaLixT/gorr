package gorr

import (
	"fmt"
	"testing"
)

func TestUnexpectedError(t *testing.T) {
	err := UnexpectedError(fmt.Errorf("Test error"))
	if err.StatusCode != 500 {
		t.Error("Incorrect status code")
		t.FailNow()
	}
	if err.ErrorCode != 10000 {
		t.Error("Incorrect error code")
		t.FailNow()
	}
	if err.ErrorMessage != "UnexpectedError" {
		t.Error("Incorrect error message")
		t.FailNow()
	}
	if err.ErrorDetail != "Test error" {
		t.Error("Incorrect error detail")
		t.FailNow()
	}
}

func TestNotImplemented(t *testing.T) {
	err := NotImplemented()
	if err.StatusCode != 501 {
		t.Error("Incorrect status code")
		t.FailNow()
	}
	if err.ErrorCode != 10001 {
		t.Error("Incorrect error code")
		t.FailNow()
	}
	if err.ErrorMessage != "NotImplemented" {
		t.Error("Incorrect error message")
		t.FailNow()
	}
	if err.ErrorDetail != "" {
		t.Error("Incorrect error detail")
		t.FailNow()
	}
}
