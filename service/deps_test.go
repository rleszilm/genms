package service_test

import (
	"reflect"
	"testing"

	"github.com/rleszilm/gen_microservice/service"
	"github.com/rleszilm/gen_microservice/service/servicefakes"
)

func TestDependencies(t *testing.T) {
	svc1 := &servicefakes.FakeService{}
	svc1.NameReturns("svc-1")

	svc2 := &servicefakes.FakeService{}
	svc2.NameReturns("svc-2")

	svc3 := &servicefakes.FakeService{}
	svc3.NameReturns("svc-3")

	svc4 := &servicefakes.FakeService{}
	svc4.NameReturns("svc-4")

	svc5 := &servicefakes.FakeService{}
	svc5.NameReturns("svc-5")

	testcases := []struct {
		desc        string
		svcs        []service.Service
		deps        map[string][]service.Service
		expectSet   map[string]struct{}
		expectOrder []string
		err         error
	}{
		{
			desc: "no dependencies",
			svcs: []service.Service{svc1, svc2, svc3, svc4, svc5},
			expectSet: map[string]struct{}{
				"svc-1": {},
				"svc-2": {},
				"svc-3": {},
				"svc-4": {},
				"svc-5": {},
			},
		},
		{
			desc: "many dependencies",
			svcs: []service.Service{svc1, svc2, svc3, svc4, svc5},
			deps: map[string][]service.Service{
				"svc-1": {},
				"svc-2": {svc1},
				"svc-3": {svc1, svc2},
				"svc-4": {svc1, svc2, svc3},
				"svc-5": {svc1, svc2, svc3, svc4},
			},
			expectOrder: []string{
				"svc-1",
				"svc-2",
				"svc-3",
				"svc-4",
				"svc-5",
			},
		},
		{
			desc: "simple cycle",
			svcs: []service.Service{svc1, svc2},
			deps: map[string][]service.Service{
				"svc-1": {svc2},
				"svc-2": {svc1},
			},
			err: service.ErrDependencyCycle,
		},
		{
			desc: "longer cycle",
			svcs: []service.Service{svc1, svc2, svc3, svc4},
			deps: map[string][]service.Service{
				"svc-1": {svc2},
				"svc-2": {svc3},
				"svc-3": {svc4},
				"svc-4": {svc1},
			},
			err: service.ErrDependencyCycle,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			deps := service.NewDependencies()
			for _, svc := range tc.svcs {
				deps.Register(svc, tc.deps[svc.Name()]...)
			}

			it := deps.Iterate()

			actualSet := map[string]struct{}{}
			actualOrder := []string{}
			for svc := range it.Next() {
				actualSet[svc.Name()] = struct{}{}
				actualOrder = append(actualOrder, svc.Name())
			}

			if it.Err() != tc.err {
				t.Errorf("Unexpected error when iterating. Expected: %v Actual: %v", tc.err, it.Err())
				return
			}

			if tc.expectSet != nil && !reflect.DeepEqual(actualSet, tc.expectSet) {
				t.Errorf("Service set is not as expected\nExpected:\n%v\nActual:\n%v\n", tc.expectSet, actualSet)
			}

			if tc.expectOrder != nil && !reflect.DeepEqual(actualOrder, tc.expectOrder) {
				t.Errorf("Service order is not as expected\nExpected:\n%v\nActual:\n%v\n", tc.expectOrder, actualOrder)
			}
		})
	}
}
