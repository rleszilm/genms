package http_service

import "errors"

// StatusError is an
type StatusError struct {
	error
	statusCode int
}

func NewStatusError(statusCode int, err error) *StatusError {
	return &StatusError{
		error:      err,
		statusCode: statusCode,
	}
}

func (s *StatusError) StatusCode() int {
	return s.statusCode
}

func (s *StatusError) Is(target error) bool {
	return errors.Is(s.error, target)
}
