// Code generated by counterfeiter. DO NOT EDIT.
package configfakes

import (
	"sync"

	"github.com/nginxinc/nginx-gateway-kubernetes/internal/nginx/config"
	"github.com/nginxinc/nginx-gateway-kubernetes/internal/state"
)

type FakeGenerator struct {
	GenerateForHostStub        func(state.Host) ([]byte, config.Warnings)
	generateForHostMutex       sync.RWMutex
	generateForHostArgsForCall []struct {
		arg1 state.Host
	}
	generateForHostReturns struct {
		result1 []byte
		result2 config.Warnings
	}
	generateForHostReturnsOnCall map[int]struct {
		result1 []byte
		result2 config.Warnings
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGenerator) GenerateForHost(arg1 state.Host) ([]byte, config.Warnings) {
	fake.generateForHostMutex.Lock()
	ret, specificReturn := fake.generateForHostReturnsOnCall[len(fake.generateForHostArgsForCall)]
	fake.generateForHostArgsForCall = append(fake.generateForHostArgsForCall, struct {
		arg1 state.Host
	}{arg1})
	stub := fake.GenerateForHostStub
	fakeReturns := fake.generateForHostReturns
	fake.recordInvocation("GenerateForHost", []interface{}{arg1})
	fake.generateForHostMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGenerator) GenerateForHostCallCount() int {
	fake.generateForHostMutex.RLock()
	defer fake.generateForHostMutex.RUnlock()
	return len(fake.generateForHostArgsForCall)
}

func (fake *FakeGenerator) GenerateForHostCalls(stub func(state.Host) ([]byte, config.Warnings)) {
	fake.generateForHostMutex.Lock()
	defer fake.generateForHostMutex.Unlock()
	fake.GenerateForHostStub = stub
}

func (fake *FakeGenerator) GenerateForHostArgsForCall(i int) state.Host {
	fake.generateForHostMutex.RLock()
	defer fake.generateForHostMutex.RUnlock()
	argsForCall := fake.generateForHostArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGenerator) GenerateForHostReturns(result1 []byte, result2 config.Warnings) {
	fake.generateForHostMutex.Lock()
	defer fake.generateForHostMutex.Unlock()
	fake.GenerateForHostStub = nil
	fake.generateForHostReturns = struct {
		result1 []byte
		result2 config.Warnings
	}{result1, result2}
}

func (fake *FakeGenerator) GenerateForHostReturnsOnCall(i int, result1 []byte, result2 config.Warnings) {
	fake.generateForHostMutex.Lock()
	defer fake.generateForHostMutex.Unlock()
	fake.GenerateForHostStub = nil
	if fake.generateForHostReturnsOnCall == nil {
		fake.generateForHostReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 config.Warnings
		})
	}
	fake.generateForHostReturnsOnCall[i] = struct {
		result1 []byte
		result2 config.Warnings
	}{result1, result2}
}

func (fake *FakeGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateForHostMutex.RLock()
	defer fake.generateForHostMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGenerator) recordInvocation(key string, args []interface{}) {
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

var _ config.Generator = new(FakeGenerator)