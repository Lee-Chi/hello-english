package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

var openaiClient *openai.Client

func Build(openaiToken string) error {
	openaiClient = openai.NewClient(openaiToken)

	return nil
}

const (
	ChatMessageRoleSystem    string = "system"
	ChatMessageRoleUser      string = "user"
	ChatMessageRoleAssistant string = "assistant"
	ChatMessageRoleFunction  string = "function"
	ChatMessageRoleTool      string = "tool"
)

type ChatCompletionMessage openai.ChatCompletionMessage

type ChatCompletionMessages []ChatCompletionMessage

func (m ChatCompletionMessages) ToOpenAI() []openai.ChatCompletionMessage {
	msgs := make([]openai.ChatCompletionMessage, len(m))
	for i, v := range m {
		msgs[i] = openai.ChatCompletionMessage(v)
	}

	return msgs
}
func Chat(ctx context.Context, messages ChatCompletionMessages) (string, error) {
	resp, err := openaiClient.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo0125,
			Messages: messages.ToOpenAI(),
		},
	)

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("ChatGPT is busy")
	}

	return resp.Choices[0].Message.Content, nil
}
