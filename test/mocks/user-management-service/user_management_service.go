// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influenzanet/user-management-service/pkg/api (interfaces: UserManagementApiClient,UserManagementApi_StreamUsersClient)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/influenzanet/user-management-service/pkg/api"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockUserManagementApiClient is a mock of UserManagementApiClient interface
type MockUserManagementApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserManagementApiClientMockRecorder
}

// MockUserManagementApiClientMockRecorder is the mock recorder for MockUserManagementApiClient
type MockUserManagementApiClientMockRecorder struct {
	mock *MockUserManagementApiClient
}

// NewMockUserManagementApiClient creates a new mock instance
func NewMockUserManagementApiClient(ctrl *gomock.Controller) *MockUserManagementApiClient {
	mock := &MockUserManagementApiClient{ctrl: ctrl}
	mock.recorder = &MockUserManagementApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserManagementApiClient) EXPECT() *MockUserManagementApiClientMockRecorder {
	return m.recorder
}

// AddEmail mocks base method
func (m *MockUserManagementApiClient) AddEmail(arg0 context.Context, arg1 *api.ContactInfoMsg, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEmail", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddEmail indicates an expected call of AddEmail
func (mr *MockUserManagementApiClientMockRecorder) AddEmail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEmail", reflect.TypeOf((*MockUserManagementApiClient)(nil).AddEmail), varargs...)
}

// AddRoleForUser mocks base method
func (m *MockUserManagementApiClient) AddRoleForUser(arg0 context.Context, arg1 *api.RoleMsg, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddRoleForUser", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRoleForUser indicates an expected call of AddRoleForUser
func (mr *MockUserManagementApiClientMockRecorder) AddRoleForUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoleForUser", reflect.TypeOf((*MockUserManagementApiClient)(nil).AddRoleForUser), varargs...)
}

// ChangeAccountIDEmail mocks base method
func (m *MockUserManagementApiClient) ChangeAccountIDEmail(arg0 context.Context, arg1 *api.EmailChangeMsg, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeAccountIDEmail", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeAccountIDEmail indicates an expected call of ChangeAccountIDEmail
func (mr *MockUserManagementApiClientMockRecorder) ChangeAccountIDEmail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAccountIDEmail", reflect.TypeOf((*MockUserManagementApiClient)(nil).ChangeAccountIDEmail), varargs...)
}

// ChangePassword mocks base method
func (m *MockUserManagementApiClient) ChangePassword(arg0 context.Context, arg1 *api.PasswordChangeMsg, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangePassword", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword
func (mr *MockUserManagementApiClientMockRecorder) ChangePassword(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserManagementApiClient)(nil).ChangePassword), varargs...)
}

// ChangePreferredLanguage mocks base method
func (m *MockUserManagementApiClient) ChangePreferredLanguage(arg0 context.Context, arg1 *api.LanguageChangeMsg, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangePreferredLanguage", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePreferredLanguage indicates an expected call of ChangePreferredLanguage
func (mr *MockUserManagementApiClientMockRecorder) ChangePreferredLanguage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePreferredLanguage", reflect.TypeOf((*MockUserManagementApiClient)(nil).ChangePreferredLanguage), varargs...)
}

// CreateUser mocks base method
func (m *MockUserManagementApiClient) CreateUser(arg0 context.Context, arg1 *api.CreateUserReq, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUser", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserManagementApiClientMockRecorder) CreateUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserManagementApiClient)(nil).CreateUser), varargs...)
}

// DeleteAccount mocks base method
func (m *MockUserManagementApiClient) DeleteAccount(arg0 context.Context, arg1 *api.UserReference, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAccount", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccount indicates an expected call of DeleteAccount
func (mr *MockUserManagementApiClientMockRecorder) DeleteAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockUserManagementApiClient)(nil).DeleteAccount), varargs...)
}

// DeleteTempToken mocks base method
func (m *MockUserManagementApiClient) DeleteTempToken(arg0 context.Context, arg1 *api.TempToken, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTempToken", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTempToken indicates an expected call of DeleteTempToken
func (mr *MockUserManagementApiClientMockRecorder) DeleteTempToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTempToken", reflect.TypeOf((*MockUserManagementApiClient)(nil).DeleteTempToken), varargs...)
}

// FindNonParticipantUsers mocks base method
func (m *MockUserManagementApiClient) FindNonParticipantUsers(arg0 context.Context, arg1 *api.FindNonParticipantUsersMsg, arg2 ...grpc.CallOption) (*api.UserListMsg, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindNonParticipantUsers", varargs...)
	ret0, _ := ret[0].(*api.UserListMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNonParticipantUsers indicates an expected call of FindNonParticipantUsers
func (mr *MockUserManagementApiClientMockRecorder) FindNonParticipantUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNonParticipantUsers", reflect.TypeOf((*MockUserManagementApiClient)(nil).FindNonParticipantUsers), varargs...)
}

// GenerateTempToken mocks base method
func (m *MockUserManagementApiClient) GenerateTempToken(arg0 context.Context, arg1 *api.TempTokenInfo, arg2 ...grpc.CallOption) (*api.TempToken, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GenerateTempToken", varargs...)
	ret0, _ := ret[0].(*api.TempToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTempToken indicates an expected call of GenerateTempToken
func (mr *MockUserManagementApiClientMockRecorder) GenerateTempToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTempToken", reflect.TypeOf((*MockUserManagementApiClient)(nil).GenerateTempToken), varargs...)
}

// GetInfosForPasswordReset mocks base method
func (m *MockUserManagementApiClient) GetInfosForPasswordReset(arg0 context.Context, arg1 *api.GetInfosForResetPasswordMsg, arg2 ...grpc.CallOption) (*api.UserInfoForPWReset, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInfosForPasswordReset", varargs...)
	ret0, _ := ret[0].(*api.UserInfoForPWReset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfosForPasswordReset indicates an expected call of GetInfosForPasswordReset
func (mr *MockUserManagementApiClientMockRecorder) GetInfosForPasswordReset(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfosForPasswordReset", reflect.TypeOf((*MockUserManagementApiClient)(nil).GetInfosForPasswordReset), varargs...)
}

// GetTempTokens mocks base method
func (m *MockUserManagementApiClient) GetTempTokens(arg0 context.Context, arg1 *api.TempTokenInfo, arg2 ...grpc.CallOption) (*api.TempTokenInfos, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTempTokens", varargs...)
	ret0, _ := ret[0].(*api.TempTokenInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTempTokens indicates an expected call of GetTempTokens
func (mr *MockUserManagementApiClientMockRecorder) GetTempTokens(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTempTokens", reflect.TypeOf((*MockUserManagementApiClient)(nil).GetTempTokens), varargs...)
}

// GetUser mocks base method
func (m *MockUserManagementApiClient) GetUser(arg0 context.Context, arg1 *api.UserReference, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserManagementApiClientMockRecorder) GetUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserManagementApiClient)(nil).GetUser), varargs...)
}

// InitiatePasswordReset mocks base method
func (m *MockUserManagementApiClient) InitiatePasswordReset(arg0 context.Context, arg1 *api.InitiateResetPasswordMsg, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InitiatePasswordReset", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitiatePasswordReset indicates an expected call of InitiatePasswordReset
func (mr *MockUserManagementApiClientMockRecorder) InitiatePasswordReset(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiatePasswordReset", reflect.TypeOf((*MockUserManagementApiClient)(nil).InitiatePasswordReset), varargs...)
}

// LoginWithEmail mocks base method
func (m *MockUserManagementApiClient) LoginWithEmail(arg0 context.Context, arg1 *api.LoginWithEmailMsg, arg2 ...grpc.CallOption) (*api.TokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoginWithEmail", varargs...)
	ret0, _ := ret[0].(*api.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginWithEmail indicates an expected call of LoginWithEmail
func (mr *MockUserManagementApiClientMockRecorder) LoginWithEmail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWithEmail", reflect.TypeOf((*MockUserManagementApiClient)(nil).LoginWithEmail), varargs...)
}

// LoginWithTempToken mocks base method
func (m *MockUserManagementApiClient) LoginWithTempToken(arg0 context.Context, arg1 *api.JWTRequest, arg2 ...grpc.CallOption) (*api.TokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoginWithTempToken", varargs...)
	ret0, _ := ret[0].(*api.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginWithTempToken indicates an expected call of LoginWithTempToken
func (mr *MockUserManagementApiClientMockRecorder) LoginWithTempToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWithTempToken", reflect.TypeOf((*MockUserManagementApiClient)(nil).LoginWithTempToken), varargs...)
}

// PurgeUserTempTokens mocks base method
func (m *MockUserManagementApiClient) PurgeUserTempTokens(arg0 context.Context, arg1 *api.TempTokenInfo, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PurgeUserTempTokens", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurgeUserTempTokens indicates an expected call of PurgeUserTempTokens
func (mr *MockUserManagementApiClientMockRecorder) PurgeUserTempTokens(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeUserTempTokens", reflect.TypeOf((*MockUserManagementApiClient)(nil).PurgeUserTempTokens), varargs...)
}

// RemoveEmail mocks base method
func (m *MockUserManagementApiClient) RemoveEmail(arg0 context.Context, arg1 *api.ContactInfoMsg, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveEmail", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveEmail indicates an expected call of RemoveEmail
func (mr *MockUserManagementApiClientMockRecorder) RemoveEmail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEmail", reflect.TypeOf((*MockUserManagementApiClient)(nil).RemoveEmail), varargs...)
}

// RemoveProfile mocks base method
func (m *MockUserManagementApiClient) RemoveProfile(arg0 context.Context, arg1 *api.ProfileRequest, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveProfile", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveProfile indicates an expected call of RemoveProfile
func (mr *MockUserManagementApiClientMockRecorder) RemoveProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProfile", reflect.TypeOf((*MockUserManagementApiClient)(nil).RemoveProfile), varargs...)
}

// RemoveRoleForUser mocks base method
func (m *MockUserManagementApiClient) RemoveRoleForUser(arg0 context.Context, arg1 *api.RoleMsg, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveRoleForUser", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveRoleForUser indicates an expected call of RemoveRoleForUser
func (mr *MockUserManagementApiClientMockRecorder) RemoveRoleForUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRoleForUser", reflect.TypeOf((*MockUserManagementApiClient)(nil).RemoveRoleForUser), varargs...)
}

// RenewJWT mocks base method
func (m *MockUserManagementApiClient) RenewJWT(arg0 context.Context, arg1 *api.RefreshJWTRequest, arg2 ...grpc.CallOption) (*api.TokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenewJWT", varargs...)
	ret0, _ := ret[0].(*api.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewJWT indicates an expected call of RenewJWT
func (mr *MockUserManagementApiClientMockRecorder) RenewJWT(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewJWT", reflect.TypeOf((*MockUserManagementApiClient)(nil).RenewJWT), varargs...)
}

// ResendContactVerification mocks base method
func (m *MockUserManagementApiClient) ResendContactVerification(arg0 context.Context, arg1 *api.ResendContactVerificationReq, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResendContactVerification", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResendContactVerification indicates an expected call of ResendContactVerification
func (mr *MockUserManagementApiClientMockRecorder) ResendContactVerification(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendContactVerification", reflect.TypeOf((*MockUserManagementApiClient)(nil).ResendContactVerification), varargs...)
}

// ResetPassword mocks base method
func (m *MockUserManagementApiClient) ResetPassword(arg0 context.Context, arg1 *api.ResetPasswordMsg, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetPassword", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetPassword indicates an expected call of ResetPassword
func (mr *MockUserManagementApiClientMockRecorder) ResetPassword(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUserManagementApiClient)(nil).ResetPassword), varargs...)
}

// RevokeAllRefreshTokens mocks base method
func (m *MockUserManagementApiClient) RevokeAllRefreshTokens(arg0 context.Context, arg1 *api.RevokeRefreshTokensReq, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokeAllRefreshTokens", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAllRefreshTokens indicates an expected call of RevokeAllRefreshTokens
func (mr *MockUserManagementApiClientMockRecorder) RevokeAllRefreshTokens(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAllRefreshTokens", reflect.TypeOf((*MockUserManagementApiClient)(nil).RevokeAllRefreshTokens), varargs...)
}

// SaveProfile mocks base method
func (m *MockUserManagementApiClient) SaveProfile(arg0 context.Context, arg1 *api.ProfileRequest, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveProfile", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveProfile indicates an expected call of SaveProfile
func (mr *MockUserManagementApiClientMockRecorder) SaveProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveProfile", reflect.TypeOf((*MockUserManagementApiClient)(nil).SaveProfile), varargs...)
}

// SignupWithEmail mocks base method
func (m *MockUserManagementApiClient) SignupWithEmail(arg0 context.Context, arg1 *api.SignupWithEmailMsg, arg2 ...grpc.CallOption) (*api.TokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignupWithEmail", varargs...)
	ret0, _ := ret[0].(*api.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupWithEmail indicates an expected call of SignupWithEmail
func (mr *MockUserManagementApiClientMockRecorder) SignupWithEmail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupWithEmail", reflect.TypeOf((*MockUserManagementApiClient)(nil).SignupWithEmail), varargs...)
}

// Status mocks base method
func (m *MockUserManagementApiClient) Status(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockUserManagementApiClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockUserManagementApiClient)(nil).Status), varargs...)
}

// StreamUsers mocks base method
func (m *MockUserManagementApiClient) StreamUsers(arg0 context.Context, arg1 *api.StreamUsersMsg, arg2 ...grpc.CallOption) (api.UserManagementApi_StreamUsersClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamUsers", varargs...)
	ret0, _ := ret[0].(api.UserManagementApi_StreamUsersClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamUsers indicates an expected call of StreamUsers
func (mr *MockUserManagementApiClientMockRecorder) StreamUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamUsers", reflect.TypeOf((*MockUserManagementApiClient)(nil).StreamUsers), varargs...)
}

// SwitchProfile mocks base method
func (m *MockUserManagementApiClient) SwitchProfile(arg0 context.Context, arg1 *api.SwitchProfileRequest, arg2 ...grpc.CallOption) (*api.TokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SwitchProfile", varargs...)
	ret0, _ := ret[0].(*api.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SwitchProfile indicates an expected call of SwitchProfile
func (mr *MockUserManagementApiClientMockRecorder) SwitchProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwitchProfile", reflect.TypeOf((*MockUserManagementApiClient)(nil).SwitchProfile), varargs...)
}

// UpdateContactPreferences mocks base method
func (m *MockUserManagementApiClient) UpdateContactPreferences(arg0 context.Context, arg1 *api.ContactPreferencesMsg, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateContactPreferences", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateContactPreferences indicates an expected call of UpdateContactPreferences
func (mr *MockUserManagementApiClientMockRecorder) UpdateContactPreferences(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContactPreferences", reflect.TypeOf((*MockUserManagementApiClient)(nil).UpdateContactPreferences), varargs...)
}

// UseUnsubscribeToken mocks base method
func (m *MockUserManagementApiClient) UseUnsubscribeToken(arg0 context.Context, arg1 *api.TempToken, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UseUnsubscribeToken", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UseUnsubscribeToken indicates an expected call of UseUnsubscribeToken
func (mr *MockUserManagementApiClientMockRecorder) UseUnsubscribeToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseUnsubscribeToken", reflect.TypeOf((*MockUserManagementApiClient)(nil).UseUnsubscribeToken), varargs...)
}

// ValidateAppToken mocks base method
func (m *MockUserManagementApiClient) ValidateAppToken(arg0 context.Context, arg1 *api.AppTokenRequest, arg2 ...grpc.CallOption) (*api.AppTokenValidation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateAppToken", varargs...)
	ret0, _ := ret[0].(*api.AppTokenValidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateAppToken indicates an expected call of ValidateAppToken
func (mr *MockUserManagementApiClientMockRecorder) ValidateAppToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAppToken", reflect.TypeOf((*MockUserManagementApiClient)(nil).ValidateAppToken), varargs...)
}

// ValidateJWT mocks base method
func (m *MockUserManagementApiClient) ValidateJWT(arg0 context.Context, arg1 *api.JWTRequest, arg2 ...grpc.CallOption) (*api.TokenInfos, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateJWT", varargs...)
	ret0, _ := ret[0].(*api.TokenInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateJWT indicates an expected call of ValidateJWT
func (mr *MockUserManagementApiClientMockRecorder) ValidateJWT(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateJWT", reflect.TypeOf((*MockUserManagementApiClient)(nil).ValidateJWT), varargs...)
}

// VerifyContact mocks base method
func (m *MockUserManagementApiClient) VerifyContact(arg0 context.Context, arg1 *api.TempToken, arg2 ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VerifyContact", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyContact indicates an expected call of VerifyContact
func (mr *MockUserManagementApiClientMockRecorder) VerifyContact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyContact", reflect.TypeOf((*MockUserManagementApiClient)(nil).VerifyContact), varargs...)
}

// MockUserManagementApi_StreamUsersClient is a mock of UserManagementApi_StreamUsersClient interface
type MockUserManagementApi_StreamUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserManagementApi_StreamUsersClientMockRecorder
}

// MockUserManagementApi_StreamUsersClientMockRecorder is the mock recorder for MockUserManagementApi_StreamUsersClient
type MockUserManagementApi_StreamUsersClientMockRecorder struct {
	mock *MockUserManagementApi_StreamUsersClient
}

// NewMockUserManagementApi_StreamUsersClient creates a new mock instance
func NewMockUserManagementApi_StreamUsersClient(ctrl *gomock.Controller) *MockUserManagementApi_StreamUsersClient {
	mock := &MockUserManagementApi_StreamUsersClient{ctrl: ctrl}
	mock.recorder = &MockUserManagementApi_StreamUsersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserManagementApi_StreamUsersClient) EXPECT() *MockUserManagementApi_StreamUsersClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockUserManagementApi_StreamUsersClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockUserManagementApi_StreamUsersClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockUserManagementApi_StreamUsersClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockUserManagementApi_StreamUsersClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockUserManagementApi_StreamUsersClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUserManagementApi_StreamUsersClient)(nil).Context))
}

// Header mocks base method
func (m *MockUserManagementApi_StreamUsersClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockUserManagementApi_StreamUsersClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockUserManagementApi_StreamUsersClient)(nil).Header))
}

// Recv mocks base method
func (m *MockUserManagementApi_StreamUsersClient) Recv() (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockUserManagementApi_StreamUsersClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockUserManagementApi_StreamUsersClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockUserManagementApi_StreamUsersClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockUserManagementApi_StreamUsersClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserManagementApi_StreamUsersClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockUserManagementApi_StreamUsersClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockUserManagementApi_StreamUsersClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserManagementApi_StreamUsersClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockUserManagementApi_StreamUsersClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockUserManagementApi_StreamUsersClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockUserManagementApi_StreamUsersClient)(nil).Trailer))
}
