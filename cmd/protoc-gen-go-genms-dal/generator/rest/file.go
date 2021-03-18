package rest

import (
	"path"
	"strings"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

// File adds functionality to the underlying File.
type File struct {
	*generator.File
	name string
}

// NewFile returns a new File.
func NewFile(outfile *protogen.GeneratedFile, file *protogen.File) *File {
	return AsFile(generator.NewFile(outfile, file))
}

// AsFile wraps a File.
func AsFile(file *generator.File) *File {
	return &File{
		File: file,
	}
}

// RestDalPackageName returns the name of the dal  package.
func (f *File) RestDalPackageName() string {
	return "dal_" + f.File.PackageName()
}

// RestDalPackagePath returns the path of the package.
func (f *File) RestDalPackagePath() string {
	toks := append([]string{strings.ReplaceAll(f.Proto().GoImportPath.String(), "\"", "")}, "dal")
	return path.Join(toks...)
}
