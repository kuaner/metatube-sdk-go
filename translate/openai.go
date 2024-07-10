package translate

import (
	"context"
	"fmt"
	"os"
	"strconv"

	openai "github.com/sashabaranov/go-openai"
)

func OpenaiTranslate(q, source, target, key string) (result string, err error) {
	openaiConf := openai.DefaultConfig(key)
	if withURL := os.Getenv("OPENAI_BASE_URL"); withURL != "" {
		openaiConf.BaseURL = withURL
	}
	msg := openai.ChatCompletionRequest{
		Model:       openai.GPT4o,
		Temperature: 1.0,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a professional, authentic machine translation engine.",
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: fmt.Sprintf(`Translate the following source text to Simplified Chinese Language, Output translation directly without any additional text.
Source Text: %s`, q),
			},
		},
	}
	if withModel := os.Getenv("OPENAI_MODEL"); withModel != "" {
		msg.Model = withModel
	}
	if withTemperature := os.Getenv("OPENAI_TEMPERATURE"); withTemperature != "" {
		// convert string to float32
		temperature, err := strconv.ParseFloat(withTemperature, 32)
		if err == nil {
			msg.Temperature = float32(temperature)
		}
	}
	resp, err := openai.NewClientWithConfig(openaiConf).CreateChatCompletion(
		context.Background(),
		msg,
	)
	if err != nil {
		return "error", err
	}
	return resp.Choices[0].Message.Content, nil
}
