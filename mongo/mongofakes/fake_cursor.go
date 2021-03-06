// Code generated by counterfeiter. DO NOT EDIT.
package mongofakes

import (
	"context"
	"sync"

	"github.com/rleszilm/genms/mongo"
)

type FakeCursor struct {
	DecodeStub        func(interface{}) error
	decodeMutex       sync.RWMutex
	decodeArgsForCall []struct {
		arg1 interface{}
	}
	decodeReturns struct {
		result1 error
	}
	decodeReturnsOnCall map[int]struct {
		result1 error
	}
	NextStub        func(context.Context) bool
	nextMutex       sync.RWMutex
	nextArgsForCall []struct {
		arg1 context.Context
	}
	nextReturns struct {
		result1 bool
	}
	nextReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCursor) Decode(arg1 interface{}) error {
	fake.decodeMutex.Lock()
	ret, specificReturn := fake.decodeReturnsOnCall[len(fake.decodeArgsForCall)]
	fake.decodeArgsForCall = append(fake.decodeArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Decode", []interface{}{arg1})
	fake.decodeMutex.Unlock()
	if fake.DecodeStub != nil {
		return fake.DecodeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.decodeReturns
	return fakeReturns.result1
}

func (fake *FakeCursor) DecodeCallCount() int {
	fake.decodeMutex.RLock()
	defer fake.decodeMutex.RUnlock()
	return len(fake.decodeArgsForCall)
}

func (fake *FakeCursor) DecodeCalls(stub func(interface{}) error) {
	fake.decodeMutex.Lock()
	defer fake.decodeMutex.Unlock()
	fake.DecodeStub = stub
}

func (fake *FakeCursor) DecodeArgsForCall(i int) interface{} {
	fake.decodeMutex.RLock()
	defer fake.decodeMutex.RUnlock()
	argsForCall := fake.decodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCursor) DecodeReturns(result1 error) {
	fake.decodeMutex.Lock()
	defer fake.decodeMutex.Unlock()
	fake.DecodeStub = nil
	fake.decodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCursor) DecodeReturnsOnCall(i int, result1 error) {
	fake.decodeMutex.Lock()
	defer fake.decodeMutex.Unlock()
	fake.DecodeStub = nil
	if fake.decodeReturnsOnCall == nil {
		fake.decodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.decodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCursor) Next(arg1 context.Context) bool {
	fake.nextMutex.Lock()
	ret, specificReturn := fake.nextReturnsOnCall[len(fake.nextArgsForCall)]
	fake.nextArgsForCall = append(fake.nextArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Next", []interface{}{arg1})
	fake.nextMutex.Unlock()
	if fake.NextStub != nil {
		return fake.NextStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nextReturns
	return fakeReturns.result1
}

func (fake *FakeCursor) NextCallCount() int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	return len(fake.nextArgsForCall)
}

func (fake *FakeCursor) NextCalls(stub func(context.Context) bool) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = stub
}

func (fake *FakeCursor) NextArgsForCall(i int) context.Context {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	argsForCall := fake.nextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCursor) NextReturns(result1 bool) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	fake.nextReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCursor) NextReturnsOnCall(i int, result1 bool) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	if fake.nextReturnsOnCall == nil {
		fake.nextReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.nextReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCursor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.decodeMutex.RLock()
	defer fake.decodeMutex.RUnlock()
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCursor) recordInvocation(key string, args []interface{}) {
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

var _ mongo.Cursor = new(FakeCursor)
