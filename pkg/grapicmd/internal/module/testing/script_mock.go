// Code generated by MockGen. DO NOT EDIT.
// Source: script.go

// Package moduletesting is a generated GoMock package.
package moduletesting

import (
	context "context"
	module "github.com/Ryo-not-rio/grapi/pkg/grapicmd/internal/module"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockScript is a mock of Script interface
type MockScript struct {
	ctrl     *gomock.Controller
	recorder *MockScriptMockRecorder
}

// MockScriptMockRecorder is the mock recorder for MockScript
type MockScriptMockRecorder struct {
	mock *MockScript
}

// NewMockScript creates a new mock instance
func NewMockScript(ctrl *gomock.Controller) *MockScript {
	mock := &MockScript{ctrl: ctrl}
	mock.recorder = &MockScriptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScript) EXPECT() *MockScriptMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockScript) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockScriptMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockScript)(nil).Name))
}

// Build mocks base method
func (m *MockScript) Build(ctx context.Context, args ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Build", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build
func (mr *MockScriptMockRecorder) Build(ctx interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockScript)(nil).Build), varargs...)
}

// Run mocks base method
func (m *MockScript) Run(ctx context.Context, args ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockScriptMockRecorder) Run(ctx interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockScript)(nil).Run), varargs...)
}

// MockScriptLoader is a mock of ScriptLoader interface
type MockScriptLoader struct {
	ctrl     *gomock.Controller
	recorder *MockScriptLoaderMockRecorder
}

// MockScriptLoaderMockRecorder is the mock recorder for MockScriptLoader
type MockScriptLoaderMockRecorder struct {
	mock *MockScriptLoader
}

// NewMockScriptLoader creates a new mock instance
func NewMockScriptLoader(ctrl *gomock.Controller) *MockScriptLoader {
	mock := &MockScriptLoader{ctrl: ctrl}
	mock.recorder = &MockScriptLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScriptLoader) EXPECT() *MockScriptLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockScriptLoader) Load(dir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load
func (mr *MockScriptLoaderMockRecorder) Load(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockScriptLoader)(nil).Load), dir)
}

// Get mocks base method
func (m *MockScriptLoader) Get(name string) (module.Script, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(module.Script)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockScriptLoaderMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockScriptLoader)(nil).Get), name)
}

// Names mocks base method
func (m *MockScriptLoader) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names
func (mr *MockScriptLoaderMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockScriptLoader)(nil).Names))
}
