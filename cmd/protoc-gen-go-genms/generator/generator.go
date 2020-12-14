package generator

import (
	"text/template"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type generator interface {
	FileRunner() *fileRunner
	FileName() string
	Replacements(map[string]interface{}) map[string]interface{}
	Imports() []string
	Constants() map[string]interface{}
	Variables() map[string]string
	Logic() *template.Template
	Outline() *template.Template
}

func generate(g generator, defaultVals map[string]interface{}) ([]*plugin_go.CodeGeneratorResponse_File, error) {
	vals := g.Replacements(defaultVals)
	resp := []*plugin_go.CodeGeneratorResponse_File{}

	if ol := g.Outline(); ol != nil {
		outline, err := generateOutline(g, vals)
		if err != nil {
			return nil, err
		}
		resp = append(resp, outline)
	}

	if len(g.Imports()) > 0 {
		imps, err := generateImports(g, vals)
		if err != nil {
			return nil, err
		}

		resp = append(resp, imps)
	}

	if len(g.Constants()) > 0 {
		constants, err := generateConstants(g, vals)
		if err != nil {
			return nil, err
		}

		resp = append(resp, constants)
	}

	if len(g.Variables()) > 0 {
		variables, err := generateVariables(g, vals)
		if err != nil {
			return nil, err
		}

		resp = append(resp, variables)
	}

	if log := g.Logic(); log != nil {
		logic, err := generateLogic(g, vals)
		if err != nil {
			return nil, err
		}
		resp = append(resp, logic)
	}

	return resp, nil
}
