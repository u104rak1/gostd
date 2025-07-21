package errutil

// customError はカスタムエラーのベースとなる構造体です。
// ジェネリック型Tを使用して、任意のエラー型を扱うことができます。
type customError[T any] struct {
	msg string
	err error
}

// newCustomError は新しいカスタムエラーを作成します。
func newCustomError[T any](message string) *customError[T] {
	return &customError[T]{msg: message}
}

// Error はエラーメッセージを返します。
func (e *customError[T]) Error() string {
	return e.msg
}

// Unwrap は元のエラーを返します。
func (e *customError[T]) Unwrap() error {
	return e.err
}

// Is はエラーの型を比較します。
func (e *customError[T]) Is(target error) bool {
	_, ok := target.(*customError[T])
	return ok
}

type (
	badRequestType          struct{}
	unauthorizedType        struct{}
	forbiddenType           struct{}
	notFoundType            struct{}
	requestTimeoutType      struct{}
	conflictType            struct{}
	unprocessableEntityType struct{}
	tooManyRequestsType     struct{}
)

type (
	errBadRequest          = customError[badRequestType]
	errUnauthorized        = customError[unauthorizedType]
	errForbidden           = customError[forbiddenType]
	errNotFound            = customError[notFoundType]
	errRequestTimeout      = customError[requestTimeoutType]
	errConflict            = customError[conflictType]
	errUnprocessableEntity = customError[unprocessableEntityType]
	errTooManyRequests     = customError[tooManyRequestsType]
)

// NewErrBadRequest はBad Requestエラーを作成します。
func NewErrBadRequest(message string) *errBadRequest {
	return newCustomError[badRequestType](message)
}

// NewErrUnauthorized はUnauthorizedエラーを作成します。
func NewErrUnauthorized(message string) *errUnauthorized {
	return newCustomError[unauthorizedType](message)
}

// NewErrForbidden はForbiddenエラーを作成します。
func NewErrForbidden(message string) *errForbidden {
	return newCustomError[forbiddenType](message)
}

// NewErrNotFound はNot Foundエラーを作成します。
func NewErrNotFound(message string) *errNotFound {
	return newCustomError[notFoundType](message)
}

// NewErrRequestTimeout はRequest Timeoutエラーを作成します。
func NewErrRequestTimeout(message string) *errRequestTimeout {
	return newCustomError[requestTimeoutType](message)
}

// NewErrConflict はConflictエラーを作成します。
func NewErrConflict(message string) *errConflict {
	return newCustomError[conflictType](message)
}

// NewErrUnprocessableEntity はUnprocessable Entityエラーを作成します。
func NewErrUnprocessableEntity(message string) *errUnprocessableEntity {
	return newCustomError[unprocessableEntityType](message)
}

// NewErrTooManyRequests はToo Many Requestsエラーを作成します。
func NewErrTooManyRequests(message string) *errTooManyRequests {
	return newCustomError[tooManyRequestsType](message)
}
