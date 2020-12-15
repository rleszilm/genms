// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	"github.com/rleszilm/gen_microservice/service"
)

type FakeService struct {
	DependenciesStub        func() service.Services
	dependenciesMutex       sync.RWMutex
	dependenciesArgsForCall []struct {
	}
	dependenciesReturns struct {
		result1 service.Services
	}
	dependenciesReturnsOnCall map[int]struct {
		result1 service.Services
	}
	InitializeStub        func(context.Context) error
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct {
		arg1 context.Context
	}
	initializeReturns struct {
		result1 error
	}
	initializeReturnsOnCall map[int]struct {
		result1 error
	}
	NameOfStub        func() string
	nameOfMutex       sync.RWMutex
	nameOfArgsForCall []struct {
	}
	nameOfReturns struct {
		result1 string
	}
	nameOfReturnsOnCall map[int]struct {
		result1 string
	}
	ShutdownStub        func(context.Context) error
	shutdownMutex       sync.RWMutex
	shutdownArgsForCall []struct {
		arg1 context.Context
	}
	shutdownReturns struct {
		result1 error
	}
	shutdownReturnsOnCall map[int]struct {
		result1 error
	}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct {
	}
	stringReturns struct {
		result1 string
	}
	stringReturnsOnCall map[int]struct {
		result1 string
	}
	WithDependenciesStub        func(...service.Service)
	withDependenciesMutex       sync.RWMutex
	withDependenciesArgsForCall []struct {
		arg1 []service.Service
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeService) Dependencies() service.Services {
	fake.dependenciesMutex.Lock()
	ret, specificReturn := fake.dependenciesReturnsOnCall[len(fake.dependenciesArgsForCall)]
	fake.dependenciesArgsForCall = append(fake.dependenciesArgsForCall, struct {
	}{})
	fake.recordInvocation("Dependencies", []interface{}{})
	fake.dependenciesMutex.Unlock()
	if fake.DependenciesStub != nil {
		return fake.DependenciesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dependenciesReturns
	return fakeReturns.result1
}

func (fake *FakeService) DependenciesCallCount() int {
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	return len(fake.dependenciesArgsForCall)
}

func (fake *FakeService) DependenciesCalls(stub func() service.Services) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = stub
}

func (fake *FakeService) DependenciesReturns(result1 service.Services) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = nil
	fake.dependenciesReturns = struct {
		result1 service.Services
	}{result1}
}

func (fake *FakeService) DependenciesReturnsOnCall(i int, result1 service.Services) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = nil
	if fake.dependenciesReturnsOnCall == nil {
		fake.dependenciesReturnsOnCall = make(map[int]struct {
			result1 service.Services
		})
	}
	fake.dependenciesReturnsOnCall[i] = struct {
		result1 service.Services
	}{result1}
}

func (fake *FakeService) Initialize(arg1 context.Context) error {
	fake.initializeMutex.Lock()
	ret, specificReturn := fake.initializeReturnsOnCall[len(fake.initializeArgsForCall)]
	fake.initializeArgsForCall = append(fake.initializeArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Initialize", []interface{}{arg1})
	fake.initializeMutex.Unlock()
	if fake.InitializeStub != nil {
		return fake.InitializeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initializeReturns
	return fakeReturns.result1
}

func (fake *FakeService) InitializeCallCount() int {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return len(fake.initializeArgsForCall)
}

func (fake *FakeService) InitializeCalls(stub func(context.Context) error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = stub
}

func (fake *FakeService) InitializeArgsForCall(i int) context.Context {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	argsForCall := fake.initializeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeService) InitializeReturns(result1 error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = nil
	fake.initializeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) InitializeReturnsOnCall(i int, result1 error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = nil
	if fake.initializeReturnsOnCall == nil {
		fake.initializeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) NameOf() string {
	fake.nameOfMutex.Lock()
	ret, specificReturn := fake.nameOfReturnsOnCall[len(fake.nameOfArgsForCall)]
	fake.nameOfArgsForCall = append(fake.nameOfArgsForCall, struct {
	}{})
	fake.recordInvocation("NameOf", []interface{}{})
	fake.nameOfMutex.Unlock()
	if fake.NameOfStub != nil {
		return fake.NameOfStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameOfReturns
	return fakeReturns.result1
}

func (fake *FakeService) NameOfCallCount() int {
	fake.nameOfMutex.RLock()
	defer fake.nameOfMutex.RUnlock()
	return len(fake.nameOfArgsForCall)
}

func (fake *FakeService) NameOfCalls(stub func() string) {
	fake.nameOfMutex.Lock()
	defer fake.nameOfMutex.Unlock()
	fake.NameOfStub = stub
}

func (fake *FakeService) NameOfReturns(result1 string) {
	fake.nameOfMutex.Lock()
	defer fake.nameOfMutex.Unlock()
	fake.NameOfStub = nil
	fake.nameOfReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeService) NameOfReturnsOnCall(i int, result1 string) {
	fake.nameOfMutex.Lock()
	defer fake.nameOfMutex.Unlock()
	fake.NameOfStub = nil
	if fake.nameOfReturnsOnCall == nil {
		fake.nameOfReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameOfReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeService) Shutdown(arg1 context.Context) error {
	fake.shutdownMutex.Lock()
	ret, specificReturn := fake.shutdownReturnsOnCall[len(fake.shutdownArgsForCall)]
	fake.shutdownArgsForCall = append(fake.shutdownArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Shutdown", []interface{}{arg1})
	fake.shutdownMutex.Unlock()
	if fake.ShutdownStub != nil {
		return fake.ShutdownStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.shutdownReturns
	return fakeReturns.result1
}

func (fake *FakeService) ShutdownCallCount() int {
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	return len(fake.shutdownArgsForCall)
}

func (fake *FakeService) ShutdownCalls(stub func(context.Context) error) {
	fake.shutdownMutex.Lock()
	defer fake.shutdownMutex.Unlock()
	fake.ShutdownStub = stub
}

func (fake *FakeService) ShutdownArgsForCall(i int) context.Context {
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	argsForCall := fake.shutdownArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeService) ShutdownReturns(result1 error) {
	fake.shutdownMutex.Lock()
	defer fake.shutdownMutex.Unlock()
	fake.ShutdownStub = nil
	fake.shutdownReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) ShutdownReturnsOnCall(i int, result1 error) {
	fake.shutdownMutex.Lock()
	defer fake.shutdownMutex.Unlock()
	fake.ShutdownStub = nil
	if fake.shutdownReturnsOnCall == nil {
		fake.shutdownReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.shutdownReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) String() string {
	fake.stringMutex.Lock()
	ret, specificReturn := fake.stringReturnsOnCall[len(fake.stringArgsForCall)]
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct {
	}{})
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if fake.StringStub != nil {
		return fake.StringStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stringReturns
	return fakeReturns.result1
}

func (fake *FakeService) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeService) StringCalls(stub func() string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = stub
}

func (fake *FakeService) StringReturns(result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeService) StringReturnsOnCall(i int, result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	if fake.stringReturnsOnCall == nil {
		fake.stringReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.stringReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeService) WithDependencies(arg1 ...service.Service) {
	fake.withDependenciesMutex.Lock()
	fake.withDependenciesArgsForCall = append(fake.withDependenciesArgsForCall, struct {
		arg1 []service.Service
	}{arg1})
	fake.recordInvocation("WithDependencies", []interface{}{arg1})
	fake.withDependenciesMutex.Unlock()
	if fake.WithDependenciesStub != nil {
		fake.WithDependenciesStub(arg1...)
	}
}

func (fake *FakeService) WithDependenciesCallCount() int {
	fake.withDependenciesMutex.RLock()
	defer fake.withDependenciesMutex.RUnlock()
	return len(fake.withDependenciesArgsForCall)
}

func (fake *FakeService) WithDependenciesCalls(stub func(...service.Service)) {
	fake.withDependenciesMutex.Lock()
	defer fake.withDependenciesMutex.Unlock()
	fake.WithDependenciesStub = stub
}

func (fake *FakeService) WithDependenciesArgsForCall(i int) []service.Service {
	fake.withDependenciesMutex.RLock()
	defer fake.withDependenciesMutex.RUnlock()
	argsForCall := fake.withDependenciesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	fake.nameOfMutex.RLock()
	defer fake.nameOfMutex.RUnlock()
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	fake.withDependenciesMutex.RLock()
	defer fake.withDependenciesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ service.Service = new(FakeService)
