// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/compute-image-import/cli_tools/gce_ovf_export/domain (interfaces: OvfManifestGenerator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOvfManifestGenerator is a mock of OvfManifestGenerator interface.
type MockOvfManifestGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockOvfManifestGeneratorMockRecorder
}

// MockOvfManifestGeneratorMockRecorder is the mock recorder for MockOvfManifestGenerator.
type MockOvfManifestGeneratorMockRecorder struct {
	mock *MockOvfManifestGenerator
}

// NewMockOvfManifestGenerator creates a new mock instance.
func NewMockOvfManifestGenerator(ctrl *gomock.Controller) *MockOvfManifestGenerator {
	mock := &MockOvfManifestGenerator{ctrl: ctrl}
	mock.recorder = &MockOvfManifestGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOvfManifestGenerator) EXPECT() *MockOvfManifestGeneratorMockRecorder {
	return m.recorder
}

// Cancel mocks base method.
func (m *MockOvfManifestGenerator) Cancel(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockOvfManifestGeneratorMockRecorder) Cancel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockOvfManifestGenerator)(nil).Cancel), arg0)
}

// GenerateAndWriteToGCS mocks base method.
func (m *MockOvfManifestGenerator) GenerateAndWriteToGCS(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAndWriteToGCS", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateAndWriteToGCS indicates an expected call of GenerateAndWriteToGCS.
func (mr *MockOvfManifestGeneratorMockRecorder) GenerateAndWriteToGCS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAndWriteToGCS", reflect.TypeOf((*MockOvfManifestGenerator)(nil).GenerateAndWriteToGCS), arg0, arg1)
}
