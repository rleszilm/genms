package generator

import (
	"path"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/types/descriptorpb"
)

type fileRunner struct {
	file      *descriptorpb.FileDescriptorProto
	opts      *options
	Imports   map[string]struct{}
	Consts    map[string]string
	Variables map[string]string
}

func (r *fileRunner) NewImports(imports []string) []string {
	return deriveFilterStrMap(func(imp string) bool {
		_, ok := r.Imports[imp]
		if !ok {
			r.Imports[imp] = struct{}{}
		}
		return !ok
	}, imports)
}

func (r *fileRunner) Generate() ([]*plugin_go.CodeGeneratorResponse_File, error) {
	defaults := map[string]interface{}{}

	resp := []*plugin_go.CodeGeneratorResponse_File{}
	for _, svc := range r.file.GetService() {
		sr := newServiceRunner(r, svc, r.opts)
		chunks, err := sr.Generate(defaults)
		if err != nil {
			return nil, err
		}

		resp = append(resp, chunks...)
	}

	if len(resp) > 0 {
		f, err := generate(newOutline(r), defaults)
		if err != nil {
			return nil, err
		}

		// the first item of a file cannot have an insertion point
		resp = append(f, resp...)
	}

	return resp, nil
}

func newFileRunner(file *descriptorpb.FileDescriptorProto, opts *options) *fileRunner {
	return &fileRunner{
		file:      file,
		opts:      opts,
		Imports:   map[string]struct{}{},
		Consts:    map[string]string{},
		Variables: map[string]string{},
	}
}

func filename(file *descriptorpb.FileDescriptorProto, opts *options) string {
	name := file.GetName()
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	name += ".genms.go"

	if opts.RelativePaths {
		return name
	}

	goPackage := file.GetOptions().GoPackage
	if goPackage != nil && *goPackage != "" {
		_, name = path.Split(name)
		return path.Join(*goPackage, name)
	}
	return name
}
