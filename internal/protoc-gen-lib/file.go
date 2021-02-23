package protocgenlib

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// File adds functionality to the underlying File.
type File struct {
	Outfile *protogen.GeneratedFile
	File    *protogen.File
}

// NewFile returns a new File.
func NewFile(outfile *protogen.GeneratedFile, file *protogen.File) *File {
	return &File{
		Outfile: outfile,
		File:    file,
	}
}

// Write writes to file.
func (f *File) Write(p []byte) (int, error) {
	return f.Outfile.Write(p)
}

// PackageName returns the name of the package.
func (f *File) PackageName() string {
	return string("dal_" + f.File.GoPackageName)
}

// PackagePath returns the path of the package.
func (f *File) PackagePath() string {
	return string(f.File.GoImportPath)
}

// QualifiedPackageName adds the import path to the outfile and returns the auto-generated
// alias used for the package.
func (f *File) QualifiedPackageName(path string) string {
	ident := protogen.GoIdent{GoImportPath: protogen.GoImportPath(path)}
	return strings.Split(f.Outfile.QualifiedGoIdent(ident), ".")[0]
}
