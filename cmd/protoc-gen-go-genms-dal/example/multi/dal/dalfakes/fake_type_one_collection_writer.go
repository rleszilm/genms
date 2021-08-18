// Code generated by counterfeiter. DO NOT EDIT.
package dalfakes

import (
	"context"
	"sync"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	dal_multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal"
)

type FakeTypeOneCollectionWriter struct {
	InsertStub        func(context.Context, *multi.TypeOne) (*multi.TypeOne, error)
	insertMutex       sync.RWMutex
	insertArgsForCall []struct {
		arg1 context.Context
		arg2 *multi.TypeOne
	}
	insertReturns struct {
		result1 *multi.TypeOne
		result2 error
	}
	insertReturnsOnCall map[int]struct {
		result1 *multi.TypeOne
		result2 error
	}
	UpdateStub        func(context.Context, *multi.TypeOne, *dal_multi.TypeOneFieldValues) (*multi.TypeOne, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 *multi.TypeOne
		arg3 *dal_multi.TypeOneFieldValues
	}
	updateReturns struct {
		result1 *multi.TypeOne
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *multi.TypeOne
		result2 error
	}
	UpsertStub        func(context.Context, *multi.TypeOne) (*multi.TypeOne, error)
	upsertMutex       sync.RWMutex
	upsertArgsForCall []struct {
		arg1 context.Context
		arg2 *multi.TypeOne
	}
	upsertReturns struct {
		result1 *multi.TypeOne
		result2 error
	}
	upsertReturnsOnCall map[int]struct {
		result1 *multi.TypeOne
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTypeOneCollectionWriter) Insert(arg1 context.Context, arg2 *multi.TypeOne) (*multi.TypeOne, error) {
	fake.insertMutex.Lock()
	ret, specificReturn := fake.insertReturnsOnCall[len(fake.insertArgsForCall)]
	fake.insertArgsForCall = append(fake.insertArgsForCall, struct {
		arg1 context.Context
		arg2 *multi.TypeOne
	}{arg1, arg2})
	fake.recordInvocation("Insert", []interface{}{arg1, arg2})
	fake.insertMutex.Unlock()
	if fake.InsertStub != nil {
		return fake.InsertStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.insertReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeOneCollectionWriter) InsertCallCount() int {
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	return len(fake.insertArgsForCall)
}

func (fake *FakeTypeOneCollectionWriter) InsertCalls(stub func(context.Context, *multi.TypeOne) (*multi.TypeOne, error)) {
	fake.insertMutex.Lock()
	defer fake.insertMutex.Unlock()
	fake.InsertStub = stub
}

func (fake *FakeTypeOneCollectionWriter) InsertArgsForCall(i int) (context.Context, *multi.TypeOne) {
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	argsForCall := fake.insertArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTypeOneCollectionWriter) InsertReturns(result1 *multi.TypeOne, result2 error) {
	fake.insertMutex.Lock()
	defer fake.insertMutex.Unlock()
	fake.InsertStub = nil
	fake.insertReturns = struct {
		result1 *multi.TypeOne
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeOneCollectionWriter) InsertReturnsOnCall(i int, result1 *multi.TypeOne, result2 error) {
	fake.insertMutex.Lock()
	defer fake.insertMutex.Unlock()
	fake.InsertStub = nil
	if fake.insertReturnsOnCall == nil {
		fake.insertReturnsOnCall = make(map[int]struct {
			result1 *multi.TypeOne
			result2 error
		})
	}
	fake.insertReturnsOnCall[i] = struct {
		result1 *multi.TypeOne
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeOneCollectionWriter) Update(arg1 context.Context, arg2 *multi.TypeOne, arg3 *dal_multi.TypeOneFieldValues) (*multi.TypeOne, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 *multi.TypeOne
		arg3 *dal_multi.TypeOneFieldValues
	}{arg1, arg2, arg3})
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeOneCollectionWriter) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeTypeOneCollectionWriter) UpdateCalls(stub func(context.Context, *multi.TypeOne, *dal_multi.TypeOneFieldValues) (*multi.TypeOne, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeTypeOneCollectionWriter) UpdateArgsForCall(i int) (context.Context, *multi.TypeOne, *dal_multi.TypeOneFieldValues) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTypeOneCollectionWriter) UpdateReturns(result1 *multi.TypeOne, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *multi.TypeOne
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeOneCollectionWriter) UpdateReturnsOnCall(i int, result1 *multi.TypeOne, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *multi.TypeOne
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *multi.TypeOne
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeOneCollectionWriter) Upsert(arg1 context.Context, arg2 *multi.TypeOne) (*multi.TypeOne, error) {
	fake.upsertMutex.Lock()
	ret, specificReturn := fake.upsertReturnsOnCall[len(fake.upsertArgsForCall)]
	fake.upsertArgsForCall = append(fake.upsertArgsForCall, struct {
		arg1 context.Context
		arg2 *multi.TypeOne
	}{arg1, arg2})
	fake.recordInvocation("Upsert", []interface{}{arg1, arg2})
	fake.upsertMutex.Unlock()
	if fake.UpsertStub != nil {
		return fake.UpsertStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.upsertReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeOneCollectionWriter) UpsertCallCount() int {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	return len(fake.upsertArgsForCall)
}

func (fake *FakeTypeOneCollectionWriter) UpsertCalls(stub func(context.Context, *multi.TypeOne) (*multi.TypeOne, error)) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = stub
}

func (fake *FakeTypeOneCollectionWriter) UpsertArgsForCall(i int) (context.Context, *multi.TypeOne) {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	argsForCall := fake.upsertArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTypeOneCollectionWriter) UpsertReturns(result1 *multi.TypeOne, result2 error) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = nil
	fake.upsertReturns = struct {
		result1 *multi.TypeOne
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeOneCollectionWriter) UpsertReturnsOnCall(i int, result1 *multi.TypeOne, result2 error) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = nil
	if fake.upsertReturnsOnCall == nil {
		fake.upsertReturnsOnCall = make(map[int]struct {
			result1 *multi.TypeOne
			result2 error
		})
	}
	fake.upsertReturnsOnCall[i] = struct {
		result1 *multi.TypeOne
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeOneCollectionWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.insertMutex.RLock()
	defer fake.insertMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTypeOneCollectionWriter) recordInvocation(key string, args []interface{}) {
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

var _ dal_multi.TypeOneCollectionWriter = new(FakeTypeOneCollectionWriter)
