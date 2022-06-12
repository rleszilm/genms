package service

// Services defines a set of services
type Services []Service

// Sort sorts the services in the order they can be successfull started.
func (s *Services) Sort() ([]Service, error) {
	if s == nil {
		return nil, nil
	}

	var err error
	var path []Service
	var cycle []Service
	sorted := []Service{}
	visited := map[Service]struct{}{}
	// go across each service to see if it's a start point
	for _, svc := range *s {
		// if the service is marked as visited it is not a start point
		if _, ok := visited[svc]; ok {
			continue
		}

		visited, cycle, path, err = s.traverse(svc, visited, []Service{}, []Service{})
		if err != nil {
			return cycle, err
		}

		sorted = append(sorted, path...)
	}

	svcs := Services(sorted)
	*s = svcs
	return nil, nil
}

func (s *Services) traverse(node Service, visited map[Service]struct{}, cycle []Service, path []Service) (map[Service]struct{}, []Service, []Service, error) {
	// check for cycle
	cycle = append(cycle, node)
	for _, svc := range cycle[:len(cycle)-1] {
		if node == svc {
			return nil, cycle, nil, ErrDependencyCycle
		}
	}

	// dependencies must be satisfied before this node can be added
	for _, node := range node.dependencies() {
		// If the node is visited skip it.
		if _, ok := visited[node]; !ok {
			tVisited, tCycle, tPath, err := s.traverse(node, visited, cycle, []Service{})
			if err == ErrDependencyCycle {
				return nil, tCycle, nil, err
			} else if err != nil {
				return nil, nil, nil, err
			}
			visited, cycle, path = tVisited, tCycle, append(path, tPath...)
		}
	}

	visited[node] = struct{}{}
	return visited, cycle, append(path, node), nil
}
