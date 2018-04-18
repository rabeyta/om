// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type CredentialsService struct {
	GetDeployedProductCredentialStub        func(api.GetDeployedProductCredentialInput) (api.CredentialOutput, error)
	getDeployedProductCredentialMutex       sync.RWMutex
	getDeployedProductCredentialArgsForCall []struct {
		arg1 api.GetDeployedProductCredentialInput
	}
	getDeployedProductCredentialReturns struct {
		result1 api.CredentialOutput
		result2 error
	}
	getDeployedProductCredentialReturnsOnCall map[int]struct {
		result1 api.CredentialOutput
		result2 error
	}
	ListDeployedProductsStub        func() ([]api.DeployedProductOutput, error)
	listDeployedProductsMutex       sync.RWMutex
	listDeployedProductsArgsForCall []struct{}
	listDeployedProductsReturns     struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	listDeployedProductsReturnsOnCall map[int]struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CredentialsService) GetDeployedProductCredential(arg1 api.GetDeployedProductCredentialInput) (api.CredentialOutput, error) {
	fake.getDeployedProductCredentialMutex.Lock()
	ret, specificReturn := fake.getDeployedProductCredentialReturnsOnCall[len(fake.getDeployedProductCredentialArgsForCall)]
	fake.getDeployedProductCredentialArgsForCall = append(fake.getDeployedProductCredentialArgsForCall, struct {
		arg1 api.GetDeployedProductCredentialInput
	}{arg1})
	fake.recordInvocation("GetDeployedProductCredential", []interface{}{arg1})
	fake.getDeployedProductCredentialMutex.Unlock()
	if fake.GetDeployedProductCredentialStub != nil {
		return fake.GetDeployedProductCredentialStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeployedProductCredentialReturns.result1, fake.getDeployedProductCredentialReturns.result2
}

func (fake *CredentialsService) GetDeployedProductCredentialCallCount() int {
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	return len(fake.getDeployedProductCredentialArgsForCall)
}

func (fake *CredentialsService) GetDeployedProductCredentialArgsForCall(i int) api.GetDeployedProductCredentialInput {
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	return fake.getDeployedProductCredentialArgsForCall[i].arg1
}

func (fake *CredentialsService) GetDeployedProductCredentialReturns(result1 api.CredentialOutput, result2 error) {
	fake.GetDeployedProductCredentialStub = nil
	fake.getDeployedProductCredentialReturns = struct {
		result1 api.CredentialOutput
		result2 error
	}{result1, result2}
}

func (fake *CredentialsService) GetDeployedProductCredentialReturnsOnCall(i int, result1 api.CredentialOutput, result2 error) {
	fake.GetDeployedProductCredentialStub = nil
	if fake.getDeployedProductCredentialReturnsOnCall == nil {
		fake.getDeployedProductCredentialReturnsOnCall = make(map[int]struct {
			result1 api.CredentialOutput
			result2 error
		})
	}
	fake.getDeployedProductCredentialReturnsOnCall[i] = struct {
		result1 api.CredentialOutput
		result2 error
	}{result1, result2}
}

func (fake *CredentialsService) ListDeployedProducts() ([]api.DeployedProductOutput, error) {
	fake.listDeployedProductsMutex.Lock()
	ret, specificReturn := fake.listDeployedProductsReturnsOnCall[len(fake.listDeployedProductsArgsForCall)]
	fake.listDeployedProductsArgsForCall = append(fake.listDeployedProductsArgsForCall, struct{}{})
	fake.recordInvocation("ListDeployedProducts", []interface{}{})
	fake.listDeployedProductsMutex.Unlock()
	if fake.ListDeployedProductsStub != nil {
		return fake.ListDeployedProductsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDeployedProductsReturns.result1, fake.listDeployedProductsReturns.result2
}

func (fake *CredentialsService) ListDeployedProductsCallCount() int {
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	return len(fake.listDeployedProductsArgsForCall)
}

func (fake *CredentialsService) ListDeployedProductsReturns(result1 []api.DeployedProductOutput, result2 error) {
	fake.ListDeployedProductsStub = nil
	fake.listDeployedProductsReturns = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *CredentialsService) ListDeployedProductsReturnsOnCall(i int, result1 []api.DeployedProductOutput, result2 error) {
	fake.ListDeployedProductsStub = nil
	if fake.listDeployedProductsReturnsOnCall == nil {
		fake.listDeployedProductsReturnsOnCall = make(map[int]struct {
			result1 []api.DeployedProductOutput
			result2 error
		})
	}
	fake.listDeployedProductsReturnsOnCall[i] = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *CredentialsService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CredentialsService) recordInvocation(key string, args []interface{}) {
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
