package models

import (
	"github.com/init-stuff/init/util"
	"github.com/openai/openai-go"
)

var (
	codeResponseSchema = util.GenerateSchema[Code]()

	CodeResponseSchemaParam = openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("code"),
		Description: openai.F("The code that satisfies the user's request."),
		Schema:      openai.F(codeResponseSchema),
		Strict:      openai.F(true),
	}
)

type Code struct {
	Code string `json:"code" jsonschema_description:"The code that satisfies the user's request. Make sure to provide the program in its entirety, including any imports."`
}
