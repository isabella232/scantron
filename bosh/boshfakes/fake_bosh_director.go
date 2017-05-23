// This file was generated by counterfeiter
package boshfakes

import (
	"sync"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
	"github.com/pivotal-cf/scantron/bosh"
	"github.com/pivotal-cf/scantron/remotemachine"
	"github.com/pivotal-cf/scantron/scanlog"
)

type FakeBoshDirector struct {
	VMsStub        func() []boshdir.VMInfo
	vMsMutex       sync.RWMutex
	vMsArgsForCall []struct{}
	vMsReturns     struct {
		result1 []boshdir.VMInfo
	}
	vMsReturnsOnCall map[int]struct {
		result1 []boshdir.VMInfo
	}
	ConnectToStub        func(scanlog.Logger, boshdir.VMInfo) remotemachine.RemoteMachine
	connectToMutex       sync.RWMutex
	connectToArgsForCall []struct {
		arg1 scanlog.Logger
		arg2 boshdir.VMInfo
	}
	connectToReturns struct {
		result1 remotemachine.RemoteMachine
	}
	connectToReturnsOnCall map[int]struct {
		result1 remotemachine.RemoteMachine
	}
	ReleasesStub        func() []boshdir.Release
	releasesMutex       sync.RWMutex
	releasesArgsForCall []struct{}
	releasesReturns     struct {
		result1 []boshdir.Release
	}
	releasesReturnsOnCall map[int]struct {
		result1 []boshdir.Release
	}
	SetupStub        func() error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct{}
	setupReturns     struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	cleanupReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBoshDirector) VMs() []boshdir.VMInfo {
	fake.vMsMutex.Lock()
	ret, specificReturn := fake.vMsReturnsOnCall[len(fake.vMsArgsForCall)]
	fake.vMsArgsForCall = append(fake.vMsArgsForCall, struct{}{})
	fake.recordInvocation("VMs", []interface{}{})
	fake.vMsMutex.Unlock()
	if fake.VMsStub != nil {
		return fake.VMsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.vMsReturns.result1
}

func (fake *FakeBoshDirector) VMsCallCount() int {
	fake.vMsMutex.RLock()
	defer fake.vMsMutex.RUnlock()
	return len(fake.vMsArgsForCall)
}

func (fake *FakeBoshDirector) VMsReturns(result1 []boshdir.VMInfo) {
	fake.VMsStub = nil
	fake.vMsReturns = struct {
		result1 []boshdir.VMInfo
	}{result1}
}

func (fake *FakeBoshDirector) VMsReturnsOnCall(i int, result1 []boshdir.VMInfo) {
	fake.VMsStub = nil
	if fake.vMsReturnsOnCall == nil {
		fake.vMsReturnsOnCall = make(map[int]struct {
			result1 []boshdir.VMInfo
		})
	}
	fake.vMsReturnsOnCall[i] = struct {
		result1 []boshdir.VMInfo
	}{result1}
}

func (fake *FakeBoshDirector) ConnectTo(arg1 scanlog.Logger, arg2 boshdir.VMInfo) remotemachine.RemoteMachine {
	fake.connectToMutex.Lock()
	ret, specificReturn := fake.connectToReturnsOnCall[len(fake.connectToArgsForCall)]
	fake.connectToArgsForCall = append(fake.connectToArgsForCall, struct {
		arg1 scanlog.Logger
		arg2 boshdir.VMInfo
	}{arg1, arg2})
	fake.recordInvocation("ConnectTo", []interface{}{arg1, arg2})
	fake.connectToMutex.Unlock()
	if fake.ConnectToStub != nil {
		return fake.ConnectToStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.connectToReturns.result1
}

func (fake *FakeBoshDirector) ConnectToCallCount() int {
	fake.connectToMutex.RLock()
	defer fake.connectToMutex.RUnlock()
	return len(fake.connectToArgsForCall)
}

func (fake *FakeBoshDirector) ConnectToArgsForCall(i int) (scanlog.Logger, boshdir.VMInfo) {
	fake.connectToMutex.RLock()
	defer fake.connectToMutex.RUnlock()
	return fake.connectToArgsForCall[i].arg1, fake.connectToArgsForCall[i].arg2
}

func (fake *FakeBoshDirector) ConnectToReturns(result1 remotemachine.RemoteMachine) {
	fake.ConnectToStub = nil
	fake.connectToReturns = struct {
		result1 remotemachine.RemoteMachine
	}{result1}
}

func (fake *FakeBoshDirector) ConnectToReturnsOnCall(i int, result1 remotemachine.RemoteMachine) {
	fake.ConnectToStub = nil
	if fake.connectToReturnsOnCall == nil {
		fake.connectToReturnsOnCall = make(map[int]struct {
			result1 remotemachine.RemoteMachine
		})
	}
	fake.connectToReturnsOnCall[i] = struct {
		result1 remotemachine.RemoteMachine
	}{result1}
}

func (fake *FakeBoshDirector) Releases() []boshdir.Release {
	fake.releasesMutex.Lock()
	ret, specificReturn := fake.releasesReturnsOnCall[len(fake.releasesArgsForCall)]
	fake.releasesArgsForCall = append(fake.releasesArgsForCall, struct{}{})
	fake.recordInvocation("Releases", []interface{}{})
	fake.releasesMutex.Unlock()
	if fake.ReleasesStub != nil {
		return fake.ReleasesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.releasesReturns.result1
}

func (fake *FakeBoshDirector) ReleasesCallCount() int {
	fake.releasesMutex.RLock()
	defer fake.releasesMutex.RUnlock()
	return len(fake.releasesArgsForCall)
}

func (fake *FakeBoshDirector) ReleasesReturns(result1 []boshdir.Release) {
	fake.ReleasesStub = nil
	fake.releasesReturns = struct {
		result1 []boshdir.Release
	}{result1}
}

func (fake *FakeBoshDirector) ReleasesReturnsOnCall(i int, result1 []boshdir.Release) {
	fake.ReleasesStub = nil
	if fake.releasesReturnsOnCall == nil {
		fake.releasesReturnsOnCall = make(map[int]struct {
			result1 []boshdir.Release
		})
	}
	fake.releasesReturnsOnCall[i] = struct {
		result1 []boshdir.Release
	}{result1}
}

func (fake *FakeBoshDirector) Setup() error {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct{}{})
	fake.recordInvocation("Setup", []interface{}{})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setupReturns.result1
}

func (fake *FakeBoshDirector) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeBoshDirector) SetupReturns(result1 error) {
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBoshDirector) SetupReturnsOnCall(i int, result1 error) {
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBoshDirector) Cleanup() error {
	fake.cleanupMutex.Lock()
	ret, specificReturn := fake.cleanupReturnsOnCall[len(fake.cleanupArgsForCall)]
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanupReturns.result1
}

func (fake *FakeBoshDirector) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeBoshDirector) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBoshDirector) CleanupReturnsOnCall(i int, result1 error) {
	fake.CleanupStub = nil
	if fake.cleanupReturnsOnCall == nil {
		fake.cleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBoshDirector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.vMsMutex.RLock()
	defer fake.vMsMutex.RUnlock()
	fake.connectToMutex.RLock()
	defer fake.connectToMutex.RUnlock()
	fake.releasesMutex.RLock()
	defer fake.releasesMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBoshDirector) recordInvocation(key string, args []interface{}) {
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

var _ bosh.BoshDirector = new(FakeBoshDirector)
