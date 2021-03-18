package protocgenlib

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// File adds functionality to the underlying File.
type File struct {
	outfile *protogen.GeneratedFile
	file    *protogen.File
}

// NewFile returns a new File.
func NewFile(outfile *protogen.GeneratedFile, file *protogen.File) *File {
	return &File{
		outfile: outfile,
		file:    file,
	}
}

// Proto returns the base protogen object.
func (f *File) Proto() *protogen.File {
	return f.file
}

// Outfile returns the file to which this field would be written.
func (f *File) Outfile() *protogen.GeneratedFile {
	return f.outfile
}

// Write writes to file.
func (f *File) Write(p []byte) (int, error) {
	return f.outfile.Write(p)
}

// PackageName returns the name of the package.
func (f *File) PackageName() string {
	return string(f.file.GoPackageName)
}

// PackagePath returns the path of the package.
func (f *File) PackagePath() string {
	return string(f.file.GoImportPath)
}

// QualifiedKind returns the qualified ident.
func (f *File) QualifiedKind(i protogen.GoIdent) string {
	return f.outfile.QualifiedGoIdent(i)
}

// QualifiedPackageName adds the import path to the outfile and returns the auto-generated
// alias used for the package.
func (f *File) QualifiedPackageName(path string) string {
	ident := protogen.GoIdent{GoImportPath: protogen.GoImportPath(path)}
	return strings.Split(f.outfile.QualifiedGoIdent(ident), ".")[0]
}
