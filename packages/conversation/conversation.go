package conversation

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/init-stuff/init/models"
	"github.com/openai/openai-go"
)

var (
	client = openai.NewClient() // API key defaults to env var OPENAI_API_KEY
)

type Conversation struct {
	ID uuid.UUID `json:"id"`
	Params openai.ChatCompletionNewParams `json:"params"`
}

func NewConversation() (*Conversation, error) {
	params := openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(""),
		}),
		Model:    openai.F(openai.ChatModelGPT4o),
	}

	_, err := client.Chat.Completions.New(context.TODO(), params)
	if err != nil {
		return nil, err
	}

	return &Conversation{
		ID: uuid.New(),
		Params: params,
	}, nil
}

func (c *Conversation) PackagesPrompt(message string) (*models.Packages, error) {
	c.Params.Messages.Value = append(c.Params.Messages.Value, openai.UserMessage(
		fmt.Sprintf("Provide a list of python packages required to create a program that satisfies the following project description:\n\n%s", message),
	))

	c.Params.ResponseFormat = openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
		openai.ResponseFormatJSONSchemaParam{
			Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
			JSONSchema: openai.F(models.PackagesResponseSchemaParam),
		},
	)

	res, err := client.Chat.Completions.New(context.TODO(), c.Params)
	if err != nil {
		return nil, err
	}

	var packages models.Packages
	err = json.Unmarshal([]byte(res.Choices[0].Message.Content), &packages)
	if err != nil {
		return nil, err
	}

	c.Params.ResponseFormat = openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](nil)

	return &packages, nil
}
