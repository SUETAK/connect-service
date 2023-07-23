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

func NewClient(key string) *Client {
	return &Client{
		Key: key,
	}
}

func (c *Client) GetCompletion(prompt string, maxTokens int) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	// User message
	userMessage := ChatMessage{
		Role:    "system",
		Content: "You are a helpful assistant.",
	}

	// Assistant message
	assistantMessage := ChatMessage{
		Role:    "user",
		Content: prompt,
	}

	payloadData := &Payload{
		Model:    "gpt-3.5-turbo",
		Messages: []ChatMessage{userMessage, assistantMessage},
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("OpenAI API returned non-200 status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
