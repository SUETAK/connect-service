package usecase

import "example/service/openai"

type ElizaServer struct {
	openaiClient openai.AIClient
}

func NewElizaServer(openaiKey string) *ElizaServer {
	return &ElizaServer{
		openaiClient: openai.NewClient(openaiKey),
	}
}
