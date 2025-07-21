/*
errutil パッケージは、HTTP API向けのエラーハンドリングユーティリティを提供します。
主な機能として、用途別のカスタムエラー生成、エラーラップ、Ginハンドラ向けのエラーレスポンス生成などがあります。
エラーレスポンスのフォーマットは、RFC 9457（Problem Details for HTTP APIs）に準拠します。（https://datatracker.ietf.org/doc/html/rfc9457）
これにより、API利用者がエラーの種類や詳細を機械的に判別しやすくなります。
*/
package errutil

import "fmt"

// Wrap はエラーメッセージをラップして新しいエラーを返します。
func Wrap(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}
