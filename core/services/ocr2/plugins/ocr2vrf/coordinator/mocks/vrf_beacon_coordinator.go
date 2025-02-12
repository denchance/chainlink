// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	mock "github.com/stretchr/testify/mock"
)

// VRFBeaconCoordinator is an autogenerated mock type for the VRFBeaconCoordinator type
type VRFBeaconCoordinator struct {
	mock.Mock
}

// IBeaconPeriodBlocks provides a mock function with given fields: opts
func (_m *VRFBeaconCoordinator) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SProvingKeyHash provides a mock function with given fields: opts
func (_m *VRFBeaconCoordinator) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	ret := _m.Called(opts)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) [32]byte); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewVRFBeaconCoordinator interface {
	mock.TestingT
	Cleanup(func())
}

// NewVRFBeaconCoordinator creates a new instance of VRFBeaconCoordinator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVRFBeaconCoordinator(t mockConstructorTestingTNewVRFBeaconCoordinator) *VRFBeaconCoordinator {
	mock := &VRFBeaconCoordinator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
