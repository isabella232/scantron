// Code generated by MockGen. DO NOT EDIT.
// Source: bosh/bosh.go

// Package bosh is a generated GoMock package.
package bosh

import (
	director "github.com/cloudfoundry/bosh-cli/director"
	gomock "github.com/golang/mock/gomock"
	remotemachine "github.com/pivotal-cf/scantron/remotemachine"
	reflect "reflect"
)

// MockTargetDeployment is a mock of TargetDeployment interface
type MockTargetDeployment struct {
	ctrl     *gomock.Controller
	recorder *MockTargetDeploymentMockRecorder
}

// MockTargetDeploymentMockRecorder is the mock recorder for MockTargetDeployment
type MockTargetDeploymentMockRecorder struct {
	mock *MockTargetDeployment
}

// NewMockTargetDeployment creates a new mock instance
func NewMockTargetDeployment(ctrl *gomock.Controller) *MockTargetDeployment {
	mock := &MockTargetDeployment{ctrl: ctrl}
	mock.recorder = &MockTargetDeploymentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTargetDeployment) EXPECT() *MockTargetDeploymentMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockTargetDeployment) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockTargetDeploymentMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockTargetDeployment)(nil).Name))
}

// VMs mocks base method
func (m *MockTargetDeployment) VMs() []director.VMInfo {
	ret := m.ctrl.Call(m, "VMs")
	ret0, _ := ret[0].([]director.VMInfo)
	return ret0
}

// VMs indicates an expected call of VMs
func (mr *MockTargetDeploymentMockRecorder) VMs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMs", reflect.TypeOf((*MockTargetDeployment)(nil).VMs))
}

// Releases mocks base method
func (m *MockTargetDeployment) Releases() []director.Release {
	ret := m.ctrl.Call(m, "Releases")
	ret0, _ := ret[0].([]director.Release)
	return ret0
}

// Releases indicates an expected call of Releases
func (mr *MockTargetDeploymentMockRecorder) Releases() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Releases", reflect.TypeOf((*MockTargetDeployment)(nil).Releases))
}

// Setup mocks base method
func (m *MockTargetDeployment) Setup() error {
	ret := m.ctrl.Call(m, "Setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup
func (mr *MockTargetDeploymentMockRecorder) Setup() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockTargetDeployment)(nil).Setup))
}

// ConnectTo mocks base method
func (m *MockTargetDeployment) ConnectTo(arg0 director.VMInfo) remotemachine.RemoteMachine {
	ret := m.ctrl.Call(m, "ConnectTo", arg0)
	ret0, _ := ret[0].(remotemachine.RemoteMachine)
	return ret0
}

// ConnectTo indicates an expected call of ConnectTo
func (mr *MockTargetDeploymentMockRecorder) ConnectTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectTo", reflect.TypeOf((*MockTargetDeployment)(nil).ConnectTo), arg0)
}

// Cleanup mocks base method
func (m *MockTargetDeployment) Cleanup() error {
	ret := m.ctrl.Call(m, "Cleanup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Cleanup indicates an expected call of Cleanup
func (mr *MockTargetDeploymentMockRecorder) Cleanup() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockTargetDeployment)(nil).Cleanup))
}
