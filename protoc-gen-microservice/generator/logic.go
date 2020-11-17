package generator

import (
	"bytes"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func generateLogic(g generator, defaultVals map[string]interface{}) (*plugin_go.CodeGeneratorResponse_File, error) {
	vals := map[string]interface{}{}
	for k, v := range defaultVals {
		vals[k] = v
	}

	buf := &bytes.Buffer{}
	if err := g.Logic().Execute(buf, vals); err != nil {
		return nil, err
	}

	return &plugin_go.CodeGeneratorResponse_File{
		Name:           strRef(g.FileName()),
		InsertionPoint: strRef("genms-logic"),
		Content:        strRef(buf.String()),
	}, nil
}
