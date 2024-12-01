package llm_models

import (
	"github.com/init-stuff/init/models"
	"github.com/openai/openai-go"
)

var (
	modifyCodeResponseSchema = models.GenerateSchema[ModifyCode]()

	ModifyCodeResponseSchemaParam = openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("code"),
		Description: openai.F("The modified code that satisfies the user's request."),
		Schema:      openai.F(modifyCodeResponseSchema),
		Strict:      openai.F(true),
	}
)

type ModifyCode struct {
	Code string `json:"code" jsonschema_description:"The modified code that satisfies the user's request. Make sure to provide the program in its entirety, including any imports."`
}
