// Code generated by counterfeiter. DO NOT EDIT.
package v2v3actionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2v3action"
	"code.cloudfoundry.org/cli/cf/manifest"
)

type FakeV2Actor struct {
	CreateManifestApplicationStub        func(string, string) (manifest.Application, v2action.Warnings, error)
	createManifestApplicationMutex       sync.RWMutex
	createManifestApplicationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createManifestApplicationReturns struct {
		result1 manifest.Application
		result2 v2action.Warnings
		result3 error
	}
	createManifestApplicationReturnsOnCall map[int]struct {
		result1 manifest.Application
		result2 v2action.Warnings
		result3 error
	}
	GetFeatureFlagsStub        func() ([]v2action.FeatureFlag, v2action.Warnings, error)
	getFeatureFlagsMutex       sync.RWMutex
	getFeatureFlagsArgsForCall []struct{}
	getFeatureFlagsReturns     struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}
	getFeatureFlagsReturnsOnCall map[int]struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}
	GetServiceStub        func(serviceGUID string) (v2action.Service, v2action.Warnings, error)
	getServiceMutex       sync.RWMutex
	getServiceArgsForCall []struct {
		serviceGUID string
	}
	getServiceReturns struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}
	getServiceReturnsOnCall map[int]struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}
	GetServiceInstanceByNameAndSpaceStub        func(serviceInstanceName string, spaceGUID string) (v2action.ServiceInstance, v2action.Warnings, error)
	getServiceInstanceByNameAndSpaceMutex       sync.RWMutex
	getServiceInstanceByNameAndSpaceArgsForCall []struct {
		serviceInstanceName string
		spaceGUID           string
	}
	getServiceInstanceByNameAndSpaceReturns struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}
	getServiceInstanceByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}
	GetServiceInstanceSharedTosByServiceInstanceStub        func(serviceInstanceGUID string) ([]v2action.ServiceInstanceSharedTo, v2action.Warnings, error)
	getServiceInstanceSharedTosByServiceInstanceMutex       sync.RWMutex
	getServiceInstanceSharedTosByServiceInstanceArgsForCall []struct {
		serviceInstanceGUID string
	}
	getServiceInstanceSharedTosByServiceInstanceReturns struct {
		result1 []v2action.ServiceInstanceSharedTo
		result2 v2action.Warnings
		result3 error
	}
	getServiceInstanceSharedTosByServiceInstanceReturnsOnCall map[int]struct {
		result1 []v2action.ServiceInstanceSharedTo
		result2 v2action.Warnings
		result3 error
	}
	GetSpaceByOrganizationAndNameStub        func(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error)
	getSpaceByOrganizationAndNameMutex       sync.RWMutex
	getSpaceByOrganizationAndNameArgsForCall []struct {
		orgGUID   string
		spaceName string
	}
	getSpaceByOrganizationAndNameReturns struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	getSpaceByOrganizationAndNameReturnsOnCall map[int]struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV2Actor) CreateManifestApplication(arg1 string, arg2 string) (manifest.Application, v2action.Warnings, error) {
	fake.createManifestApplicationMutex.Lock()
	ret, specificReturn := fake.createManifestApplicationReturnsOnCall[len(fake.createManifestApplicationArgsForCall)]
	fake.createManifestApplicationArgsForCall = append(fake.createManifestApplicationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateManifestApplication", []interface{}{arg1, arg2})
	fake.createManifestApplicationMutex.Unlock()
	if fake.CreateManifestApplicationStub != nil {
		return fake.CreateManifestApplicationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createManifestApplicationReturns.result1, fake.createManifestApplicationReturns.result2, fake.createManifestApplicationReturns.result3
}

func (fake *FakeV2Actor) CreateManifestApplicationCallCount() int {
	fake.createManifestApplicationMutex.RLock()
	defer fake.createManifestApplicationMutex.RUnlock()
	return len(fake.createManifestApplicationArgsForCall)
}

func (fake *FakeV2Actor) CreateManifestApplicationArgsForCall(i int) (string, string) {
	fake.createManifestApplicationMutex.RLock()
	defer fake.createManifestApplicationMutex.RUnlock()
	return fake.createManifestApplicationArgsForCall[i].arg1, fake.createManifestApplicationArgsForCall[i].arg2
}

func (fake *FakeV2Actor) CreateManifestApplicationReturns(result1 manifest.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateManifestApplicationStub = nil
	fake.createManifestApplicationReturns = struct {
		result1 manifest.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CreateManifestApplicationReturnsOnCall(i int, result1 manifest.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateManifestApplicationStub = nil
	if fake.createManifestApplicationReturnsOnCall == nil {
		fake.createManifestApplicationReturnsOnCall = make(map[int]struct {
			result1 manifest.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createManifestApplicationReturnsOnCall[i] = struct {
		result1 manifest.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetFeatureFlags() ([]v2action.FeatureFlag, v2action.Warnings, error) {
	fake.getFeatureFlagsMutex.Lock()
	ret, specificReturn := fake.getFeatureFlagsReturnsOnCall[len(fake.getFeatureFlagsArgsForCall)]
	fake.getFeatureFlagsArgsForCall = append(fake.getFeatureFlagsArgsForCall, struct{}{})
	fake.recordInvocation("GetFeatureFlags", []interface{}{})
	fake.getFeatureFlagsMutex.Unlock()
	if fake.GetFeatureFlagsStub != nil {
		return fake.GetFeatureFlagsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getFeatureFlagsReturns.result1, fake.getFeatureFlagsReturns.result2, fake.getFeatureFlagsReturns.result3
}

func (fake *FakeV2Actor) GetFeatureFlagsCallCount() int {
	fake.getFeatureFlagsMutex.RLock()
	defer fake.getFeatureFlagsMutex.RUnlock()
	return len(fake.getFeatureFlagsArgsForCall)
}

func (fake *FakeV2Actor) GetFeatureFlagsReturns(result1 []v2action.FeatureFlag, result2 v2action.Warnings, result3 error) {
	fake.GetFeatureFlagsStub = nil
	fake.getFeatureFlagsReturns = struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetFeatureFlagsReturnsOnCall(i int, result1 []v2action.FeatureFlag, result2 v2action.Warnings, result3 error) {
	fake.GetFeatureFlagsStub = nil
	if fake.getFeatureFlagsReturnsOnCall == nil {
		fake.getFeatureFlagsReturnsOnCall = make(map[int]struct {
			result1 []v2action.FeatureFlag
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getFeatureFlagsReturnsOnCall[i] = struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetService(serviceGUID string) (v2action.Service, v2action.Warnings, error) {
	fake.getServiceMutex.Lock()
	ret, specificReturn := fake.getServiceReturnsOnCall[len(fake.getServiceArgsForCall)]
	fake.getServiceArgsForCall = append(fake.getServiceArgsForCall, struct {
		serviceGUID string
	}{serviceGUID})
	fake.recordInvocation("GetService", []interface{}{serviceGUID})
	fake.getServiceMutex.Unlock()
	if fake.GetServiceStub != nil {
		return fake.GetServiceStub(serviceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getServiceReturns.result1, fake.getServiceReturns.result2, fake.getServiceReturns.result3
}

func (fake *FakeV2Actor) GetServiceCallCount() int {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return len(fake.getServiceArgsForCall)
}

func (fake *FakeV2Actor) GetServiceArgsForCall(i int) string {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return fake.getServiceArgsForCall[i].serviceGUID
}

func (fake *FakeV2Actor) GetServiceReturns(result1 v2action.Service, result2 v2action.Warnings, result3 error) {
	fake.GetServiceStub = nil
	fake.getServiceReturns = struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceReturnsOnCall(i int, result1 v2action.Service, result2 v2action.Warnings, result3 error) {
	fake.GetServiceStub = nil
	if fake.getServiceReturnsOnCall == nil {
		fake.getServiceReturnsOnCall = make(map[int]struct {
			result1 v2action.Service
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceReturnsOnCall[i] = struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpace(serviceInstanceName string, spaceGUID string) (v2action.ServiceInstance, v2action.Warnings, error) {
	fake.getServiceInstanceByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceByNameAndSpaceReturnsOnCall[len(fake.getServiceInstanceByNameAndSpaceArgsForCall)]
	fake.getServiceInstanceByNameAndSpaceArgsForCall = append(fake.getServiceInstanceByNameAndSpaceArgsForCall, struct {
		serviceInstanceName string
		spaceGUID           string
	}{serviceInstanceName, spaceGUID})
	fake.recordInvocation("GetServiceInstanceByNameAndSpace", []interface{}{serviceInstanceName, spaceGUID})
	fake.getServiceInstanceByNameAndSpaceMutex.Unlock()
	if fake.GetServiceInstanceByNameAndSpaceStub != nil {
		return fake.GetServiceInstanceByNameAndSpaceStub(serviceInstanceName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getServiceInstanceByNameAndSpaceReturns.result1, fake.getServiceInstanceByNameAndSpaceReturns.result2, fake.getServiceInstanceByNameAndSpaceReturns.result3
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpaceCallCount() int {
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	return len(fake.getServiceInstanceByNameAndSpaceArgsForCall)
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	return fake.getServiceInstanceByNameAndSpaceArgsForCall[i].serviceInstanceName, fake.getServiceInstanceByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpaceReturns(result1 v2action.ServiceInstance, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstanceByNameAndSpaceStub = nil
	fake.getServiceInstanceByNameAndSpaceReturns = struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpaceReturnsOnCall(i int, result1 v2action.ServiceInstance, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstanceByNameAndSpaceStub = nil
	if fake.getServiceInstanceByNameAndSpaceReturnsOnCall == nil {
		fake.getServiceInstanceByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.ServiceInstance
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceInstanceByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstance(serviceInstanceGUID string) ([]v2action.ServiceInstanceSharedTo, v2action.Warnings, error) {
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceSharedTosByServiceInstanceReturnsOnCall[len(fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall)]
	fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall = append(fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall, struct {
		serviceInstanceGUID string
	}{serviceInstanceGUID})
	fake.recordInvocation("GetServiceInstanceSharedTosByServiceInstance", []interface{}{serviceInstanceGUID})
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.Unlock()
	if fake.GetServiceInstanceSharedTosByServiceInstanceStub != nil {
		return fake.GetServiceInstanceSharedTosByServiceInstanceStub(serviceInstanceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getServiceInstanceSharedTosByServiceInstanceReturns.result1, fake.getServiceInstanceSharedTosByServiceInstanceReturns.result2, fake.getServiceInstanceSharedTosByServiceInstanceReturns.result3
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstanceCallCount() int {
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceSharedTosByServiceInstanceMutex.RUnlock()
	return len(fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall)
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstanceArgsForCall(i int) string {
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceSharedTosByServiceInstanceMutex.RUnlock()
	return fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall[i].serviceInstanceGUID
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstanceReturns(result1 []v2action.ServiceInstanceSharedTo, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstanceSharedTosByServiceInstanceStub = nil
	fake.getServiceInstanceSharedTosByServiceInstanceReturns = struct {
		result1 []v2action.ServiceInstanceSharedTo
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstanceReturnsOnCall(i int, result1 []v2action.ServiceInstanceSharedTo, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstanceSharedTosByServiceInstanceStub = nil
	if fake.getServiceInstanceSharedTosByServiceInstanceReturnsOnCall == nil {
		fake.getServiceInstanceSharedTosByServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 []v2action.ServiceInstanceSharedTo
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceInstanceSharedTosByServiceInstanceReturnsOnCall[i] = struct {
		result1 []v2action.ServiceInstanceSharedTo
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndName(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error) {
	fake.getSpaceByOrganizationAndNameMutex.Lock()
	ret, specificReturn := fake.getSpaceByOrganizationAndNameReturnsOnCall[len(fake.getSpaceByOrganizationAndNameArgsForCall)]
	fake.getSpaceByOrganizationAndNameArgsForCall = append(fake.getSpaceByOrganizationAndNameArgsForCall, struct {
		orgGUID   string
		spaceName string
	}{orgGUID, spaceName})
	fake.recordInvocation("GetSpaceByOrganizationAndName", []interface{}{orgGUID, spaceName})
	fake.getSpaceByOrganizationAndNameMutex.Unlock()
	if fake.GetSpaceByOrganizationAndNameStub != nil {
		return fake.GetSpaceByOrganizationAndNameStub(orgGUID, spaceName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getSpaceByOrganizationAndNameReturns.result1, fake.getSpaceByOrganizationAndNameReturns.result2, fake.getSpaceByOrganizationAndNameReturns.result3
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndNameCallCount() int {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return len(fake.getSpaceByOrganizationAndNameArgsForCall)
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndNameArgsForCall(i int) (string, string) {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return fake.getSpaceByOrganizationAndNameArgsForCall[i].orgGUID, fake.getSpaceByOrganizationAndNameArgsForCall[i].spaceName
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndNameReturns(result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	fake.getSpaceByOrganizationAndNameReturns = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndNameReturnsOnCall(i int, result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	if fake.getSpaceByOrganizationAndNameReturnsOnCall == nil {
		fake.getSpaceByOrganizationAndNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Space
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getSpaceByOrganizationAndNameReturnsOnCall[i] = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createManifestApplicationMutex.RLock()
	defer fake.createManifestApplicationMutex.RUnlock()
	fake.getFeatureFlagsMutex.RLock()
	defer fake.getFeatureFlagsMutex.RUnlock()
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceSharedTosByServiceInstanceMutex.RUnlock()
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV2Actor) recordInvocation(key string, args []interface{}) {
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

var _ v2v3action.V2Actor = new(FakeV2Actor)
