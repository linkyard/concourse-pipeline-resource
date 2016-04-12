// This file was generated by counterfeiter
package flyfakes

import (
	"sync"

	"github.com/robdimsdale/concourse-pipeline-resource/fly"
)

type FakeFlyConn struct {
	LoginStub        func(url string, username string, password string) ([]byte, error)
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		url      string
		username string
		password string
	}
	loginReturns struct {
		result1 []byte
		result2 error
	}
	GetPipelineStub        func(pipelineName string) ([]byte, error)
	getPipelineMutex       sync.RWMutex
	getPipelineArgsForCall []struct {
		pipelineName string
	}
	getPipelineReturns struct {
		result1 []byte
		result2 error
	}
	SetPipelineStub        func(pipelineName string, configFilepath string, varsFilepaths []string) ([]byte, error)
	setPipelineMutex       sync.RWMutex
	setPipelineArgsForCall []struct {
		pipelineName   string
		configFilepath string
		varsFilepaths  []string
	}
	setPipelineReturns struct {
		result1 []byte
		result2 error
	}
	DestroyPipelineStub        func(pipelineName string) ([]byte, error)
	destroyPipelineMutex       sync.RWMutex
	destroyPipelineArgsForCall []struct {
		pipelineName string
	}
	destroyPipelineReturns struct {
		result1 []byte
		result2 error
	}
}

func (fake *FakeFlyConn) Login(url string, username string, password string) ([]byte, error) {
	fake.loginMutex.Lock()
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		url      string
		username string
		password string
	}{url, username, password})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub(url, username, password)
	} else {
		return fake.loginReturns.result1, fake.loginReturns.result2
	}
}

func (fake *FakeFlyConn) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeFlyConn) LoginArgsForCall(i int) (string, string, string) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return fake.loginArgsForCall[i].url, fake.loginArgsForCall[i].username, fake.loginArgsForCall[i].password
}

func (fake *FakeFlyConn) LoginReturns(result1 []byte, result2 error) {
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlyConn) GetPipeline(pipelineName string) ([]byte, error) {
	fake.getPipelineMutex.Lock()
	fake.getPipelineArgsForCall = append(fake.getPipelineArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.getPipelineMutex.Unlock()
	if fake.GetPipelineStub != nil {
		return fake.GetPipelineStub(pipelineName)
	} else {
		return fake.getPipelineReturns.result1, fake.getPipelineReturns.result2
	}
}

func (fake *FakeFlyConn) GetPipelineCallCount() int {
	fake.getPipelineMutex.RLock()
	defer fake.getPipelineMutex.RUnlock()
	return len(fake.getPipelineArgsForCall)
}

func (fake *FakeFlyConn) GetPipelineArgsForCall(i int) string {
	fake.getPipelineMutex.RLock()
	defer fake.getPipelineMutex.RUnlock()
	return fake.getPipelineArgsForCall[i].pipelineName
}

func (fake *FakeFlyConn) GetPipelineReturns(result1 []byte, result2 error) {
	fake.GetPipelineStub = nil
	fake.getPipelineReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlyConn) SetPipeline(pipelineName string, configFilepath string, varsFilepaths []string) ([]byte, error) {
	fake.setPipelineMutex.Lock()
	fake.setPipelineArgsForCall = append(fake.setPipelineArgsForCall, struct {
		pipelineName   string
		configFilepath string
		varsFilepaths  []string
	}{pipelineName, configFilepath, varsFilepaths})
	fake.setPipelineMutex.Unlock()
	if fake.SetPipelineStub != nil {
		return fake.SetPipelineStub(pipelineName, configFilepath, varsFilepaths)
	} else {
		return fake.setPipelineReturns.result1, fake.setPipelineReturns.result2
	}
}

func (fake *FakeFlyConn) SetPipelineCallCount() int {
	fake.setPipelineMutex.RLock()
	defer fake.setPipelineMutex.RUnlock()
	return len(fake.setPipelineArgsForCall)
}

func (fake *FakeFlyConn) SetPipelineArgsForCall(i int) (string, string, []string) {
	fake.setPipelineMutex.RLock()
	defer fake.setPipelineMutex.RUnlock()
	return fake.setPipelineArgsForCall[i].pipelineName, fake.setPipelineArgsForCall[i].configFilepath, fake.setPipelineArgsForCall[i].varsFilepaths
}

func (fake *FakeFlyConn) SetPipelineReturns(result1 []byte, result2 error) {
	fake.SetPipelineStub = nil
	fake.setPipelineReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlyConn) DestroyPipeline(pipelineName string) ([]byte, error) {
	fake.destroyPipelineMutex.Lock()
	fake.destroyPipelineArgsForCall = append(fake.destroyPipelineArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.destroyPipelineMutex.Unlock()
	if fake.DestroyPipelineStub != nil {
		return fake.DestroyPipelineStub(pipelineName)
	} else {
		return fake.destroyPipelineReturns.result1, fake.destroyPipelineReturns.result2
	}
}

func (fake *FakeFlyConn) DestroyPipelineCallCount() int {
	fake.destroyPipelineMutex.RLock()
	defer fake.destroyPipelineMutex.RUnlock()
	return len(fake.destroyPipelineArgsForCall)
}

func (fake *FakeFlyConn) DestroyPipelineArgsForCall(i int) string {
	fake.destroyPipelineMutex.RLock()
	defer fake.destroyPipelineMutex.RUnlock()
	return fake.destroyPipelineArgsForCall[i].pipelineName
}

func (fake *FakeFlyConn) DestroyPipelineReturns(result1 []byte, result2 error) {
	fake.DestroyPipelineStub = nil
	fake.destroyPipelineReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

var _ fly.FlyConn = new(FakeFlyConn)
