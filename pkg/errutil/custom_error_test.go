package errutil_test

import (
	"errors"
	"testing"

	"github.com/u104rak1/gostd/pkg/errutil"
)

func TestCustomError_Error(t *testing.T) {
	t.Run("Error()は設定したメッセージを返す", func(t *testing.T) {
		err := errutil.NewErrBadRequest("bad request")
		if err.Error() != "bad request" {
			t.Errorf("Error() = %v, want %v", err.Error(), "bad request")
		}
	})
}

func TestCustomError_Unwrap(t *testing.T) {
	t.Run("WrapでラップしたエラーはUnwrap()で元のエラーを取得できる", func(t *testing.T) {
		originalErr := errors.New("original error")
		wrappedErr := errutil.Wrap("bad request", originalErr)
		unwrapped := errors.Unwrap(wrappedErr)
		if unwrapped != originalErr {
			t.Errorf("Unwrap() = %v, want %v", unwrapped, originalErr)
		}
	})
}

func TestCustomError_Is(t *testing.T) {
	t.Run("同じ型のcustomError同士はerrors.Isでtrueになる", func(t *testing.T) {
		err1 := errutil.NewErrBadRequest("bad request")
		err2 := errutil.NewErrBadRequest("another bad request")
		if !errors.Is(err1, err2) {
			t.Error("errors.Is should be true for same custom error type")
		}
	})
	t.Run("異なる型のcustomError同士はerrors.Isでfalseになる", func(t *testing.T) {
		err1 := errutil.NewErrBadRequest("bad request")
		err3 := errutil.NewErrUnauthorized("unauthorized")
		if errors.Is(err1, err3) {
			t.Error("errors.Is should be false for different custom error type")
		}
	})
}

func TestCustomError_Constructors(t *testing.T) {
	tests := []struct {
		caseName string
		fn       func(string) error
		message  string
	}{
		{
			caseName: "NewErrBadRequestでエラーを生成しError()が正しいメッセージを返す",
			fn: func(msg string) error {
				return errutil.NewErrBadRequest(msg)
			},
			message: "bad request",
		},
		{
			caseName: "NewErrUnauthorizedでエラーを生成しError()が正しいメッセージを返す",
			fn: func(msg string) error {
				return errutil.NewErrUnauthorized(msg)
			},
			message: "unauthorized",
		},
		{
			caseName: "NewErrForbiddenでエラーを生成しError()が正しいメッセージを返す",
			fn: func(msg string) error {
				return errutil.NewErrForbidden(msg)
			},
			message: "forbidden",
		},
		{
			caseName: "NewErrNotFoundでエラーを生成しError()が正しいメッセージを返す",
			fn: func(msg string) error {
				return errutil.NewErrNotFound(msg)
			},
			message: "not found",
		},
		{
			caseName: "NewErrRequestTimeoutでエラーを生成しError()が正しいメッセージを返す",
			fn: func(msg string) error {
				return errutil.NewErrRequestTimeout(msg)
			},
			message: "timeout",
		},
		{
			caseName: "NewErrConflictでエラーを生成しError()が正しいメッセージを返す",
			fn: func(msg string) error {
				return errutil.NewErrConflict(msg)
			},
			message: "conflict",
		},
		{
			caseName: "NewErrUnprocessableEntityでエラーを生成しError()が正しいメッセージを返す",
			fn: func(msg string) error {
				return errutil.NewErrUnprocessableEntity(msg)
			},
			message: "unprocessable",
		},
		{
			caseName: "NewErrTooManyRequestsでエラーを生成しError()が正しいメッセージを返す",
			fn: func(msg string) error {
				return errutil.NewErrTooManyRequests(msg)
			},
			message: "too many",
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.caseName, func(t *testing.T) {
			err := tt.fn(tt.message)
			if err.Error() != tt.message {
				t.Errorf("Error() = %v, want %v", err.Error(), tt.message)
			}
		})
	}
}
