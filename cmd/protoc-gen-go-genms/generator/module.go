package generator

import (
	"fmt"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/types/descriptorpb"
)

type moduleRunner struct {
	Consts    map[string]string
	Variables map[string]string
}

func (r *moduleRunner) Generate(file *descriptorpb.FileDescriptorProto, opts *options) ([]*plugin_go.CodeGeneratorResponse_File, error) {

	fr := newFileRunner(file, opts)
	res, err := fr.Generate()
	if err != nil {
		return nil, err
	}

	for con := range fr.Consts {
		if _, ok := r.Consts[con]; ok {
			return nil, fmt.Errorf("constant (%v) already defined", con)
		}
	}

	for vari := range fr.Variables {
		if _, ok := r.Variables[vari]; ok {
			return nil, fmt.Errorf("variable (%v) already defined", vari)
		}
	}

	return res, nil
}

func newModuleRunner() *moduleRunner {
	return &moduleRunner{
		Consts:    map[string]string{},
		Variables: map[string]string{},
	}
}
