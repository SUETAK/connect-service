package di

import (
	"connectrpc.com/connect"
	"example/interfaces/proto/eliza/v1/elizav1connect"
	"example/interfaces/proto/sign/v1/signv1connect"
	"example/repository"
	"example/usecase"
	"github.com/uptrace/bun"
	"net/http"
)

func InitSign(mux *http.ServeMux, db *bun.DB, authInterceptor connect.Option) {
	repo := repository.NewUserRepository(db)
	user := usecase.NewSignServer(repo)
	signPath, signHandler := signv1connect.NewSignServiceHandler(user, authInterceptor)
	mux.Handle(signPath, signHandler)
}

func InitEliza(mux *http.ServeMux, openaiKey string, db *bun.DB, authInterceptor connect.Option) {
	repo := repository.NewUserRepository(db)
	eliza := usecase.NewElizaServer(openaiKey, repo)
	path, handler := elizav1connect.NewElizaServiceHandler(eliza, authInterceptor)
	mux.Handle(path, handler)
}
