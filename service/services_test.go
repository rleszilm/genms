package service_test

import (
	"reflect"
	"testing"

	"github.com/rleszilm/genms/service"
	"github.com/rleszilm/genms/service/servicefakes"
)

func TestServicesSort(t *testing.T) {
	svc1 := &servicefakes.FakeService{}
	svc1.StringReturns("svc-1")

	svc2 := &servicefakes.FakeService{}
	svc2.StringReturns("svc-2")

	svc3 := &servicefakes.FakeService{}
	svc3.StringReturns("svc-3")

	svc4 := &servicefakes.FakeService{}
	svc4.StringReturns("svc-4")

	svc5 := &servicefakes.FakeService{}
	svc5.StringReturns("svc-5")

	testcases := []struct {
		desc   string
		svcs   []service.Service
		deps   map[*servicefakes.FakeService][]service.Service
		expect []service.Service
		cycle  []service.Service
		err    error
	}{
		{
			desc:   "no dependencies",
			svcs:   []service.Service{svc1, svc2, svc3, svc4, svc5},
			expect: []service.Service{svc1, svc2, svc3, svc4, svc5},
		},
		{
			desc: "many dependencies",
			svcs: []service.Service{svc1, svc2, svc3, svc4, svc5},
			deps: map[*servicefakes.FakeService][]service.Service{
				svc1: {svc2, svc3, svc4, svc5},
				svc2: {svc3, svc4, svc5},
				svc3: {svc4, svc5},
				svc4: {svc5},
				svc5: {},
			},
			expect: []service.Service{svc5, svc4, svc3, svc2, svc1},
		},
		{
			desc: "branching dependencies",
			svcs: []service.Service{svc1, svc2, svc3, svc4, svc5},
			deps: map[*servicefakes.FakeService][]service.Service{
				svc1: {svc2, svc3},
				svc2: {svc4},
				svc3: {svc5},
				svc4: {},
				svc5: {},
			},
			expect: []service.Service{svc4, svc2, svc5, svc3, svc1},
		},
		{
			desc: "simple cycle",
			svcs: []service.Service{svc1, svc2},
			deps: map[*servicefakes.FakeService][]service.Service{
				svc1: {svc2},
				svc2: {svc1},
			},
			cycle: []service.Service{svc1, svc2, svc1},
			err:   service.ErrDependencyCycle,
		},
		{
			desc: "longer cycle",
			svcs: []service.Service{svc1, svc2, svc3, svc4},
			deps: map[*servicefakes.FakeService][]service.Service{
				svc1: {svc2},
				svc2: {svc3},
				svc3: {svc4},
				svc4: {svc1},
			},
			cycle: []service.Service{svc1, svc2, svc3, svc4, svc1},
			err:   service.ErrDependencyCycle,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			for svc, deps := range tc.deps {
				svc.DependantsReturns(deps)
			}

			svcs := service.Services(tc.svcs)
			cycle, err := svcs.Sort()
			if tc.err != err {
				t.Errorf("Unexpected err when sorting. Expected: %v Actual: %v", tc.err, err)
				return
			} else if err != nil {
				return
			}

			if !reflect.DeepEqual(cycle, tc.cycle) {
				t.Errorf("Cycle order is not as expected.Expected:\n%+v\nActual:%+v\n", tc.cycle, cycle)
			}

			expected := service.Services(tc.expect)
			if !reflect.DeepEqual(svcs, expected) {
				t.Errorf("Service order is not as expected.\nExpected:\n%+v\nActual:\n%+v\n", expected, svcs)
			}
		})
	}
}
