package generator

import (
	"bytes"
	"text/template"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	importsTemplateSource = `{{ range .imports -}}
{{ . }}
{{ end }}`
)

var (
	importsTemplate = template.Must(
		template.New("imports").
			Funcs(template.FuncMap{}).
			Parse(importsTemplateSource))
)

func generateImports(g generator, defaultVals map[string]interface{}) (*plugin_go.CodeGeneratorResponse_File, error) {
	vals := g.Replacements(defaultVals)
	vals["imports"] = g.FileRunner().NewImports(g.Imports())

	buf := &bytes.Buffer{}
	if err := importsTemplate.Execute(buf, vals); err != nil {
		return nil, err
	}

	return &plugin_go.CodeGeneratorResponse_File{
		Name:           strRef(g.FileName()),
		InsertionPoint: strRef("genms-imports"),
		Content:        strRef(buf.String()),
	}, nil
}
