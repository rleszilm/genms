package rest

import (
	"fmt"
	"path"
	"strings"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

// File adds functionality to the underlying File.
type File struct {
	*generator.File
	name string
}

// NewFile returns a new File.
func NewFile(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message) *File {
	base := path.Base(file.GeneratedFilenamePrefix)
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/rest/%s.genms.dal.%s.go", base, strings.ToLower(msg.GoIdent.GoName)))
	outfile := plugin.NewGeneratedFile(filename, ".")

	return &File{
		File: generator.NewFile(outfile, file),
		name: filename,
	}
}

// Outfile returns the underlying generated file of the file.
func (f *File) Outfile() *protogen.GeneratedFile {
	return f.File.Outfile
}

// Name returns the name of the file.
func (f *File) Name() string {
	return f.name
}

// BaseName returns the base name of the file.
func (f *File) BaseName() string {
	return path.Base(f.name)
}

// DirName returns the directory name of the file.
func (f *File) DirName() string {
	return path.Dir(f.name)
}
