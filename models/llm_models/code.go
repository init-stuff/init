package llm_models

import (
	"github.com/init-stuff/init/models"
	"github.com/openai/openai-go"
)

var (
	codeResponseSchema = models.GenerateSchema[Code]()

	CodeResponseSchemaParam = openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("code"),
		Description: openai.F("The code to run."),
		Schema:      openai.F(codeResponseSchema),
		Strict:      openai.F(true),
	}
)

type Code struct {
	Code string `json:"code" jsonschema_description:"The code that satisfies the user's request. Make sure to provide the program in its entirety, including any imports."`
}
