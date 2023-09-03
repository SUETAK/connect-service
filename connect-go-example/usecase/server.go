package usecase

import (
	"example/repository"
	"example/service/openai"
)

type ElizaServer struct {
	openaiClient   openai.AIClient
	userRepository repository.UserRepository
}

func NewElizaServer(openaiKey string, userRepository repository.UserRepository) *ElizaServer {
	return &ElizaServer{
		openaiClient:   openai.NewClient(openaiKey),
		userRepository: userRepository,
	}
}
