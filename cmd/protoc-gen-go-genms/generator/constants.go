package generator

import (
	"bytes"
	"text/template"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	constantsTemplateSource = `{{ range $label, $value := .constants }}
	{{ $label }} = {{ printf "%q" $value }}
{{ end }}`
)

var (
	constantsTemplate = template.Must(
		template.New("constants").
			Funcs(template.FuncMap{}).
			Parse(constantsTemplateSource))
)

func generateConstants(g generator, defaultVals map[string]interface{}) (*plugin_go.CodeGeneratorResponse_File, error) {
	vals := map[string]interface{}{}
	for k, v := range defaultVals {
		vals[k] = v
	}
	vals["constants"] = g.Constants()

	buf := &bytes.Buffer{}
	if err := constantsTemplate.Execute(buf, vals); err != nil {
		return nil, err
	}

	return &plugin_go.CodeGeneratorResponse_File{
		Name:           strRef(g.FileName()),
		InsertionPoint: strRef("genms-constants"),
		Content:        strRef(buf.String()),
	}, nil
}
