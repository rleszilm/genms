// Code generated by counterfeiter. DO NOT EDIT.
package keyvaluefakes

import (
	"context"
	"sync"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
	keyvalue_dal_single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single/dal/keyvalue"
)

type FakeSingleReadWriter struct {
	GetByKeyStub        func(context.Context, keyvalue_dal_single.SingleKey) (*single.Single, error)
	getByKeyMutex       sync.RWMutex
	getByKeyArgsForCall []struct {
		arg1 context.Context
		arg2 keyvalue_dal_single.SingleKey
	}
	getByKeyReturns struct {
		result1 *single.Single
		result2 error
	}
	getByKeyReturnsOnCall map[int]struct {
		result1 *single.Single
		result2 error
	}
	SetByKeyStub        func(context.Context, keyvalue_dal_single.SingleKey, *single.Single) error
	setByKeyMutex       sync.RWMutex
	setByKeyArgsForCall []struct {
		arg1 context.Context
		arg2 keyvalue_dal_single.SingleKey
		arg3 *single.Single
	}
	setByKeyReturns struct {
		result1 error
	}
	setByKeyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSingleReadWriter) GetByKey(arg1 context.Context, arg2 keyvalue_dal_single.SingleKey) (*single.Single, error) {
	fake.getByKeyMutex.Lock()
	ret, specificReturn := fake.getByKeyReturnsOnCall[len(fake.getByKeyArgsForCall)]
	fake.getByKeyArgsForCall = append(fake.getByKeyArgsForCall, struct {
		arg1 context.Context
		arg2 keyvalue_dal_single.SingleKey
	}{arg1, arg2})
	fake.recordInvocation("GetByKey", []interface{}{arg1, arg2})
	fake.getByKeyMutex.Unlock()
	if fake.GetByKeyStub != nil {
		return fake.GetByKeyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getByKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSingleReadWriter) GetByKeyCallCount() int {
	fake.getByKeyMutex.RLock()
	defer fake.getByKeyMutex.RUnlock()
	return len(fake.getByKeyArgsForCall)
}

func (fake *FakeSingleReadWriter) GetByKeyCalls(stub func(context.Context, keyvalue_dal_single.SingleKey) (*single.Single, error)) {
	fake.getByKeyMutex.Lock()
	defer fake.getByKeyMutex.Unlock()
	fake.GetByKeyStub = stub
}

func (fake *FakeSingleReadWriter) GetByKeyArgsForCall(i int) (context.Context, keyvalue_dal_single.SingleKey) {
	fake.getByKeyMutex.RLock()
	defer fake.getByKeyMutex.RUnlock()
	argsForCall := fake.getByKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSingleReadWriter) GetByKeyReturns(result1 *single.Single, result2 error) {
	fake.getByKeyMutex.Lock()
	defer fake.getByKeyMutex.Unlock()
	fake.GetByKeyStub = nil
	fake.getByKeyReturns = struct {
		result1 *single.Single
		result2 error
	}{result1, result2}
}

func (fake *FakeSingleReadWriter) GetByKeyReturnsOnCall(i int, result1 *single.Single, result2 error) {
	fake.getByKeyMutex.Lock()
	defer fake.getByKeyMutex.Unlock()
	fake.GetByKeyStub = nil
	if fake.getByKeyReturnsOnCall == nil {
		fake.getByKeyReturnsOnCall = make(map[int]struct {
			result1 *single.Single
			result2 error
		})
	}
	fake.getByKeyReturnsOnCall[i] = struct {
		result1 *single.Single
		result2 error
	}{result1, result2}
}

func (fake *FakeSingleReadWriter) SetByKey(arg1 context.Context, arg2 keyvalue_dal_single.SingleKey, arg3 *single.Single) error {
	fake.setByKeyMutex.Lock()
	ret, specificReturn := fake.setByKeyReturnsOnCall[len(fake.setByKeyArgsForCall)]
	fake.setByKeyArgsForCall = append(fake.setByKeyArgsForCall, struct {
		arg1 context.Context
		arg2 keyvalue_dal_single.SingleKey
		arg3 *single.Single
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetByKey", []interface{}{arg1, arg2, arg3})
	fake.setByKeyMutex.Unlock()
	if fake.SetByKeyStub != nil {
		return fake.SetByKeyStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setByKeyReturns
	return fakeReturns.result1
}

func (fake *FakeSingleReadWriter) SetByKeyCallCount() int {
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
	return len(fake.setByKeyArgsForCall)
}

func (fake *FakeSingleReadWriter) SetByKeyCalls(stub func(context.Context, keyvalue_dal_single.SingleKey, *single.Single) error) {
	fake.setByKeyMutex.Lock()
	defer fake.setByKeyMutex.Unlock()
	fake.SetByKeyStub = stub
}

func (fake *FakeSingleReadWriter) SetByKeyArgsForCall(i int) (context.Context, keyvalue_dal_single.SingleKey, *single.Single) {
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
	argsForCall := fake.setByKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSingleReadWriter) SetByKeyReturns(result1 error) {
	fake.setByKeyMutex.Lock()
	defer fake.setByKeyMutex.Unlock()
	fake.SetByKeyStub = nil
	fake.setByKeyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSingleReadWriter) SetByKeyReturnsOnCall(i int, result1 error) {
	fake.setByKeyMutex.Lock()
	defer fake.setByKeyMutex.Unlock()
	fake.SetByKeyStub = nil
	if fake.setByKeyReturnsOnCall == nil {
		fake.setByKeyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setByKeyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSingleReadWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getByKeyMutex.RLock()
	defer fake.getByKeyMutex.RUnlock()
	fake.setByKeyMutex.RLock()
	defer fake.setByKeyMutex.RUnlock()
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

var _ keyvalue_dal_single.SingleReadWriter = new(FakeSingleReadWriter)
