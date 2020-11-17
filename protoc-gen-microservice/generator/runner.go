package generator

import (
	"errors"
	"path"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Runner generates files for the requests.
type Runner map[string]*moduleRunner

// Generate generates files.
func (r *Runner) Generate(req *plugin_go.CodeGeneratorRequest) ([]*plugin_go.CodeGeneratorResponse_File, error) {
	if r == nil {
		return nil, errors.New("cannot generate nil runner")
	}

	opts, err := newOptions(req)
	if err != nil {
		return nil, err
	}

	resp := []*plugin_go.CodeGeneratorResponse_File{}
	for _, file := range req.GetProtoFile() {
		module := path.Dir(file.GetName()) + ":" + file.GetPackage()
		if _, ok := (*r)[module]; !ok {
			(*r)[module] = &moduleRunner{}
		}
		chunks, err := (*r)[module].Generate(file, opts)
		if err != nil {
			return nil, err
		}
		resp = append(resp, chunks...)
	}

	return resp, nil
}

// NewRunner returns a new Runner
func NewRunner() *Runner {
	return &Runner{}
}
