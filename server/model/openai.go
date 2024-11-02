package model

import "github.com/sashabaranov/go-openai"

func GetOpenAiClientFromToken(authToken string) *openai.Client {
	config := openai.DefaultConfig(authToken)
	c := openai.NewClientWithConfig(config)
	return c
}
