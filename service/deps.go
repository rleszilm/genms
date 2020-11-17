package service

// Deps manages a services dependencies.
type Deps Services

// Dependencies returns the services dependencies.
func (d *Deps) Dependencies() Services {
	if d == nil {
		return nil
	}

	return Services(*d)
}

// WithDependencies adds dependencies to the service.
func (d *Deps) WithDependencies(svcs ...Service) {
	if d == nil {
		return
	}

	*d = append(*d, svcs...)
}
