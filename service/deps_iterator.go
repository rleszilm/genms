package service

// DependencyIterator is the interface that defines a Dependency Iterator.
type DependencyIterator interface {
	Next() <-chan Service
	Err() error
}

// Iterator iterates over the sorted elements.
type Iterator struct {
	*Dependencies
}

// Next implements DependencyIterator.Next
func (i *Iterator) Next() <-chan Service {
	next := make(chan Service)
	go func() {
		for _, svc := range i.path {
			next <- i.svcs[svc]
		}
		close(next)
	}()
	return next
}

// Err implements DependencyIterator.Err
func (i Iterator) Err() error {
	return i.err
}

// ReverseIterator iterates over the sorted elements in reverse order.
type ReverseIterator struct {
	*Dependencies
}

// Next implements DependencyIterator.Next
func (i *ReverseIterator) Next() <-chan Service {
	next := make(chan Service)
	go func() {
		for x := len(i.path); x > 0; x-- {
			next <- i.svcs[i.path[x-1]]
		}
		close(next)
	}()
	return next
}

// Err implements DependencyIterator.Err
func (i ReverseIterator) Err() error {
	return i.err
}
