package openai

import (
	"context"
	"fmt"
	"os"
	"strconv"

	openai "github.com/sashabaranov/go-openai"

	"github.com/metatube-community/metatube-sdk-go/translate"
)

var _ translate.Translator = (*OpenAI)(nil)

type OpenAI struct {
	APIKey string `json:"openai-api-key"`
}

func (oa *OpenAI) Translate(q, source, target string) (result string, err error) {
	openaiConf := openai.DefaultConfig(oa.APIKey)
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
				Content: fmt.Sprintf(`将我发给你的内容翻译成简体中文，直接返回翻译后的结果，不要附加其他信息。
%s`, q),
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

func init() {
	translate.Register(&OpenAI{})
}
