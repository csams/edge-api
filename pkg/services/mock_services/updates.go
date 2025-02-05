// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/updates.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/redhatinsights/edge-api/pkg/models"
)

// MockUpdateServiceInterface is a mock of UpdateServiceInterface interface.
type MockUpdateServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateServiceInterfaceMockRecorder
}

// MockUpdateServiceInterfaceMockRecorder is the mock recorder for MockUpdateServiceInterface.
type MockUpdateServiceInterfaceMockRecorder struct {
	mock *MockUpdateServiceInterface
}

// NewMockUpdateServiceInterface creates a new mock instance.
func NewMockUpdateServiceInterface(ctrl *gomock.Controller) *MockUpdateServiceInterface {
	mock := &MockUpdateServiceInterface{ctrl: ctrl}
	mock.recorder = &MockUpdateServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateServiceInterface) EXPECT() *MockUpdateServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateUpdate mocks base method.
func (m *MockUpdateServiceInterface) CreateUpdate(id uint) (*models.UpdateTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpdate", id)
	ret0, _ := ret[0].(*models.UpdateTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUpdate indicates an expected call of CreateUpdate.
func (mr *MockUpdateServiceInterfaceMockRecorder) CreateUpdate(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpdate", reflect.TypeOf((*MockUpdateServiceInterface)(nil).CreateUpdate), id)
}

// GetUpdatePlaybook mocks base method.
func (m *MockUpdateServiceInterface) GetUpdatePlaybook(update *models.UpdateTransaction) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdatePlaybook", update)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdatePlaybook indicates an expected call of GetUpdatePlaybook.
func (mr *MockUpdateServiceInterfaceMockRecorder) GetUpdatePlaybook(update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdatePlaybook", reflect.TypeOf((*MockUpdateServiceInterface)(nil).GetUpdatePlaybook), update)
}

// GetUpdateTransactionsForDevice mocks base method.
func (m *MockUpdateServiceInterface) GetUpdateTransactionsForDevice(device *models.Device) (*[]models.UpdateTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateTransactionsForDevice", device)
	ret0, _ := ret[0].(*[]models.UpdateTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdateTransactionsForDevice indicates an expected call of GetUpdateTransactionsForDevice.
func (mr *MockUpdateServiceInterfaceMockRecorder) GetUpdateTransactionsForDevice(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateTransactionsForDevice", reflect.TypeOf((*MockUpdateServiceInterface)(nil).GetUpdateTransactionsForDevice), device)
}

// ProcessPlaybookDispatcherRunEvent mocks base method.
func (m *MockUpdateServiceInterface) ProcessPlaybookDispatcherRunEvent(message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPlaybookDispatcherRunEvent", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessPlaybookDispatcherRunEvent indicates an expected call of ProcessPlaybookDispatcherRunEvent.
func (mr *MockUpdateServiceInterfaceMockRecorder) ProcessPlaybookDispatcherRunEvent(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPlaybookDispatcherRunEvent", reflect.TypeOf((*MockUpdateServiceInterface)(nil).ProcessPlaybookDispatcherRunEvent), message)
}
