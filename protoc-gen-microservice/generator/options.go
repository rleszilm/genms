package generator

import (
	"fmt"
	"strings"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type options struct {
	RelativePaths bool
}

func newOptions(req *plugin_go.CodeGeneratorRequest) (*options, error) {
	opts := &options{}

	tokens := strings.Split(req.GetParameter(), ",")
	for _, token := range tokens {
		kv := strings.Split(token, "=")

		switch kv[0] {
		case "paths":
			if len(kv) != 2 {
				return nil, fmt.Errorf("%s requires an argument", kv[0])
			}
			opts.RelativePaths = kv[1] == "source_relative"
		default:
			return nil, fmt.Errorf("%s is an invalid option", kv[0])
		}
	}

	return opts, nil
}
