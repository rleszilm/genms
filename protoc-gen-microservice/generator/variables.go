package generator

import (
	"bytes"
	"text/template"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	variablesTemplateSource = `{{ range $label, $value := .variables }}
	{{ $label }} = {{ printf "%s" $value }}
{{ end }}`
)

var (
	variablesTemplate = template.Must(
		template.New("variables").
			Funcs(template.FuncMap{}).
			Parse(variablesTemplateSource))
)

func generateVariables(g generator, defaultVals map[string]interface{}) (*plugin_go.CodeGeneratorResponse_File, error) {
	vals := map[string]interface{}{}
	for k, v := range defaultVals {
		vals[k] = v
	}
	vals["variables"] = g.Variables()

	buf := &bytes.Buffer{}
	if err := variablesTemplate.Execute(buf, vals); err != nil {
		return nil, err
	}

	return &plugin_go.CodeGeneratorResponse_File{
		Name:           strRef(g.FileName()),
		InsertionPoint: strRef("genms-variables"),
		Content:        strRef(buf.String()),
	}, nil
}
