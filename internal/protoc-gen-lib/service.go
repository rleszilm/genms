package protocgenlib

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Service adds functionality to the underlying Service.
type Service struct {
	file    *File
	service *protogen.Service
}

// NewService returns a new Service.
func NewService(file *File, service *protogen.Service) *Service {
	return &Service{
		file:    file,
		service: service,
	}
}

// Proto returns the base protogen object.
func (s *Service) Proto() *protogen.Service {
	return s.service
}

// File returns the base File object.
func (s *Service) File() *File {
	return s.file
}

// Outfile returns the file to which this field would be written.
func (s *Service) Outfile() *protogen.GeneratedFile {
	return s.file.Outfile()
}

// Name returns the name.
func (s *Service) Name() string {
	return s.service.GoName
}

// Kind returns the fields go type.
func (s *Service) Kind() string {
	return s.service.GoName
}

// QualifiedKind returns the fully qualified type.
func (s *Service) QualifiedKind() string {
	return s.service.GoName
}
