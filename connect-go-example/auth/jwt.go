package auth

import _ "embed"

// os.ReadFile を使うとファイルも適切なpath に置く必要があるため、embed を使って実行バイナリにファイルを埋め込む

//go:embed cert/secret.pem
var rawPrivateKey []byte

//go:embed cert/public.pem
var rawPublicKey []byte
