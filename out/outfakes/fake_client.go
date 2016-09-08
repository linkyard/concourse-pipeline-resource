// This file was generated by counterfeiter
package outfakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/robdimsdale/concourse-pipeline-resource/concourse/api"
	"github.com/robdimsdale/concourse-pipeline-resource/out"
)

type FakeClient struct {
	PipelinesStub        func() ([]api.Pipeline, error)
	pipelinesMutex       sync.RWMutex
	pipelinesArgsForCall []struct{}
	pipelinesReturns     struct {
		result1 []api.Pipeline
		result2 error
	}
	PipelineConfigStub        func(pipelineName string) (config atc.Config, rawConfig string, version string, err error)
	pipelineConfigMutex       sync.RWMutex
	pipelineConfigArgsForCall []struct {
		pipelineName string
	}
	pipelineConfigReturns struct {
		result1 atc.Config
		result2 string
		result3 string
		result4 error
	}
	SetPipelineConfigStub        func(pipelineName string, configVersion string, passedConfig atc.Config) error
	setPipelineConfigMutex       sync.RWMutex
	setPipelineConfigArgsForCall []struct {
		pipelineName  string
		configVersion string
		passedConfig  atc.Config
	}
	setPipelineConfigReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Pipelines() ([]api.Pipeline, error) {
	fake.pipelinesMutex.Lock()
	fake.pipelinesArgsForCall = append(fake.pipelinesArgsForCall, struct{}{})
	fake.recordInvocation("Pipelines", []interface{}{})
	fake.pipelinesMutex.Unlock()
	if fake.PipelinesStub != nil {
		return fake.PipelinesStub()
	} else {
		return fake.pipelinesReturns.result1, fake.pipelinesReturns.result2
	}
}

func (fake *FakeClient) PipelinesCallCount() int {
	fake.pipelinesMutex.RLock()
	defer fake.pipelinesMutex.RUnlock()
	return len(fake.pipelinesArgsForCall)
}

func (fake *FakeClient) PipelinesReturns(result1 []api.Pipeline, result2 error) {
	fake.PipelinesStub = nil
	fake.pipelinesReturns = struct {
		result1 []api.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) PipelineConfig(pipelineName string) (config atc.Config, rawConfig string, version string, err error) {
	fake.pipelineConfigMutex.Lock()
	fake.pipelineConfigArgsForCall = append(fake.pipelineConfigArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.recordInvocation("PipelineConfig", []interface{}{pipelineName})
	fake.pipelineConfigMutex.Unlock()
	if fake.PipelineConfigStub != nil {
		return fake.PipelineConfigStub(pipelineName)
	} else {
		return fake.pipelineConfigReturns.result1, fake.pipelineConfigReturns.result2, fake.pipelineConfigReturns.result3, fake.pipelineConfigReturns.result4
	}
}

func (fake *FakeClient) PipelineConfigCallCount() int {
	fake.pipelineConfigMutex.RLock()
	defer fake.pipelineConfigMutex.RUnlock()
	return len(fake.pipelineConfigArgsForCall)
}

func (fake *FakeClient) PipelineConfigArgsForCall(i int) string {
	fake.pipelineConfigMutex.RLock()
	defer fake.pipelineConfigMutex.RUnlock()
	return fake.pipelineConfigArgsForCall[i].pipelineName
}

func (fake *FakeClient) PipelineConfigReturns(result1 atc.Config, result2 string, result3 string, result4 error) {
	fake.PipelineConfigStub = nil
	fake.pipelineConfigReturns = struct {
		result1 atc.Config
		result2 string
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeClient) SetPipelineConfig(pipelineName string, configVersion string, passedConfig atc.Config) error {
	fake.setPipelineConfigMutex.Lock()
	fake.setPipelineConfigArgsForCall = append(fake.setPipelineConfigArgsForCall, struct {
		pipelineName  string
		configVersion string
		passedConfig  atc.Config
	}{pipelineName, configVersion, passedConfig})
	fake.recordInvocation("SetPipelineConfig", []interface{}{pipelineName, configVersion, passedConfig})
	fake.setPipelineConfigMutex.Unlock()
	if fake.SetPipelineConfigStub != nil {
		return fake.SetPipelineConfigStub(pipelineName, configVersion, passedConfig)
	} else {
		return fake.setPipelineConfigReturns.result1
	}
}

func (fake *FakeClient) SetPipelineConfigCallCount() int {
	fake.setPipelineConfigMutex.RLock()
	defer fake.setPipelineConfigMutex.RUnlock()
	return len(fake.setPipelineConfigArgsForCall)
}

func (fake *FakeClient) SetPipelineConfigArgsForCall(i int) (string, string, atc.Config) {
	fake.setPipelineConfigMutex.RLock()
	defer fake.setPipelineConfigMutex.RUnlock()
	return fake.setPipelineConfigArgsForCall[i].pipelineName, fake.setPipelineConfigArgsForCall[i].configVersion, fake.setPipelineConfigArgsForCall[i].passedConfig
}

func (fake *FakeClient) SetPipelineConfigReturns(result1 error) {
	fake.SetPipelineConfigStub = nil
	fake.setPipelineConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pipelinesMutex.RLock()
	defer fake.pipelinesMutex.RUnlock()
	fake.pipelineConfigMutex.RLock()
	defer fake.pipelineConfigMutex.RUnlock()
	fake.setPipelineConfigMutex.RLock()
	defer fake.setPipelineConfigMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ out.Client = new(FakeClient)