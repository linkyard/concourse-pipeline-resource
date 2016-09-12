// This file was generated by counterfeiter
package checkfakes

import (
	"sync"

	"github.com/robdimsdale/concourse-pipeline-resource/check"
)

type FakeLogger struct {
	DebugfStub        func(format string, a ...interface{}) (n int, err error)
	debugfMutex       sync.RWMutex
	debugfArgsForCall []struct {
		format string
		a      []interface{}
	}
	debugfReturns struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogger) Debugf(format string, a ...interface{}) (n int, err error) {
	fake.debugfMutex.Lock()
	fake.debugfArgsForCall = append(fake.debugfArgsForCall, struct {
		format string
		a      []interface{}
	}{format, a})
	fake.recordInvocation("Debugf", []interface{}{format, a})
	fake.debugfMutex.Unlock()
	if fake.DebugfStub != nil {
		return fake.DebugfStub(format, a...)
	} else {
		return fake.debugfReturns.result1, fake.debugfReturns.result2
	}
}

func (fake *FakeLogger) DebugfCallCount() int {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return len(fake.debugfArgsForCall)
}

func (fake *FakeLogger) DebugfArgsForCall(i int) (string, []interface{}) {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return fake.debugfArgsForCall[i].format, fake.debugfArgsForCall[i].a
}

func (fake *FakeLogger) DebugfReturns(result1 int, result2 error) {
	fake.DebugfStub = nil
	fake.debugfReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLogger) recordInvocation(key string, args []interface{}) {
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

var _ check.Logger = new(FakeLogger)