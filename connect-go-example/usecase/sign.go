package usecase

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	v1 "example/interfaces/proto/sign/v1"
	"example/repository"
)

type SignServer struct {
	userRepository repository.UserRepository
}

func NewSignServer(userRepository repository.UserRepository) *SignServer {
	return &SignServer{
		userRepository: userRepository,
	}
}

func (s *SignServer) SignIn(ctx context.Context, request *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error) {
	// TODO user テーブルに接続して、ユーザーを検索する
	msg := request.Msg
	user, err := s.userRepository.FindUserByID(ctx, msg.Username)
	if err != nil {
		return nil, err
	}
	// TODO ユーザーが存在しない場合は、エラーを返す
	if user == nil {
		return nil, errors.New("user not found")
	}
	// TODO ユーザーが存在する場合は、JWTを生成して返す
	res := connect.NewResponse(&v1.SignInResponse{
		Token: "123",
	})
	return res, nil
}

func (s *SignServer) SignUp(ctx context.Context, request *connect.Request[v1.SignUpRequest]) (*connect.Response[v1.SignUpResponse], error) {
	msg := request.Msg
	_, err := s.userRepository.CreateUser(ctx, msg.Username)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&v1.SignUpResponse{
		Token: "123",
	})
	return res, nil
}
