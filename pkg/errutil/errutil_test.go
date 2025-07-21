package errutil_test

import (
	"errors"
	"testing"

	"github.com/u104rak1/gostd/pkg/errutil"
)

func TestWrap(t *testing.T) {
	t.Run("Wrapでエラーメッセージがラップされ、Unwrapで元のエラーが取得できる", func(t *testing.T) {
		originalErr := errors.New("base error")
		wrappedErr := errutil.Wrap("wrapped error", originalErr)

		// Error()の内容を確認
		wantMsg := "wrapped error: base error"
		if wrappedErr.Error() != wantMsg {
			t.Errorf("Error() = %v, want %v", wrappedErr.Error(), wantMsg)
		}

		// Unwrapで元のエラーが取得できることを確認
		if errors.Unwrap(wrappedErr) != originalErr {
			t.Errorf("Unwrap() = %v, want %v", errors.Unwrap(wrappedErr), originalErr)
		}
	})
}
