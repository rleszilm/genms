package generator

import (
	"path"
	"strings"

	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// File adds functionality to the underlying File.
type File struct {
	*protocgenlib.File
}

// NewFile returns a new File.
func NewFile(outfile *protogen.GeneratedFile, file *protogen.File) *File {
	return AsFile(protocgenlib.NewFile(outfile, file))
}

// AsFile wraps a File.
func AsFile(file *protocgenlib.File) *File {
	return &File{
		File: file,
	}
}

// Generator returns the underlying generator.File
func (f *File) Generator() *File {
	return f
}

// DalPackageName returns the name of the dal  package.
func (f *File) DalPackageName() string {
	return "dal_" + f.File.PackageName()
}

// DalPackagePath returns the path of the package.
func (f *File) DalPackagePath() string {
	toks := append([]string{strings.ReplaceAll(f.Proto().GoImportPath.String(), "\"", "")}, "dal")
	return path.Join(toks...)
}
