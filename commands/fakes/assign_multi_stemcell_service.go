// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type AssignMultiStemcellService struct {
	AssignMultiStemcellStub        func(api.ProductMultiStemcells) error
	assignMultiStemcellMutex       sync.RWMutex
	assignMultiStemcellArgsForCall []struct {
		arg1 api.ProductMultiStemcells
	}
	assignMultiStemcellReturns struct {
		result1 error
	}
	assignMultiStemcellReturnsOnCall map[int]struct {
		result1 error
	}
	InfoStub        func() (api.Info, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
	}
	infoReturns struct {
		result1 api.Info
		result2 error
	}
	infoReturnsOnCall map[int]struct {
		result1 api.Info
		result2 error
	}
	ListMultiStemcellsStub        func() (api.ProductMultiStemcells, error)
	listMultiStemcellsMutex       sync.RWMutex
	listMultiStemcellsArgsForCall []struct {
	}
	listMultiStemcellsReturns struct {
		result1 api.ProductMultiStemcells
		result2 error
	}
	listMultiStemcellsReturnsOnCall map[int]struct {
		result1 api.ProductMultiStemcells
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AssignMultiStemcellService) AssignMultiStemcell(arg1 api.ProductMultiStemcells) error {
	fake.assignMultiStemcellMutex.Lock()
	ret, specificReturn := fake.assignMultiStemcellReturnsOnCall[len(fake.assignMultiStemcellArgsForCall)]
	fake.assignMultiStemcellArgsForCall = append(fake.assignMultiStemcellArgsForCall, struct {
		arg1 api.ProductMultiStemcells
	}{arg1})
	fake.recordInvocation("AssignMultiStemcell", []interface{}{arg1})
	fake.assignMultiStemcellMutex.Unlock()
	if fake.AssignMultiStemcellStub != nil {
		return fake.AssignMultiStemcellStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.assignMultiStemcellReturns
	return fakeReturns.result1
}

func (fake *AssignMultiStemcellService) AssignMultiStemcellCallCount() int {
	fake.assignMultiStemcellMutex.RLock()
	defer fake.assignMultiStemcellMutex.RUnlock()
	return len(fake.assignMultiStemcellArgsForCall)
}

func (fake *AssignMultiStemcellService) AssignMultiStemcellCalls(stub func(api.ProductMultiStemcells) error) {
	fake.assignMultiStemcellMutex.Lock()
	defer fake.assignMultiStemcellMutex.Unlock()
	fake.AssignMultiStemcellStub = stub
}

func (fake *AssignMultiStemcellService) AssignMultiStemcellArgsForCall(i int) api.ProductMultiStemcells {
	fake.assignMultiStemcellMutex.RLock()
	defer fake.assignMultiStemcellMutex.RUnlock()
	argsForCall := fake.assignMultiStemcellArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AssignMultiStemcellService) AssignMultiStemcellReturns(result1 error) {
	fake.assignMultiStemcellMutex.Lock()
	defer fake.assignMultiStemcellMutex.Unlock()
	fake.AssignMultiStemcellStub = nil
	fake.assignMultiStemcellReturns = struct {
		result1 error
	}{result1}
}

func (fake *AssignMultiStemcellService) AssignMultiStemcellReturnsOnCall(i int, result1 error) {
	fake.assignMultiStemcellMutex.Lock()
	defer fake.assignMultiStemcellMutex.Unlock()
	fake.AssignMultiStemcellStub = nil
	if fake.assignMultiStemcellReturnsOnCall == nil {
		fake.assignMultiStemcellReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.assignMultiStemcellReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *AssignMultiStemcellService) Info() (api.Info, error) {
	fake.infoMutex.Lock()
	ret, specificReturn := fake.infoReturnsOnCall[len(fake.infoArgsForCall)]
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
	}{})
	fake.recordInvocation("Info", []interface{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.infoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AssignMultiStemcellService) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *AssignMultiStemcellService) InfoCalls(stub func() (api.Info, error)) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = stub
}

func (fake *AssignMultiStemcellService) InfoReturns(result1 api.Info, result2 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 api.Info
		result2 error
	}{result1, result2}
}

func (fake *AssignMultiStemcellService) InfoReturnsOnCall(i int, result1 api.Info, result2 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	if fake.infoReturnsOnCall == nil {
		fake.infoReturnsOnCall = make(map[int]struct {
			result1 api.Info
			result2 error
		})
	}
	fake.infoReturnsOnCall[i] = struct {
		result1 api.Info
		result2 error
	}{result1, result2}
}

func (fake *AssignMultiStemcellService) ListMultiStemcells() (api.ProductMultiStemcells, error) {
	fake.listMultiStemcellsMutex.Lock()
	ret, specificReturn := fake.listMultiStemcellsReturnsOnCall[len(fake.listMultiStemcellsArgsForCall)]
	fake.listMultiStemcellsArgsForCall = append(fake.listMultiStemcellsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListMultiStemcells", []interface{}{})
	fake.listMultiStemcellsMutex.Unlock()
	if fake.ListMultiStemcellsStub != nil {
		return fake.ListMultiStemcellsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listMultiStemcellsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AssignMultiStemcellService) ListMultiStemcellsCallCount() int {
	fake.listMultiStemcellsMutex.RLock()
	defer fake.listMultiStemcellsMutex.RUnlock()
	return len(fake.listMultiStemcellsArgsForCall)
}

func (fake *AssignMultiStemcellService) ListMultiStemcellsCalls(stub func() (api.ProductMultiStemcells, error)) {
	fake.listMultiStemcellsMutex.Lock()
	defer fake.listMultiStemcellsMutex.Unlock()
	fake.ListMultiStemcellsStub = stub
}

func (fake *AssignMultiStemcellService) ListMultiStemcellsReturns(result1 api.ProductMultiStemcells, result2 error) {
	fake.listMultiStemcellsMutex.Lock()
	defer fake.listMultiStemcellsMutex.Unlock()
	fake.ListMultiStemcellsStub = nil
	fake.listMultiStemcellsReturns = struct {
		result1 api.ProductMultiStemcells
		result2 error
	}{result1, result2}
}

func (fake *AssignMultiStemcellService) ListMultiStemcellsReturnsOnCall(i int, result1 api.ProductMultiStemcells, result2 error) {
	fake.listMultiStemcellsMutex.Lock()
	defer fake.listMultiStemcellsMutex.Unlock()
	fake.ListMultiStemcellsStub = nil
	if fake.listMultiStemcellsReturnsOnCall == nil {
		fake.listMultiStemcellsReturnsOnCall = make(map[int]struct {
			result1 api.ProductMultiStemcells
			result2 error
		})
	}
	fake.listMultiStemcellsReturnsOnCall[i] = struct {
		result1 api.ProductMultiStemcells
		result2 error
	}{result1, result2}
}

func (fake *AssignMultiStemcellService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.assignMultiStemcellMutex.RLock()
	defer fake.assignMultiStemcellMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.listMultiStemcellsMutex.RLock()
	defer fake.listMultiStemcellsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AssignMultiStemcellService) recordInvocation(key string, args []interface{}) {
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