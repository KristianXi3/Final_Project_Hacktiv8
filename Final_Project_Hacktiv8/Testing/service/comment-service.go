package mock_service

import (
	context "context"
	model "golang-crud-sql/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommentService is a mock of CommentService interface.
type MockCommentService struct {
	ctrl     *gomock.Controller
	recorder *MockCommentServiceMockRecorder
}

// MockCommentServiceMockRecorder is the mock recorder for MockCommentService.
type MockCommentServiceMockRecorder struct {
	mock *MockCommentService
}

// NewMockCommentService creates a new mock instance.
func NewMockCommentService(ctrl *gomock.Controller) *MockCommentService {
	mock := &MockCommentService{ctrl: ctrl}
	mock.recorder = &MockCommentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentService) EXPECT() *MockCommentServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCommentService) Add(ctx context.Context, comment *model.CommentRequest, userId int) (*model.CreateCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, comment, userId)
	ret0, _ := ret[0].(*model.CreateCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockCommentServiceMockRecorder) Add(ctx, comment, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCommentService)(nil).Add), ctx, comment, userId)
}

// Delete mocks base method.
func (m *MockCommentService) Delete(ctx context.Context, id, userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentServiceMockRecorder) Delete(ctx, id, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommentService)(nil).Delete), ctx, id, userId)
}

// Get mocks base method.
func (m *MockCommentService) Get(ctx context.Context) ([]*model.CommentDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]*model.CommentDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCommentServiceMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCommentService)(nil).Get), ctx)
}

// Update mocks base method.
func (m *MockCommentService) Update(ctx context.Context, id, userId int, comment *model.EditCommentRequest) (*model.EditCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, userId, comment)
	ret0, _ := ret[0].(*model.EditCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCommentServiceMockRecorder) Update(ctx, id, userId, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommentService)(nil).Update), ctx, id, userId, comment)
}

// MockCommentRepository is a mock of CommentRepository interface.
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository.
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance.
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return m.recorder
}

// AddComment mocks base method.
func (m *MockCommentRepository) AddComment(ctx context.Context, comment *model.CommentRequest, userId int) (*model.CreateCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", ctx, comment, userId)
	ret0, _ := ret[0].(*model.CreateCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment.
func (mr *MockCommentRepositoryMockRecorder) AddComment(ctx, comment, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockCommentRepository)(nil).AddComment), ctx, comment, userId)
}

// DeleteComment mocks base method.
func (m *MockCommentRepository) DeleteComment(ctx context.Context, id, userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ctx, id, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockCommentRepositoryMockRecorder) DeleteComment(ctx, id, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockCommentRepository)(nil).DeleteComment), ctx, id, userId)
}

// GetComments mocks base method.
func (m *MockCommentRepository) GetComments(ctx context.Context) ([]*model.CommentDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", ctx)
	ret0, _ := ret[0].([]*model.CommentDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockCommentRepositoryMockRecorder) GetComments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockCommentRepository)(nil).GetComments), ctx)
}

// UpdateComment mocks base method.
func (m *MockCommentRepository) UpdateComment(ctx context.Context, id, userId int, comment *model.EditCommentRequest) (*model.EditCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", ctx, id, userId, comment)
	ret0, _ := ret[0].(*model.EditCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockCommentRepositoryMockRecorder) UpdateComment(ctx, id, userId, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockCommentRepository)(nil).UpdateComment), ctx, id, userId, comment)
}
