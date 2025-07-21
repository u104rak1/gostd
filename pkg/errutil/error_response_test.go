package errutil_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/u104rak1/gostd/pkg/errutil"
)

// テスト用のレスポンス解析構造体
type errorResponseTest struct {
	Type   string    `json:"type"`
	Title  string    `json:"title"`
	Status int       `json:"status"`
	Detail string    `json:"detail"`
	Errors *[]string `json:"errors,omitempty"`
}

func TestHandleError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		caseName   string
		err        error
		wantStatus int
		wantType   string
		wantTitle  string
		wantDetail string
	}{
		{
			caseName:   "errBadRequestを受け取った時、400 Bad Requestを返す",
			err:        errutil.NewErrBadRequest("bad request error"),
			wantStatus: http.StatusBadRequest,
			wantType:   "https://example.com/errors/bad-request",
			wantTitle:  "Bad Request",
			wantDetail: "The request could not be processed due to invalid syntax or parameters.",
		},
		{
			caseName:   "errUnauthorizedを受け取った時、401 Unauthorizedを返す",
			err:        errutil.NewErrUnauthorized("unauthorized error"),
			wantStatus: http.StatusUnauthorized,
			wantType:   "https://example.com/errors/unauthorized",
			wantTitle:  "Unauthorized",
			wantDetail: "Authentication is required to access this resource.",
		},
		{
			caseName:   "errForbiddenを受け取った時、403 Forbiddenを返す",
			err:        errutil.NewErrForbidden("forbidden error"),
			wantStatus: http.StatusForbidden,
			wantType:   "https://example.com/errors/forbidden",
			wantTitle:  "Forbidden",
			wantDetail: "You do not have permission to access this resource.",
		},
		{
			caseName:   "errNotFoundを受け取った時、404 Not Foundを返す",
			err:        errutil.NewErrNotFound("not found error"),
			wantStatus: http.StatusNotFound,
			wantType:   "https://example.com/errors/not-found",
			wantTitle:  "Not Found",
			wantDetail: "The requested resource could not be found.",
		},
		{
			caseName:   "errRequestTimeoutを受け取った時、408 Request Timeoutを返す",
			err:        errutil.NewErrRequestTimeout("timeout error"),
			wantStatus: http.StatusRequestTimeout,
			wantType:   "https://example.com/errors/request-timeout",
			wantTitle:  "Request Timeout",
			wantDetail: "The request took too long to process and timed out.",
		},
		{
			caseName:   "errConflictを受け取った時、409 Conflictを返す",
			err:        errutil.NewErrConflict("conflict error"),
			wantStatus: http.StatusConflict,
			wantType:   "https://example.com/errors/conflict",
			wantTitle:  "Conflict",
			wantDetail: "The request could not be completed due to a conflict with the current state of the resource.",
		},
		{
			caseName:   "errTooManyRequestsを受け取った時、429 Too Many Requestsを返す",
			err:        errutil.NewErrTooManyRequests("too many requests error"),
			wantStatus: http.StatusTooManyRequests,
			wantType:   "https://example.com/errors/too-many-requests",
			wantTitle:  "Too Many Requests",
			wantDetail: "You have sent too many requests in a given amount of time.",
		},
		{
			caseName:   "その他のエラーを受け取った時、500 Internal Server Errorを返す",
			err:        errors.New("internal error"),
			wantStatus: http.StatusInternalServerError,
			wantType:   "https://example.com/errors/internal-server-error",
			wantTitle:  "Internal Server Error",
			wantDetail: "An unexpected error occurred.",
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.caseName, func(t *testing.T) {
			// Arrange
			t.Parallel()
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest(http.MethodGet, "/test-path", nil)

			// Act
			errutil.HandleError(c, tt.err)

			// Assert
			assert.Equal(t, tt.wantStatus, w.Code)

			var resp errorResponseTest
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantType, resp.Type)
			assert.Equal(t, tt.wantTitle, resp.Title)
			assert.Equal(t, tt.wantStatus, resp.Status)
			assert.Equal(t, tt.wantDetail, resp.Detail)
		})
	}
}

func TestHandleValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("バリデーションエラーメッセージを複数含むレスポンスを返す", func(t *testing.T) {
		// Arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/test", nil)

		validationErrors := []string{
			"name is required",
			"email format is invalid",
			"age must be greater than or equal to 0",
		}

		// Act
		errutil.HandleValidationError(c, validationErrors)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

		var resp errorResponseTest
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, "https://example.com/errors/unprocessable-entity", resp.Type)
		assert.Equal(t, "Unprocessable Entity", resp.Title)
		assert.Equal(t, http.StatusUnprocessableEntity, resp.Status)
		assert.Equal(t, "The input parameters are invalid. See the errors member for details.", resp.Detail)
		assert.NotNil(t, resp.Errors)
		assert.Len(t, *resp.Errors, 3)
		assert.Contains(t, *resp.Errors, "name is required")
		assert.Contains(t, *resp.Errors, "email format is invalid")
		assert.Contains(t, *resp.Errors, "age must be greater than or equal to 0")
	})
}
