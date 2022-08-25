package mock_service

import (
	context "context"
	model "golang-crud-sql/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// DeleteUser mocks base method.
func (m *MockUserService) DeleteUser(ctx context.Context, id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserService)(nil).DeleteUser), ctx, id)
}

// GetUserById mocks base method.
func (m *MockUserService) GetUserById(ctx context.Context, id int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserService)(nil).GetUserById), ctx, id)
}

// GetUsers mocks base method.
func (m *MockUserService) GetUsers(ctx context.Context) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServiceMockRecorder) GetUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserService)(nil).GetUsers), ctx)
}

// Login mocks base method.
func (m *MockUserService) Login(ctx context.Context, login model.UserLogin) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, login)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), ctx, login)
}

// Register mocks base method.
func (m *MockUserService) Register(ctx context.Context, user *model.UserRegistration) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserService)(nil).Register), ctx, user)
}

// UpdateUser mocks base method.
func (m *MockUserService) UpdateUser(ctx context.Context, id int, user *model.UserRequest) (*model.EditUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, user)
	ret0, _ := ret[0].(*model.EditUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServiceMockRecorder) UpdateUser(ctx, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserService)(nil).UpdateUser), ctx, id, user)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, user *model.UserRegistration) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(ctx context.Context, id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), ctx, id)
}

// GetUserById mocks base method.
func (m *MockUserRepository) GetUserById(ctx context.Context, id int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserRepositoryMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserRepository)(nil).GetUserById), ctx, id)
}

// GetUsers mocks base method.
func (m *MockUserRepository) GetUsers(ctx context.Context) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserRepositoryMockRecorder) GetUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserRepository)(nil).GetUsers), ctx)
}

// LoginUser mocks base method.
func (m *MockUserRepository) LoginUser(ctx context.Context, login model.UserLogin) (*model.UserDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", ctx, login)
	ret0, _ := ret[0].(*model.UserDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockUserRepositoryMockRecorder) LoginUser(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockUserRepository)(nil).LoginUser), ctx, login)
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(ctx context.Context, id int, user *model.UserRequest) (*model.EditUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, user)
	ret0, _ := ret[0].(*model.EditUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(ctx, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), ctx, id, user)
}
