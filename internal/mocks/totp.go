// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/v4/internal/totp (interfaces: Provider)
//
// Generated by this command:
//
//	mockgen -package mocks -destination totp.go -mock_names Provider=MockTOTP github.com/authelia/authelia/v4/internal/totp Provider
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/authelia/authelia/v4/internal/model"
	totp "github.com/authelia/authelia/v4/internal/totp"
	gomock "go.uber.org/mock/gomock"
)

// MockTOTP is a mock of Provider interface.
type MockTOTP struct {
	ctrl     *gomock.Controller
	recorder *MockTOTPMockRecorder
	isgomock struct{}
}

// MockTOTPMockRecorder is the mock recorder for MockTOTP.
type MockTOTPMockRecorder struct {
	mock *MockTOTP
}

// NewMockTOTP creates a new mock instance.
func NewMockTOTP(ctrl *gomock.Controller) *MockTOTP {
	mock := &MockTOTP{ctrl: ctrl}
	mock.recorder = &MockTOTPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTOTP) EXPECT() *MockTOTPMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockTOTP) Generate(ctx totp.Context, username string) (*model.TOTPConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx, username)
	ret0, _ := ret[0].(*model.TOTPConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockTOTPMockRecorder) Generate(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockTOTP)(nil).Generate), ctx, username)
}

// GenerateCustom mocks base method.
func (m *MockTOTP) GenerateCustom(ctx totp.Context, username, algorithm, secret string, digits uint32, period, secretSize uint) (*model.TOTPConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCustom", ctx, username, algorithm, secret, digits, period, secretSize)
	ret0, _ := ret[0].(*model.TOTPConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCustom indicates an expected call of GenerateCustom.
func (mr *MockTOTPMockRecorder) GenerateCustom(ctx, username, algorithm, secret, digits, period, secretSize any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCustom", reflect.TypeOf((*MockTOTP)(nil).GenerateCustom), ctx, username, algorithm, secret, digits, period, secretSize)
}

// Options mocks base method.
func (m *MockTOTP) Options() model.TOTPOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(model.TOTPOptions)
	return ret0
}

// Options indicates an expected call of Options.
func (mr *MockTOTPMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockTOTP)(nil).Options))
}

// Validate mocks base method.
func (m *MockTOTP) Validate(ctx totp.Context, token string, config *model.TOTPConfiguration) (bool, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, token, config)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Validate indicates an expected call of Validate.
func (mr *MockTOTPMockRecorder) Validate(ctx, token, config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTOTP)(nil).Validate), ctx, token, config)
}
