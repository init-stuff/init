package models

import (
	"github.com/init-stuff/init/util"
	"github.com/openai/openai-go"
)

var (
	packagesResponseSchema = util.GenerateSchema[Packages]()

	PackagesResponseSchemaParam = openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("packages"),
		Description: openai.F("A list of python packages required to run the code."),
		Schema:      openai.F(packagesResponseSchema),
		Strict:      openai.F(true),
	}
)

type Packages = []Package

type Package struct {
	Name    string `json:"name" jsonschema_description:"The name of the package."`
	Version string `json:"version" jsonschema_description:"The version of the package. Always set this to \"latest\" unless a specific version is required."`
}
