package usecase

import (
	"connectrpc.com/connect"
	"context"
	v1 "example/interfaces/proto/login/v1"
)

type LoginServer struct {
}

func NewLoginServer() *LoginServer {
	return &LoginServer{}
}

func (s *LoginServer) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	// TODO user テーブルに接続して、ユーザーを検索する

	// TODO ユーザーが存在しない場合は、エラーを返す
	// TODO ユーザーが存在する場合は、JWTを生成して返す

	res := connect.NewResponse(&v1.LoginResponse{
		Token: "123",
	})
	return res, nil
}
