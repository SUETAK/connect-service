package contextkey

// コンテキスト値に対応するキー文字列
type ctxKey string

const (
	// ユーザーを識別するID。認証後にコンテキストにセットされる
	CtxKeyUserID ctxKey = "ctx-user-id"
)
