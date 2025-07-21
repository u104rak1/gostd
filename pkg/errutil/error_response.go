package errutil

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// errorResponse はRFC 9457に準拠したエラーレスポンスの構造体です。
// https://www.rfc-editor.org/rfc/rfc9457.html
type errorResponse struct {
	Type   string    `json:"type"`             // エラーの種類を示すURI
	Title  string    `json:"title"`            // 人間が読めるエラーの要約
	Status int       `json:"status"`           // HTTPステータスコード
	Detail string    `json:"detail"`           // エラーの詳細説明
	Errors *[]string `json:"errors,omitempty"` // バリデーションエラーメッセージの配列（オプション）
}

// エラータイプのベースURI
const errorTypeBaseURI = "https://example.com/errors"

// HandleError はエラーを処理し、RFC 9457に準拠したhttpエラーレスポンスを返します。
func HandleError(ctx *gin.Context, err error) {
	var status int
	var title string
	var detail string
	var errType string

	switch {
	// 400 Bad Request
	case errors.Is(err, &errBadRequest{}):
		status = http.StatusBadRequest
		title = "Bad Request"
		detail = "The request could not be processed due to invalid syntax or parameters."
		errType = fmt.Sprintf("%s/bad-request", errorTypeBaseURI)

	// 401 Unauthorized
	case errors.Is(err, &errUnauthorized{}):
		status = http.StatusUnauthorized
		title = "Unauthorized"
		detail = "Authentication is required to access this resource."
		errType = fmt.Sprintf("%s/unauthorized", errorTypeBaseURI)

	// 403 Forbidden
	case errors.Is(err, &errForbidden{}):
		status = http.StatusForbidden
		title = "Forbidden"
		detail = "You do not have permission to access this resource."
		errType = fmt.Sprintf("%s/forbidden", errorTypeBaseURI)

	// 404 Not Found
	case errors.Is(err, &errNotFound{}):
		status = http.StatusNotFound
		title = "Not Found"
		detail = "The requested resource could not be found."
		errType = fmt.Sprintf("%s/not-found", errorTypeBaseURI)

	// 408 Request Timeout
	case errors.Is(err, &errRequestTimeout{}):
		status = http.StatusRequestTimeout
		title = "Request Timeout"
		detail = "The request took too long to process and timed out."
		errType = fmt.Sprintf("%s/request-timeout", errorTypeBaseURI)

	// 409 Conflict
	case errors.Is(err, &errConflict{}):
		status = http.StatusConflict
		title = "Conflict"
		detail = "The request could not be completed due to a conflict with the current state of the resource."
		errType = fmt.Sprintf("%s/conflict", errorTypeBaseURI)

	// 429 Too Many Requests
	case errors.Is(err, &errTooManyRequests{}):
		status = http.StatusTooManyRequests
		title = "Too Many Requests"
		detail = "You have sent too many requests in a given amount of time."
		errType = fmt.Sprintf("%s/too-many-requests", errorTypeBaseURI)

	// 500 Internal Server Error
	default:
		status = http.StatusInternalServerError
		title = "Internal Server Error"
		detail = "An unexpected error occurred."
		errType = fmt.Sprintf("%s/internal-server-error", errorTypeBaseURI)
	}

	response := errorResponse{
		Type:   errType,
		Title:  title,
		Status: status,
		Detail: detail,
	}

	ctx.JSON(status, response)
}

// HandleValidationError はバリデーションエラーを処理するための専用関数です。
// 複数のバリデーションエラーメッセージを含むレスポンスを生成します。
func HandleValidationError(ctx *gin.Context, validationErrors []string) {
	response := errorResponse{
		Type:   fmt.Sprintf("%s/unprocessable-entity", errorTypeBaseURI),
		Title:  "Unprocessable Entity",
		Status: http.StatusUnprocessableEntity,
		Detail: "The input parameters are invalid. See the errors member for details.",
		Errors: &validationErrors,
	}

	ctx.JSON(http.StatusUnprocessableEntity, response)
}
