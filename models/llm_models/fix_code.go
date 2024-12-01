package llm_models

import (
	"github.com/init-stuff/init/models"
	"github.com/openai/openai-go"
)

var (
	fixCodeResponseSchema = models.GenerateSchema[FixCode]()

	FixCodeResponseSchemaParam = openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("code"),
		Description: openai.F("The code that fixes the errors presented."),
		Schema:      openai.F(fixCodeResponseSchema),
		Strict:      openai.F(true),
	}
)

type FixCode struct {
	Code string `json:"code" jsonschema_description:"The code that fixes the errors presented. Make sure to provide the program in its entirety, including any imports."`
}
