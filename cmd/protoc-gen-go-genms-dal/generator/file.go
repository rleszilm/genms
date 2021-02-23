package generator

import (
	"path"
	"strings"

	protocgenlib "github.com/rleszilm/gen_microservice/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// File adds functionality to the underlying File.
type File struct {
	protocgenlib.File
}

// NewFile returns a new File.
func NewFile(outfile *protogen.GeneratedFile, file *protogen.File) *File {
	return &File{
		File: protocgenlib.File{
			Outfile: outfile,
			File:    file,
		},
	}
}

// DalPackageName returns the name of the dal  package.
func (f *File) DalPackageName() string {
	return string(f.File.File.GoPackageName)
}

// DalPackagePath returns the path of the package.
func (f *File) DalPackagePath() string {
	toks := append([]string{strings.ReplaceAll(f.File.File.GoImportPath.String(), "\"", "")}, "dal")
	return path.Join(toks...)
}
