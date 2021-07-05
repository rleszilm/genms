// Code generated by counterfeiter. DO NOT EDIT.
package cachefakes

import (
	"context"
	"sync"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
	cache_dal_single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single/dal/cache"
)

type FakeSingleReadWriter struct {
	GetStub        func(context.Context, cache_dal_single.SingleKey) (*single.Single, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 cache_dal_single.SingleKey
	}
	getReturns struct {
		result1 *single.Single
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *single.Single
		result2 error
	}
	SetStub        func(context.Context, cache_dal_single.SingleKey, *single.Single) (*single.Single, error)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 context.Context
		arg2 cache_dal_single.SingleKey
		arg3 *single.Single
	}
	setReturns struct {
		result1 *single.Single
		result2 error
	}
	setReturnsOnCall map[int]struct {
		result1 *single.Single
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSingleReadWriter) Get(arg1 context.Context, arg2 cache_dal_single.SingleKey) (*single.Single, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 cache_dal_single.SingleKey
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSingleReadWriter) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSingleReadWriter) GetCalls(stub func(context.Context, cache_dal_single.SingleKey) (*single.Single, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeSingleReadWriter) GetArgsForCall(i int) (context.Context, cache_dal_single.SingleKey) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSingleReadWriter) GetReturns(result1 *single.Single, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *single.Single
		result2 error
	}{result1, result2}
}

func (fake *FakeSingleReadWriter) GetReturnsOnCall(i int, result1 *single.Single, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *single.Single
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *single.Single
		result2 error
	}{result1, result2}
}

func (fake *FakeSingleReadWriter) Set(arg1 context.Context, arg2 cache_dal_single.SingleKey, arg3 *single.Single) (*single.Single, error) {
	fake.setMutex.Lock()
	ret, specificReturn := fake.setReturnsOnCall[len(fake.setArgsForCall)]
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 context.Context
		arg2 cache_dal_single.SingleKey
		arg3 *single.Single
	}{arg1, arg2, arg3})
	fake.recordInvocation("Set", []interface{}{arg1, arg2, arg3})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		return fake.SetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSingleReadWriter) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeSingleReadWriter) SetCalls(stub func(context.Context, cache_dal_single.SingleKey, *single.Single) (*single.Single, error)) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *FakeSingleReadWriter) SetArgsForCall(i int) (context.Context, cache_dal_single.SingleKey, *single.Single) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSingleReadWriter) SetReturns(result1 *single.Single, result2 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 *single.Single
		result2 error
	}{result1, result2}
}

func (fake *FakeSingleReadWriter) SetReturnsOnCall(i int, result1 *single.Single, result2 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	if fake.setReturnsOnCall == nil {
		fake.setReturnsOnCall = make(map[int]struct {
			result1 *single.Single
			result2 error
		})
	}
	fake.setReturnsOnCall[i] = struct {
		result1 *single.Single
		result2 error
	}{result1, result2}
}

func (fake *FakeSingleReadWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSingleReadWriter) recordInvocation(key string, args []interface{}) {
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

var _ cache_dal_single.SingleReadWriter = new(FakeSingleReadWriter)