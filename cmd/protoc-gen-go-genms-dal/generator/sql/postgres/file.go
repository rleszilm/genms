package postgres

import (
	"path"
	"strings"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// File adds functionality to the underlying File.
type File struct {
	*generator.File
}

// NewFile returns a new File.
func NewFile(outfile *protogen.GeneratedFile, file *protogen.File) *File {
	return AsFile(protocgenlib.NewFile(outfile, file))
}

// AsFile wraps a File.
func AsFile(file *protocgenlib.File) *File {
	return &File{
		File: generator.AsFile(file),
	}
}

// PostgresPackageName returns the name of the dal  package.
func (f *File) PostgresPackageName() string {
	return "postgres_" + f.File.DalPackageName()
}

// PostgresPackagePath returns the path of the package.
func (f *File) PostgresPackagePath() string {
	toks := append([]string{strings.ReplaceAll(f.Proto().GoImportPath.String(), "\"", "")}, "dal", "postgres")
	return path.Join(toks...)
}
