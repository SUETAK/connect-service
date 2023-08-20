package usecase

import (
	"connectrpc.com/connect"
	elizav1 "example/gen/eliza/v1"
	"golang.org/x/net/context"
)

// TODO say関数を別のpackage に移動させる
func (s *ElizaServer) Hello(
	ctx context.Context,
	req *connect.Request[elizav1.HelloRequest],
) (*connect.Response[elizav1.HelloResponse], error) {

	res := connect.NewResponse(&elizav1.HelloResponse{
		Age: 10,
	})

	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
