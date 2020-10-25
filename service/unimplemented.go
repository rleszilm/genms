package service

import "context"

// Unimplemented is a Service that has no needed Intialize or Shutdown logic.
type Unimplemented struct{}

// Initialize implements service.Initialize
func (n *Unimplemented) Initialize(context.Context) error {
	return nil
}

// Shutdown implements service.Shutdown
func (n *Unimplemented) Shutdown(context.Context) error {
	return nil
}
