// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient/insights (interfaces: ActivityLogsClient)

// Package mock_insights is a generated GoMock package.
package mock_insights

import (
	context "context"
	reflect "reflect"

	insights "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2018-09-01/insights"
	gomock "github.com/golang/mock/gomock"
)

// MockActivityLogsClient is a mock of ActivityLogsClient interface
type MockActivityLogsClient struct {
	ctrl     *gomock.Controller
	recorder *MockActivityLogsClientMockRecorder
}

// MockActivityLogsClientMockRecorder is the mock recorder for MockActivityLogsClient
type MockActivityLogsClientMockRecorder struct {
	mock *MockActivityLogsClient
}

// NewMockActivityLogsClient creates a new mock instance
func NewMockActivityLogsClient(ctrl *gomock.Controller) *MockActivityLogsClient {
	mock := &MockActivityLogsClient{ctrl: ctrl}
	mock.recorder = &MockActivityLogsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActivityLogsClient) EXPECT() *MockActivityLogsClientMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockActivityLogsClient) List(arg0 context.Context, arg1, arg2 string) (insights.EventDataCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(insights.EventDataCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockActivityLogsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockActivityLogsClient)(nil).List), arg0, arg1, arg2)
}