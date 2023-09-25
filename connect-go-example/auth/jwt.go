package auth

import (
	_ "embed"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// os.ReadFile を使うとファイルも適切なpath に置く必要があるため、embed を使って実行バイナリにファイルを埋め込む

//go:embed cert/secret.pem
var rawPrivateKey []byte

//go:embed cert/public.pem
var rawPublicKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	//Clocker               clock.Clocker

}
