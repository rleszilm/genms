package service

// Deps manages a services dependencies.
type Deps []Service

// Dependencies returns the services dependencies.
func (d *Deps) Dependencies() []Service {
	if d == nil {
		return nil
	}

	return []Service(*d)
}

// WithDependencies adds dependencies to the service.
func (d *Deps) WithDependencies(svcs ...Service) {
	if d == nil {
		return
	}

	*d = append(*d, svcs...)
}
