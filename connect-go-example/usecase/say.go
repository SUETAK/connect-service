package usecase

import (
	elizav1 "example/gen/eliza/v1"
	"example/service/openai"
	"fmt"
	"github.com/bufbuild/connect-go"
	"golang.org/x/net/context"
	"log"
)

// TODO say関数を別のpackage に移動させる
func (s *ElizaServer) Say(
	ctx context.Context,
	req *connect.Request[elizav1.SayRequest],
) (*connect.Response[elizav1.SayResponse], error) {
	log.Println("Request headers: ", req.Header())

	// TODO openaiClient に移動させる。interface にする
	openaiRes, err := s.openaiClient.GetCompletion(req.Msg.Sentence, 60)
	if err != nil {
		return nil, err
	}
	log.Println(openaiRes)
	var messages []openai.ChatMessage
	for _, choice := range openaiRes.Choices {
		messages = append(messages, choice.Message)
	}

	var joinedMessages string
	for _, message := range messages {
		joinedMessages += message.Content
	}

	res := connect.NewResponse(&elizav1.SayResponse{
		Sentence: fmt.Sprintf("Hello, %s! Here is a completion: %s", req.Msg.Sentence, joinedMessages),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
