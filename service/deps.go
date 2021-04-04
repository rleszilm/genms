package service

// Dependencies manages a services dependencies.
type Dependencies Services

// Dependants returns the services dependencies.
func (d *Dependencies) Dependants() Services {
	if d == nil {
		return nil
	}

	return Services(*d)
}

// WithDependencies adds dependencies to the service.
func (d *Dependencies) WithDependencies(svcs ...Service) {
	if d == nil {
		return
	}

	*d = append(*d, svcs...)
}
