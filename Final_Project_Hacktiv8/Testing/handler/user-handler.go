package mock_handler

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserHandlerIface is a mock of UserHandlerIface interface.
type MockUserHandlerIface struct {
	ctrl     *gomock.Controller
	recorder *MockUserHandlerIfaceMockRecorder
}

// MockUserHandlerIfaceMockRecorder is the mock recorder for MockUserHandlerIface.
type MockUserHandlerIfaceMockRecorder struct {
	mock *MockUserHandlerIface
}

// NewMockUserHandlerIface creates a new mock instance.
func NewMockUserHandlerIface(ctrl *gomock.Controller) *MockUserHandlerIface {
	mock := &MockUserHandlerIface{ctrl: ctrl}
	mock.recorder = &MockUserHandlerIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserHandlerIface) EXPECT() *MockUserHandlerIfaceMockRecorder {
	return m.recorder
}

// UserHandler mocks base method.
func (m *MockUserHandlerIface) UserHandler(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserHandler", w, r)
}

// UserHandler indicates an expected call of UserHandler.
func (mr *MockUserHandlerIfaceMockRecorder) UserHandler(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHandler", reflect.TypeOf((*MockUserHandlerIface)(nil).UserHandler), w, r)
}
