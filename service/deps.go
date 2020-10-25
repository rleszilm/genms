package service

// Dependencies tracks services based on their dependencies.
type Dependencies struct {
	first string
	path  []string
	svcs  map[string]Service
	deps  map[string]map[string]struct{}
	err   error
	cycle []string
}

// Register adds a service to the dependencies.
func (d *Dependencies) Register(svc Service, deps ...string) {
	if len(d.svcs) == 0 {
		d.first = svc.Name()
	}

	d.svcs[svc.Name()] = svc
	d.deps[svc.Name()] = deriveSetDeps(deps)
	d.sort()
}

// Iterate returns an iterator for the Dependencies
func (d *Dependencies) Iterate() DependencyIterator {
	res := &Iterator{
		Dependencies: d,
	}

	return res
}

// Reverse returns a reverse iterator for the Dependencies
func (d *Dependencies) Reverse() DependencyIterator {
	res := &ReverseIterator{
		Dependencies: d,
	}

	return res
}

// Err returns the last known error about the state. This is updated after
// each call to append.
func (d *Dependencies) Err() error {
	return d.err
}

// Cycle returns the dependency cycle that caused an error if one is present.
func (d *Dependencies) Cycle() []string {
	return d.cycle
}

func (d *Dependencies) sort() {
	next := d.first
	deps := deriveCloneDeps(d.deps)
	path := []string{}
	for len(deps) > 0 {
		tpath, cycle, err := d.traverse(next, deps, []string{}, path)
		if err != nil {
			if err == ErrDependencyCycle {
				d.cycle = cycle
			}
			d.err = err
			return
		}

		for k := range deps {
			next = k
			break
		}

		path = tpath
	}

	d.path = path
	d.err = nil
	d.cycle = nil
}

func (d *Dependencies) traverse(node string, graph map[string]map[string]struct{}, cycle []string, path []string) ([]string, []string, error) {
	edges, ok := graph[node]
	_, exists := d.deps[node]
	if !ok {
		if !exists {
			return nil, nil, ErrMissingDependency
		}
		return path, cycle, nil
	}

	cycle = append(cycle, node)
	if deriveContains(cycle[:len(cycle)-1], node) {
		return nil, cycle, ErrDependencyCycle
	}

	for edge := range edges {
		tpath, tcycle, err := d.traverse(edge, graph, cycle, path)
		if err == ErrDependencyCycle {
			return nil, tcycle, err
		} else if err != nil {
			return nil, nil, err
		}
		path, cycle = tpath, tcycle
	}

	delete(graph, node)
	return append(path, node), cycle[:len(cycle)-1], nil
}

// NewDependencies returns a new Dependencies
func NewDependencies() *Dependencies {
	return &Dependencies{
		svcs: map[string]Service{},
		deps: map[string]map[string]struct{}{},
	}
}
