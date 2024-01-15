// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prometheus/client_golang/prometheus (interfaces: Gauge)
//
// Generated by this command:
//
//	mockgen -destination=mock_gauge_test.go -package state github.com/prometheus/client_golang/prometheus Gauge
//

// Package state is a generated GoMock package.
package state

import (
	reflect "reflect"

	prometheus "github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	gomock "go.uber.org/mock/gomock"
)

// MockGauge is a mock of Gauge interface.
type MockGauge struct {
	ctrl     *gomock.Controller
	recorder *MockGaugeMockRecorder
}

// MockGaugeMockRecorder is the mock recorder for MockGauge.
type MockGaugeMockRecorder struct {
	mock *MockGauge
}

// NewMockGauge creates a new mock instance.
func NewMockGauge(ctrl *gomock.Controller) *MockGauge {
	mock := &MockGauge{ctrl: ctrl}
	mock.recorder = &MockGaugeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGauge) EXPECT() *MockGaugeMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockGauge) Add(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockGaugeMockRecorder) Add(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockGauge)(nil).Add), arg0)
}

// Collect mocks base method.
func (m *MockGauge) Collect(arg0 chan<- prometheus.Metric) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Collect", arg0)
}

// Collect indicates an expected call of Collect.
func (mr *MockGaugeMockRecorder) Collect(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockGauge)(nil).Collect), arg0)
}

// Dec mocks base method.
func (m *MockGauge) Dec() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dec")
}

// Dec indicates an expected call of Dec.
func (mr *MockGaugeMockRecorder) Dec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dec", reflect.TypeOf((*MockGauge)(nil).Dec))
}

// Desc mocks base method.
func (m *MockGauge) Desc() *prometheus.Desc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Desc")
	ret0, _ := ret[0].(*prometheus.Desc)
	return ret0
}

// Desc indicates an expected call of Desc.
func (mr *MockGaugeMockRecorder) Desc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Desc", reflect.TypeOf((*MockGauge)(nil).Desc))
}

// Describe mocks base method.
func (m *MockGauge) Describe(arg0 chan<- *prometheus.Desc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Describe", arg0)
}

// Describe indicates an expected call of Describe.
func (mr *MockGaugeMockRecorder) Describe(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockGauge)(nil).Describe), arg0)
}

// Inc mocks base method.
func (m *MockGauge) Inc() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Inc")
}

// Inc indicates an expected call of Inc.
func (mr *MockGaugeMockRecorder) Inc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inc", reflect.TypeOf((*MockGauge)(nil).Inc))
}

// Set mocks base method.
func (m *MockGauge) Set(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0)
}

// Set indicates an expected call of Set.
func (mr *MockGaugeMockRecorder) Set(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockGauge)(nil).Set), arg0)
}

// SetToCurrentTime mocks base method.
func (m *MockGauge) SetToCurrentTime() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToCurrentTime")
}

// SetToCurrentTime indicates an expected call of SetToCurrentTime.
func (mr *MockGaugeMockRecorder) SetToCurrentTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToCurrentTime", reflect.TypeOf((*MockGauge)(nil).SetToCurrentTime))
}

// Sub mocks base method.
func (m *MockGauge) Sub(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sub", arg0)
}

// Sub indicates an expected call of Sub.
func (mr *MockGaugeMockRecorder) Sub(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockGauge)(nil).Sub), arg0)
}

// Write mocks base method.
func (m *MockGauge) Write(arg0 *io_prometheus_client.Metric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockGaugeMockRecorder) Write(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockGauge)(nil).Write), arg0)
}
