// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/cloudfoundry-incubator/tcp-emitter/routing_table"
)

type FakeRoutingTable struct {
	RouteCountStub        func() int
	routeCountMutex       sync.RWMutex
	routeCountArgsForCall []struct{}
	routeCountReturns struct {
		result1 int
	}
	AddRoutesStub        func(desiredLRP *models.DesiredLRP) routing_table.RoutingEvents
	addRoutesMutex       sync.RWMutex
	addRoutesArgsForCall []struct {
		desiredLRP *models.DesiredLRP
	}
	addRoutesReturns struct {
		result1 routing_table.RoutingEvents
	}
	UpdateRoutesStub        func(beforeLRP, afterLRP *models.DesiredLRP) routing_table.RoutingEvents
	updateRoutesMutex       sync.RWMutex
	updateRoutesArgsForCall []struct {
		beforeLRP *models.DesiredLRP
		afterLRP  *models.DesiredLRP
	}
	updateRoutesReturns struct {
		result1 routing_table.RoutingEvents
	}
	RemoveRoutesStub        func(desiredLRP *models.DesiredLRP) routing_table.RoutingEvents
	removeRoutesMutex       sync.RWMutex
	removeRoutesArgsForCall []struct {
		desiredLRP *models.DesiredLRP
	}
	removeRoutesReturns struct {
		result1 routing_table.RoutingEvents
	}
	AddEndpointStub        func(actualLRP *models.ActualLRPGroup) routing_table.RoutingEvents
	addEndpointMutex       sync.RWMutex
	addEndpointArgsForCall []struct {
		actualLRP *models.ActualLRPGroup
	}
	addEndpointReturns struct {
		result1 routing_table.RoutingEvents
	}
	RemoveEndpointStub        func(actualLRP *models.ActualLRPGroup) routing_table.RoutingEvents
	removeEndpointMutex       sync.RWMutex
	removeEndpointArgsForCall []struct {
		actualLRP *models.ActualLRPGroup
	}
	removeEndpointReturns struct {
		result1 routing_table.RoutingEvents
	}
	SwapStub        func(t routing_table.RoutingTable) routing_table.RoutingEvents
	swapMutex       sync.RWMutex
	swapArgsForCall []struct {
		t routing_table.RoutingTable
	}
	swapReturns struct {
		result1 routing_table.RoutingEvents
	}
	GetRoutingEventsStub        func() routing_table.RoutingEvents
	getRoutingEventsMutex       sync.RWMutex
	getRoutingEventsArgsForCall []struct{}
	getRoutingEventsReturns struct {
		result1 routing_table.RoutingEvents
	}
}

func (fake *FakeRoutingTable) RouteCount() int {
	fake.routeCountMutex.Lock()
	fake.routeCountArgsForCall = append(fake.routeCountArgsForCall, struct{}{})
	fake.routeCountMutex.Unlock()
	if fake.RouteCountStub != nil {
		return fake.RouteCountStub()
	} else {
		return fake.routeCountReturns.result1
	}
}

func (fake *FakeRoutingTable) RouteCountCallCount() int {
	fake.routeCountMutex.RLock()
	defer fake.routeCountMutex.RUnlock()
	return len(fake.routeCountArgsForCall)
}

func (fake *FakeRoutingTable) RouteCountReturns(result1 int) {
	fake.RouteCountStub = nil
	fake.routeCountReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeRoutingTable) AddRoutes(desiredLRP *models.DesiredLRP) routing_table.RoutingEvents {
	fake.addRoutesMutex.Lock()
	fake.addRoutesArgsForCall = append(fake.addRoutesArgsForCall, struct {
		desiredLRP *models.DesiredLRP
	}{desiredLRP})
	fake.addRoutesMutex.Unlock()
	if fake.AddRoutesStub != nil {
		return fake.AddRoutesStub(desiredLRP)
	} else {
		return fake.addRoutesReturns.result1
	}
}

func (fake *FakeRoutingTable) AddRoutesCallCount() int {
	fake.addRoutesMutex.RLock()
	defer fake.addRoutesMutex.RUnlock()
	return len(fake.addRoutesArgsForCall)
}

func (fake *FakeRoutingTable) AddRoutesArgsForCall(i int) *models.DesiredLRP {
	fake.addRoutesMutex.RLock()
	defer fake.addRoutesMutex.RUnlock()
	return fake.addRoutesArgsForCall[i].desiredLRP
}

func (fake *FakeRoutingTable) AddRoutesReturns(result1 routing_table.RoutingEvents) {
	fake.AddRoutesStub = nil
	fake.addRoutesReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) UpdateRoutes(beforeLRP *models.DesiredLRP, afterLRP *models.DesiredLRP) routing_table.RoutingEvents {
	fake.updateRoutesMutex.Lock()
	fake.updateRoutesArgsForCall = append(fake.updateRoutesArgsForCall, struct {
		beforeLRP *models.DesiredLRP
		afterLRP  *models.DesiredLRP
	}{beforeLRP, afterLRP})
	fake.updateRoutesMutex.Unlock()
	if fake.UpdateRoutesStub != nil {
		return fake.UpdateRoutesStub(beforeLRP, afterLRP)
	} else {
		return fake.updateRoutesReturns.result1
	}
}

func (fake *FakeRoutingTable) UpdateRoutesCallCount() int {
	fake.updateRoutesMutex.RLock()
	defer fake.updateRoutesMutex.RUnlock()
	return len(fake.updateRoutesArgsForCall)
}

func (fake *FakeRoutingTable) UpdateRoutesArgsForCall(i int) (*models.DesiredLRP, *models.DesiredLRP) {
	fake.updateRoutesMutex.RLock()
	defer fake.updateRoutesMutex.RUnlock()
	return fake.updateRoutesArgsForCall[i].beforeLRP, fake.updateRoutesArgsForCall[i].afterLRP
}

func (fake *FakeRoutingTable) UpdateRoutesReturns(result1 routing_table.RoutingEvents) {
	fake.UpdateRoutesStub = nil
	fake.updateRoutesReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) RemoveRoutes(desiredLRP *models.DesiredLRP) routing_table.RoutingEvents {
	fake.removeRoutesMutex.Lock()
	fake.removeRoutesArgsForCall = append(fake.removeRoutesArgsForCall, struct {
		desiredLRP *models.DesiredLRP
	}{desiredLRP})
	fake.removeRoutesMutex.Unlock()
	if fake.RemoveRoutesStub != nil {
		return fake.RemoveRoutesStub(desiredLRP)
	} else {
		return fake.removeRoutesReturns.result1
	}
}

func (fake *FakeRoutingTable) RemoveRoutesCallCount() int {
	fake.removeRoutesMutex.RLock()
	defer fake.removeRoutesMutex.RUnlock()
	return len(fake.removeRoutesArgsForCall)
}

func (fake *FakeRoutingTable) RemoveRoutesArgsForCall(i int) *models.DesiredLRP {
	fake.removeRoutesMutex.RLock()
	defer fake.removeRoutesMutex.RUnlock()
	return fake.removeRoutesArgsForCall[i].desiredLRP
}

func (fake *FakeRoutingTable) RemoveRoutesReturns(result1 routing_table.RoutingEvents) {
	fake.RemoveRoutesStub = nil
	fake.removeRoutesReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) AddEndpoint(actualLRP *models.ActualLRPGroup) routing_table.RoutingEvents {
	fake.addEndpointMutex.Lock()
	fake.addEndpointArgsForCall = append(fake.addEndpointArgsForCall, struct {
		actualLRP *models.ActualLRPGroup
	}{actualLRP})
	fake.addEndpointMutex.Unlock()
	if fake.AddEndpointStub != nil {
		return fake.AddEndpointStub(actualLRP)
	} else {
		return fake.addEndpointReturns.result1
	}
}

func (fake *FakeRoutingTable) AddEndpointCallCount() int {
	fake.addEndpointMutex.RLock()
	defer fake.addEndpointMutex.RUnlock()
	return len(fake.addEndpointArgsForCall)
}

func (fake *FakeRoutingTable) AddEndpointArgsForCall(i int) *models.ActualLRPGroup {
	fake.addEndpointMutex.RLock()
	defer fake.addEndpointMutex.RUnlock()
	return fake.addEndpointArgsForCall[i].actualLRP
}

func (fake *FakeRoutingTable) AddEndpointReturns(result1 routing_table.RoutingEvents) {
	fake.AddEndpointStub = nil
	fake.addEndpointReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) RemoveEndpoint(actualLRP *models.ActualLRPGroup) routing_table.RoutingEvents {
	fake.removeEndpointMutex.Lock()
	fake.removeEndpointArgsForCall = append(fake.removeEndpointArgsForCall, struct {
		actualLRP *models.ActualLRPGroup
	}{actualLRP})
	fake.removeEndpointMutex.Unlock()
	if fake.RemoveEndpointStub != nil {
		return fake.RemoveEndpointStub(actualLRP)
	} else {
		return fake.removeEndpointReturns.result1
	}
}

func (fake *FakeRoutingTable) RemoveEndpointCallCount() int {
	fake.removeEndpointMutex.RLock()
	defer fake.removeEndpointMutex.RUnlock()
	return len(fake.removeEndpointArgsForCall)
}

func (fake *FakeRoutingTable) RemoveEndpointArgsForCall(i int) *models.ActualLRPGroup {
	fake.removeEndpointMutex.RLock()
	defer fake.removeEndpointMutex.RUnlock()
	return fake.removeEndpointArgsForCall[i].actualLRP
}

func (fake *FakeRoutingTable) RemoveEndpointReturns(result1 routing_table.RoutingEvents) {
	fake.RemoveEndpointStub = nil
	fake.removeEndpointReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) Swap(t routing_table.RoutingTable) routing_table.RoutingEvents {
	fake.swapMutex.Lock()
	fake.swapArgsForCall = append(fake.swapArgsForCall, struct {
		t routing_table.RoutingTable
	}{t})
	fake.swapMutex.Unlock()
	if fake.SwapStub != nil {
		return fake.SwapStub(t)
	} else {
		return fake.swapReturns.result1
	}
}

func (fake *FakeRoutingTable) SwapCallCount() int {
	fake.swapMutex.RLock()
	defer fake.swapMutex.RUnlock()
	return len(fake.swapArgsForCall)
}

func (fake *FakeRoutingTable) SwapArgsForCall(i int) routing_table.RoutingTable {
	fake.swapMutex.RLock()
	defer fake.swapMutex.RUnlock()
	return fake.swapArgsForCall[i].t
}

func (fake *FakeRoutingTable) SwapReturns(result1 routing_table.RoutingEvents) {
	fake.SwapStub = nil
	fake.swapReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) GetRoutingEvents() routing_table.RoutingEvents {
	fake.getRoutingEventsMutex.Lock()
	fake.getRoutingEventsArgsForCall = append(fake.getRoutingEventsArgsForCall, struct{}{})
	fake.getRoutingEventsMutex.Unlock()
	if fake.GetRoutingEventsStub != nil {
		return fake.GetRoutingEventsStub()
	} else {
		return fake.getRoutingEventsReturns.result1
	}
}

func (fake *FakeRoutingTable) GetRoutingEventsCallCount() int {
	fake.getRoutingEventsMutex.RLock()
	defer fake.getRoutingEventsMutex.RUnlock()
	return len(fake.getRoutingEventsArgsForCall)
}

func (fake *FakeRoutingTable) GetRoutingEventsReturns(result1 routing_table.RoutingEvents) {
	fake.GetRoutingEventsStub = nil
	fake.getRoutingEventsReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

var _ routing_table.RoutingTable = new(FakeRoutingTable)
