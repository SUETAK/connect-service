package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Payload struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type Client struct {
	Key string
}

func NewClient(key string) AIClient {
	return &Client{
		Key: key,
	}
}

type AIClient interface {
	GetCompletion(prompt string, maxTokens int) (GptResponse, error)
}

func (c *Client) GetCompletion(prompt string, maxTokens int) (GptResponse, error) {
	url := "https://api.openai.com/v1/chat/completions"

	// system の立場を設定
	systemMessage := ChatMessage{
		Role:    "system",
		Content: "you are a good english speaker.",
	}

	// 対象にしたい文章
	assistantMessage := ChatMessage{
		Role:    "assistant",
		Content: prompt,
	}

	// system に何をさせたいか
	userMessage := ChatMessage{
		Role:    "user",
		Content: "Please convert the following sentence into one that uses a relative pronoun",
	}

	payloadData := &Payload{
		Model:    "gpt-3.5-turbo",
		Messages: []ChatMessage{systemMessage, userMessage, assistantMessage},
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		return GptResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return GptResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GptResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return GptResponse{}, fmt.Errorf("OpenAI API returned non-200 status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GptResponse{}, err
	}

	var gptResponse GptResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		fmt.Println(err)
	}

	return gptResponse, nil
}

type GptResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
