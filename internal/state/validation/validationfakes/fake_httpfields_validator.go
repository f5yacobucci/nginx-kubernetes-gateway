// Code generated by counterfeiter. DO NOT EDIT.
package validationfakes

import (
	"sync"

	"github.com/nginxinc/nginx-kubernetes-gateway/internal/state/validation"
)

type FakeHTTPFieldsValidator struct {
	ValidateHeaderNameInMatchStub        func(string) error
	validateHeaderNameInMatchMutex       sync.RWMutex
	validateHeaderNameInMatchArgsForCall []struct {
		arg1 string
	}
	validateHeaderNameInMatchReturns struct {
		result1 error
	}
	validateHeaderNameInMatchReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateHeaderValueInMatchStub        func(string) error
	validateHeaderValueInMatchMutex       sync.RWMutex
	validateHeaderValueInMatchArgsForCall []struct {
		arg1 string
	}
	validateHeaderValueInMatchReturns struct {
		result1 error
	}
	validateHeaderValueInMatchReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateMethodInMatchStub        func(string) (bool, []string)
	validateMethodInMatchMutex       sync.RWMutex
	validateMethodInMatchArgsForCall []struct {
		arg1 string
	}
	validateMethodInMatchReturns struct {
		result1 bool
		result2 []string
	}
	validateMethodInMatchReturnsOnCall map[int]struct {
		result1 bool
		result2 []string
	}
	ValidatePathInPrefixMatchStub        func(string) error
	validatePathInPrefixMatchMutex       sync.RWMutex
	validatePathInPrefixMatchArgsForCall []struct {
		arg1 string
	}
	validatePathInPrefixMatchReturns struct {
		result1 error
	}
	validatePathInPrefixMatchReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateQueryParamNameInMatchStub        func(string) error
	validateQueryParamNameInMatchMutex       sync.RWMutex
	validateQueryParamNameInMatchArgsForCall []struct {
		arg1 string
	}
	validateQueryParamNameInMatchReturns struct {
		result1 error
	}
	validateQueryParamNameInMatchReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateQueryParamValueInMatchStub        func(string) error
	validateQueryParamValueInMatchMutex       sync.RWMutex
	validateQueryParamValueInMatchArgsForCall []struct {
		arg1 string
	}
	validateQueryParamValueInMatchReturns struct {
		result1 error
	}
	validateQueryParamValueInMatchReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateRedirectHostnameStub        func(string) error
	validateRedirectHostnameMutex       sync.RWMutex
	validateRedirectHostnameArgsForCall []struct {
		arg1 string
	}
	validateRedirectHostnameReturns struct {
		result1 error
	}
	validateRedirectHostnameReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateRedirectPortStub        func(int32) error
	validateRedirectPortMutex       sync.RWMutex
	validateRedirectPortArgsForCall []struct {
		arg1 int32
	}
	validateRedirectPortReturns struct {
		result1 error
	}
	validateRedirectPortReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateRedirectSchemeStub        func(string) (bool, []string)
	validateRedirectSchemeMutex       sync.RWMutex
	validateRedirectSchemeArgsForCall []struct {
		arg1 string
	}
	validateRedirectSchemeReturns struct {
		result1 bool
		result2 []string
	}
	validateRedirectSchemeReturnsOnCall map[int]struct {
		result1 bool
		result2 []string
	}
	ValidateRedirectStatusCodeStub        func(int) (bool, []string)
	validateRedirectStatusCodeMutex       sync.RWMutex
	validateRedirectStatusCodeArgsForCall []struct {
		arg1 int
	}
	validateRedirectStatusCodeReturns struct {
		result1 bool
		result2 []string
	}
	validateRedirectStatusCodeReturnsOnCall map[int]struct {
		result1 bool
		result2 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderNameInMatch(arg1 string) error {
	fake.validateHeaderNameInMatchMutex.Lock()
	ret, specificReturn := fake.validateHeaderNameInMatchReturnsOnCall[len(fake.validateHeaderNameInMatchArgsForCall)]
	fake.validateHeaderNameInMatchArgsForCall = append(fake.validateHeaderNameInMatchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidateHeaderNameInMatchStub
	fakeReturns := fake.validateHeaderNameInMatchReturns
	fake.recordInvocation("ValidateHeaderNameInMatch", []interface{}{arg1})
	fake.validateHeaderNameInMatchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderNameInMatchCallCount() int {
	fake.validateHeaderNameInMatchMutex.RLock()
	defer fake.validateHeaderNameInMatchMutex.RUnlock()
	return len(fake.validateHeaderNameInMatchArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderNameInMatchCalls(stub func(string) error) {
	fake.validateHeaderNameInMatchMutex.Lock()
	defer fake.validateHeaderNameInMatchMutex.Unlock()
	fake.ValidateHeaderNameInMatchStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderNameInMatchArgsForCall(i int) string {
	fake.validateHeaderNameInMatchMutex.RLock()
	defer fake.validateHeaderNameInMatchMutex.RUnlock()
	argsForCall := fake.validateHeaderNameInMatchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderNameInMatchReturns(result1 error) {
	fake.validateHeaderNameInMatchMutex.Lock()
	defer fake.validateHeaderNameInMatchMutex.Unlock()
	fake.ValidateHeaderNameInMatchStub = nil
	fake.validateHeaderNameInMatchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderNameInMatchReturnsOnCall(i int, result1 error) {
	fake.validateHeaderNameInMatchMutex.Lock()
	defer fake.validateHeaderNameInMatchMutex.Unlock()
	fake.ValidateHeaderNameInMatchStub = nil
	if fake.validateHeaderNameInMatchReturnsOnCall == nil {
		fake.validateHeaderNameInMatchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateHeaderNameInMatchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderValueInMatch(arg1 string) error {
	fake.validateHeaderValueInMatchMutex.Lock()
	ret, specificReturn := fake.validateHeaderValueInMatchReturnsOnCall[len(fake.validateHeaderValueInMatchArgsForCall)]
	fake.validateHeaderValueInMatchArgsForCall = append(fake.validateHeaderValueInMatchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidateHeaderValueInMatchStub
	fakeReturns := fake.validateHeaderValueInMatchReturns
	fake.recordInvocation("ValidateHeaderValueInMatch", []interface{}{arg1})
	fake.validateHeaderValueInMatchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderValueInMatchCallCount() int {
	fake.validateHeaderValueInMatchMutex.RLock()
	defer fake.validateHeaderValueInMatchMutex.RUnlock()
	return len(fake.validateHeaderValueInMatchArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderValueInMatchCalls(stub func(string) error) {
	fake.validateHeaderValueInMatchMutex.Lock()
	defer fake.validateHeaderValueInMatchMutex.Unlock()
	fake.ValidateHeaderValueInMatchStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderValueInMatchArgsForCall(i int) string {
	fake.validateHeaderValueInMatchMutex.RLock()
	defer fake.validateHeaderValueInMatchMutex.RUnlock()
	argsForCall := fake.validateHeaderValueInMatchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderValueInMatchReturns(result1 error) {
	fake.validateHeaderValueInMatchMutex.Lock()
	defer fake.validateHeaderValueInMatchMutex.Unlock()
	fake.ValidateHeaderValueInMatchStub = nil
	fake.validateHeaderValueInMatchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateHeaderValueInMatchReturnsOnCall(i int, result1 error) {
	fake.validateHeaderValueInMatchMutex.Lock()
	defer fake.validateHeaderValueInMatchMutex.Unlock()
	fake.ValidateHeaderValueInMatchStub = nil
	if fake.validateHeaderValueInMatchReturnsOnCall == nil {
		fake.validateHeaderValueInMatchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateHeaderValueInMatchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateMethodInMatch(arg1 string) (bool, []string) {
	fake.validateMethodInMatchMutex.Lock()
	ret, specificReturn := fake.validateMethodInMatchReturnsOnCall[len(fake.validateMethodInMatchArgsForCall)]
	fake.validateMethodInMatchArgsForCall = append(fake.validateMethodInMatchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidateMethodInMatchStub
	fakeReturns := fake.validateMethodInMatchReturns
	fake.recordInvocation("ValidateMethodInMatch", []interface{}{arg1})
	fake.validateMethodInMatchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHTTPFieldsValidator) ValidateMethodInMatchCallCount() int {
	fake.validateMethodInMatchMutex.RLock()
	defer fake.validateMethodInMatchMutex.RUnlock()
	return len(fake.validateMethodInMatchArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateMethodInMatchCalls(stub func(string) (bool, []string)) {
	fake.validateMethodInMatchMutex.Lock()
	defer fake.validateMethodInMatchMutex.Unlock()
	fake.ValidateMethodInMatchStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateMethodInMatchArgsForCall(i int) string {
	fake.validateMethodInMatchMutex.RLock()
	defer fake.validateMethodInMatchMutex.RUnlock()
	argsForCall := fake.validateMethodInMatchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateMethodInMatchReturns(result1 bool, result2 []string) {
	fake.validateMethodInMatchMutex.Lock()
	defer fake.validateMethodInMatchMutex.Unlock()
	fake.ValidateMethodInMatchStub = nil
	fake.validateMethodInMatchReturns = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeHTTPFieldsValidator) ValidateMethodInMatchReturnsOnCall(i int, result1 bool, result2 []string) {
	fake.validateMethodInMatchMutex.Lock()
	defer fake.validateMethodInMatchMutex.Unlock()
	fake.ValidateMethodInMatchStub = nil
	if fake.validateMethodInMatchReturnsOnCall == nil {
		fake.validateMethodInMatchReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 []string
		})
	}
	fake.validateMethodInMatchReturnsOnCall[i] = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeHTTPFieldsValidator) ValidatePathInPrefixMatch(arg1 string) error {
	fake.validatePathInPrefixMatchMutex.Lock()
	ret, specificReturn := fake.validatePathInPrefixMatchReturnsOnCall[len(fake.validatePathInPrefixMatchArgsForCall)]
	fake.validatePathInPrefixMatchArgsForCall = append(fake.validatePathInPrefixMatchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidatePathInPrefixMatchStub
	fakeReturns := fake.validatePathInPrefixMatchReturns
	fake.recordInvocation("ValidatePathInPrefixMatch", []interface{}{arg1})
	fake.validatePathInPrefixMatchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPFieldsValidator) ValidatePathInPrefixMatchCallCount() int {
	fake.validatePathInPrefixMatchMutex.RLock()
	defer fake.validatePathInPrefixMatchMutex.RUnlock()
	return len(fake.validatePathInPrefixMatchArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidatePathInPrefixMatchCalls(stub func(string) error) {
	fake.validatePathInPrefixMatchMutex.Lock()
	defer fake.validatePathInPrefixMatchMutex.Unlock()
	fake.ValidatePathInPrefixMatchStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidatePathInPrefixMatchArgsForCall(i int) string {
	fake.validatePathInPrefixMatchMutex.RLock()
	defer fake.validatePathInPrefixMatchMutex.RUnlock()
	argsForCall := fake.validatePathInPrefixMatchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidatePathInPrefixMatchReturns(result1 error) {
	fake.validatePathInPrefixMatchMutex.Lock()
	defer fake.validatePathInPrefixMatchMutex.Unlock()
	fake.ValidatePathInPrefixMatchStub = nil
	fake.validatePathInPrefixMatchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidatePathInPrefixMatchReturnsOnCall(i int, result1 error) {
	fake.validatePathInPrefixMatchMutex.Lock()
	defer fake.validatePathInPrefixMatchMutex.Unlock()
	fake.ValidatePathInPrefixMatchStub = nil
	if fake.validatePathInPrefixMatchReturnsOnCall == nil {
		fake.validatePathInPrefixMatchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validatePathInPrefixMatchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamNameInMatch(arg1 string) error {
	fake.validateQueryParamNameInMatchMutex.Lock()
	ret, specificReturn := fake.validateQueryParamNameInMatchReturnsOnCall[len(fake.validateQueryParamNameInMatchArgsForCall)]
	fake.validateQueryParamNameInMatchArgsForCall = append(fake.validateQueryParamNameInMatchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidateQueryParamNameInMatchStub
	fakeReturns := fake.validateQueryParamNameInMatchReturns
	fake.recordInvocation("ValidateQueryParamNameInMatch", []interface{}{arg1})
	fake.validateQueryParamNameInMatchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamNameInMatchCallCount() int {
	fake.validateQueryParamNameInMatchMutex.RLock()
	defer fake.validateQueryParamNameInMatchMutex.RUnlock()
	return len(fake.validateQueryParamNameInMatchArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamNameInMatchCalls(stub func(string) error) {
	fake.validateQueryParamNameInMatchMutex.Lock()
	defer fake.validateQueryParamNameInMatchMutex.Unlock()
	fake.ValidateQueryParamNameInMatchStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamNameInMatchArgsForCall(i int) string {
	fake.validateQueryParamNameInMatchMutex.RLock()
	defer fake.validateQueryParamNameInMatchMutex.RUnlock()
	argsForCall := fake.validateQueryParamNameInMatchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamNameInMatchReturns(result1 error) {
	fake.validateQueryParamNameInMatchMutex.Lock()
	defer fake.validateQueryParamNameInMatchMutex.Unlock()
	fake.ValidateQueryParamNameInMatchStub = nil
	fake.validateQueryParamNameInMatchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamNameInMatchReturnsOnCall(i int, result1 error) {
	fake.validateQueryParamNameInMatchMutex.Lock()
	defer fake.validateQueryParamNameInMatchMutex.Unlock()
	fake.ValidateQueryParamNameInMatchStub = nil
	if fake.validateQueryParamNameInMatchReturnsOnCall == nil {
		fake.validateQueryParamNameInMatchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateQueryParamNameInMatchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamValueInMatch(arg1 string) error {
	fake.validateQueryParamValueInMatchMutex.Lock()
	ret, specificReturn := fake.validateQueryParamValueInMatchReturnsOnCall[len(fake.validateQueryParamValueInMatchArgsForCall)]
	fake.validateQueryParamValueInMatchArgsForCall = append(fake.validateQueryParamValueInMatchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidateQueryParamValueInMatchStub
	fakeReturns := fake.validateQueryParamValueInMatchReturns
	fake.recordInvocation("ValidateQueryParamValueInMatch", []interface{}{arg1})
	fake.validateQueryParamValueInMatchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamValueInMatchCallCount() int {
	fake.validateQueryParamValueInMatchMutex.RLock()
	defer fake.validateQueryParamValueInMatchMutex.RUnlock()
	return len(fake.validateQueryParamValueInMatchArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamValueInMatchCalls(stub func(string) error) {
	fake.validateQueryParamValueInMatchMutex.Lock()
	defer fake.validateQueryParamValueInMatchMutex.Unlock()
	fake.ValidateQueryParamValueInMatchStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamValueInMatchArgsForCall(i int) string {
	fake.validateQueryParamValueInMatchMutex.RLock()
	defer fake.validateQueryParamValueInMatchMutex.RUnlock()
	argsForCall := fake.validateQueryParamValueInMatchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamValueInMatchReturns(result1 error) {
	fake.validateQueryParamValueInMatchMutex.Lock()
	defer fake.validateQueryParamValueInMatchMutex.Unlock()
	fake.ValidateQueryParamValueInMatchStub = nil
	fake.validateQueryParamValueInMatchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateQueryParamValueInMatchReturnsOnCall(i int, result1 error) {
	fake.validateQueryParamValueInMatchMutex.Lock()
	defer fake.validateQueryParamValueInMatchMutex.Unlock()
	fake.ValidateQueryParamValueInMatchStub = nil
	if fake.validateQueryParamValueInMatchReturnsOnCall == nil {
		fake.validateQueryParamValueInMatchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateQueryParamValueInMatchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectHostname(arg1 string) error {
	fake.validateRedirectHostnameMutex.Lock()
	ret, specificReturn := fake.validateRedirectHostnameReturnsOnCall[len(fake.validateRedirectHostnameArgsForCall)]
	fake.validateRedirectHostnameArgsForCall = append(fake.validateRedirectHostnameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidateRedirectHostnameStub
	fakeReturns := fake.validateRedirectHostnameReturns
	fake.recordInvocation("ValidateRedirectHostname", []interface{}{arg1})
	fake.validateRedirectHostnameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectHostnameCallCount() int {
	fake.validateRedirectHostnameMutex.RLock()
	defer fake.validateRedirectHostnameMutex.RUnlock()
	return len(fake.validateRedirectHostnameArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectHostnameCalls(stub func(string) error) {
	fake.validateRedirectHostnameMutex.Lock()
	defer fake.validateRedirectHostnameMutex.Unlock()
	fake.ValidateRedirectHostnameStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectHostnameArgsForCall(i int) string {
	fake.validateRedirectHostnameMutex.RLock()
	defer fake.validateRedirectHostnameMutex.RUnlock()
	argsForCall := fake.validateRedirectHostnameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectHostnameReturns(result1 error) {
	fake.validateRedirectHostnameMutex.Lock()
	defer fake.validateRedirectHostnameMutex.Unlock()
	fake.ValidateRedirectHostnameStub = nil
	fake.validateRedirectHostnameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectHostnameReturnsOnCall(i int, result1 error) {
	fake.validateRedirectHostnameMutex.Lock()
	defer fake.validateRedirectHostnameMutex.Unlock()
	fake.ValidateRedirectHostnameStub = nil
	if fake.validateRedirectHostnameReturnsOnCall == nil {
		fake.validateRedirectHostnameReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateRedirectHostnameReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectPort(arg1 int32) error {
	fake.validateRedirectPortMutex.Lock()
	ret, specificReturn := fake.validateRedirectPortReturnsOnCall[len(fake.validateRedirectPortArgsForCall)]
	fake.validateRedirectPortArgsForCall = append(fake.validateRedirectPortArgsForCall, struct {
		arg1 int32
	}{arg1})
	stub := fake.ValidateRedirectPortStub
	fakeReturns := fake.validateRedirectPortReturns
	fake.recordInvocation("ValidateRedirectPort", []interface{}{arg1})
	fake.validateRedirectPortMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectPortCallCount() int {
	fake.validateRedirectPortMutex.RLock()
	defer fake.validateRedirectPortMutex.RUnlock()
	return len(fake.validateRedirectPortArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectPortCalls(stub func(int32) error) {
	fake.validateRedirectPortMutex.Lock()
	defer fake.validateRedirectPortMutex.Unlock()
	fake.ValidateRedirectPortStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectPortArgsForCall(i int) int32 {
	fake.validateRedirectPortMutex.RLock()
	defer fake.validateRedirectPortMutex.RUnlock()
	argsForCall := fake.validateRedirectPortArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectPortReturns(result1 error) {
	fake.validateRedirectPortMutex.Lock()
	defer fake.validateRedirectPortMutex.Unlock()
	fake.ValidateRedirectPortStub = nil
	fake.validateRedirectPortReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectPortReturnsOnCall(i int, result1 error) {
	fake.validateRedirectPortMutex.Lock()
	defer fake.validateRedirectPortMutex.Unlock()
	fake.ValidateRedirectPortStub = nil
	if fake.validateRedirectPortReturnsOnCall == nil {
		fake.validateRedirectPortReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateRedirectPortReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectScheme(arg1 string) (bool, []string) {
	fake.validateRedirectSchemeMutex.Lock()
	ret, specificReturn := fake.validateRedirectSchemeReturnsOnCall[len(fake.validateRedirectSchemeArgsForCall)]
	fake.validateRedirectSchemeArgsForCall = append(fake.validateRedirectSchemeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ValidateRedirectSchemeStub
	fakeReturns := fake.validateRedirectSchemeReturns
	fake.recordInvocation("ValidateRedirectScheme", []interface{}{arg1})
	fake.validateRedirectSchemeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectSchemeCallCount() int {
	fake.validateRedirectSchemeMutex.RLock()
	defer fake.validateRedirectSchemeMutex.RUnlock()
	return len(fake.validateRedirectSchemeArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectSchemeCalls(stub func(string) (bool, []string)) {
	fake.validateRedirectSchemeMutex.Lock()
	defer fake.validateRedirectSchemeMutex.Unlock()
	fake.ValidateRedirectSchemeStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectSchemeArgsForCall(i int) string {
	fake.validateRedirectSchemeMutex.RLock()
	defer fake.validateRedirectSchemeMutex.RUnlock()
	argsForCall := fake.validateRedirectSchemeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectSchemeReturns(result1 bool, result2 []string) {
	fake.validateRedirectSchemeMutex.Lock()
	defer fake.validateRedirectSchemeMutex.Unlock()
	fake.ValidateRedirectSchemeStub = nil
	fake.validateRedirectSchemeReturns = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectSchemeReturnsOnCall(i int, result1 bool, result2 []string) {
	fake.validateRedirectSchemeMutex.Lock()
	defer fake.validateRedirectSchemeMutex.Unlock()
	fake.ValidateRedirectSchemeStub = nil
	if fake.validateRedirectSchemeReturnsOnCall == nil {
		fake.validateRedirectSchemeReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 []string
		})
	}
	fake.validateRedirectSchemeReturnsOnCall[i] = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectStatusCode(arg1 int) (bool, []string) {
	fake.validateRedirectStatusCodeMutex.Lock()
	ret, specificReturn := fake.validateRedirectStatusCodeReturnsOnCall[len(fake.validateRedirectStatusCodeArgsForCall)]
	fake.validateRedirectStatusCodeArgsForCall = append(fake.validateRedirectStatusCodeArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.ValidateRedirectStatusCodeStub
	fakeReturns := fake.validateRedirectStatusCodeReturns
	fake.recordInvocation("ValidateRedirectStatusCode", []interface{}{arg1})
	fake.validateRedirectStatusCodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectStatusCodeCallCount() int {
	fake.validateRedirectStatusCodeMutex.RLock()
	defer fake.validateRedirectStatusCodeMutex.RUnlock()
	return len(fake.validateRedirectStatusCodeArgsForCall)
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectStatusCodeCalls(stub func(int) (bool, []string)) {
	fake.validateRedirectStatusCodeMutex.Lock()
	defer fake.validateRedirectStatusCodeMutex.Unlock()
	fake.ValidateRedirectStatusCodeStub = stub
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectStatusCodeArgsForCall(i int) int {
	fake.validateRedirectStatusCodeMutex.RLock()
	defer fake.validateRedirectStatusCodeMutex.RUnlock()
	argsForCall := fake.validateRedirectStatusCodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectStatusCodeReturns(result1 bool, result2 []string) {
	fake.validateRedirectStatusCodeMutex.Lock()
	defer fake.validateRedirectStatusCodeMutex.Unlock()
	fake.ValidateRedirectStatusCodeStub = nil
	fake.validateRedirectStatusCodeReturns = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeHTTPFieldsValidator) ValidateRedirectStatusCodeReturnsOnCall(i int, result1 bool, result2 []string) {
	fake.validateRedirectStatusCodeMutex.Lock()
	defer fake.validateRedirectStatusCodeMutex.Unlock()
	fake.ValidateRedirectStatusCodeStub = nil
	if fake.validateRedirectStatusCodeReturnsOnCall == nil {
		fake.validateRedirectStatusCodeReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 []string
		})
	}
	fake.validateRedirectStatusCodeReturnsOnCall[i] = struct {
		result1 bool
		result2 []string
	}{result1, result2}
}

func (fake *FakeHTTPFieldsValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateHeaderNameInMatchMutex.RLock()
	defer fake.validateHeaderNameInMatchMutex.RUnlock()
	fake.validateHeaderValueInMatchMutex.RLock()
	defer fake.validateHeaderValueInMatchMutex.RUnlock()
	fake.validateMethodInMatchMutex.RLock()
	defer fake.validateMethodInMatchMutex.RUnlock()
	fake.validatePathInPrefixMatchMutex.RLock()
	defer fake.validatePathInPrefixMatchMutex.RUnlock()
	fake.validateQueryParamNameInMatchMutex.RLock()
	defer fake.validateQueryParamNameInMatchMutex.RUnlock()
	fake.validateQueryParamValueInMatchMutex.RLock()
	defer fake.validateQueryParamValueInMatchMutex.RUnlock()
	fake.validateRedirectHostnameMutex.RLock()
	defer fake.validateRedirectHostnameMutex.RUnlock()
	fake.validateRedirectPortMutex.RLock()
	defer fake.validateRedirectPortMutex.RUnlock()
	fake.validateRedirectSchemeMutex.RLock()
	defer fake.validateRedirectSchemeMutex.RUnlock()
	fake.validateRedirectStatusCodeMutex.RLock()
	defer fake.validateRedirectStatusCodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHTTPFieldsValidator) recordInvocation(key string, args []interface{}) {
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

var _ validation.HTTPFieldsValidator = new(FakeHTTPFieldsValidator)