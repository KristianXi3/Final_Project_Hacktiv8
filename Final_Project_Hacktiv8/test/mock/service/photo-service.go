package mock_service

import (
	context "context"
	model "golang-crud-sql/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPhotoService is a mock of PhotoService interface.
type MockPhotoService struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoServiceMockRecorder
}

// MockPhotoServiceMockRecorder is the mock recorder for MockPhotoService.
type MockPhotoServiceMockRecorder struct {
	mock *MockPhotoService
}

// NewMockPhotoService creates a new mock instance.
func NewMockPhotoService(ctrl *gomock.Controller) *MockPhotoService {
	mock := &MockPhotoService{ctrl: ctrl}
	mock.recorder = &MockPhotoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoService) EXPECT() *MockPhotoServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPhotoService) Add(ctx context.Context, photo *model.PhotoRequest, userId int) (*model.CreatePhotoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, photo, userId)
	ret0, _ := ret[0].(*model.CreatePhotoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockPhotoServiceMockRecorder) Add(ctx, photo, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPhotoService)(nil).Add), ctx, photo, userId)
}

// Delete mocks base method.
func (m *MockPhotoService) Delete(ctx context.Context, id, userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockPhotoServiceMockRecorder) Delete(ctx, id, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPhotoService)(nil).Delete), ctx, id, userId)
}

// Get mocks base method.
func (m *MockPhotoService) Get(ctx context.Context) ([]*model.PhotoDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]*model.PhotoDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPhotoServiceMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPhotoService)(nil).Get), ctx)
}

// Update mocks base method.
func (m *MockPhotoService) Update(ctx context.Context, id, userId int, photo *model.PhotoRequest) (*model.EditPhotoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, userId, photo)
	ret0, _ := ret[0].(*model.EditPhotoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPhotoServiceMockRecorder) Update(ctx, id, userId, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPhotoService)(nil).Update), ctx, id, userId, photo)
}

// MockPhotoRepository is a mock of PhotoRepository interface.
type MockPhotoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoRepositoryMockRecorder
}

// MockPhotoRepositoryMockRecorder is the mock recorder for MockPhotoRepository.
type MockPhotoRepositoryMockRecorder struct {
	mock *MockPhotoRepository
}

// NewMockPhotoRepository creates a new mock instance.
func NewMockPhotoRepository(ctrl *gomock.Controller) *MockPhotoRepository {
	mock := &MockPhotoRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoRepository) EXPECT() *MockPhotoRepositoryMockRecorder {
	return m.recorder
}

// AddPhoto mocks base method.
func (m *MockPhotoRepository) AddPhoto(ctx context.Context, photo *model.PhotoRequest, userId int) (*model.CreatePhotoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPhoto", ctx, photo, userId)
	ret0, _ := ret[0].(*model.CreatePhotoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPhoto indicates an expected call of AddPhoto.
func (mr *MockPhotoRepositoryMockRecorder) AddPhoto(ctx, photo, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPhoto", reflect.TypeOf((*MockPhotoRepository)(nil).AddPhoto), ctx, photo, userId)
}

// DeletePhoto mocks base method.
func (m *MockPhotoRepository) DeletePhoto(ctx context.Context, id, userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhoto", ctx, id, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePhoto indicates an expected call of DeletePhoto.
func (mr *MockPhotoRepositoryMockRecorder) DeletePhoto(ctx, id, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).DeletePhoto), ctx, id, userId)
}

// GetPhotos mocks base method.
func (m *MockPhotoRepository) GetPhotos(ctx context.Context) ([]*model.PhotoDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotos", ctx)
	ret0, _ := ret[0].([]*model.PhotoDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotos indicates an expected call of GetPhotos.
func (mr *MockPhotoRepositoryMockRecorder) GetPhotos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotos", reflect.TypeOf((*MockPhotoRepository)(nil).GetPhotos), ctx)
}

// UpdatePhoto mocks base method.
func (m *MockPhotoRepository) UpdatePhoto(ctx context.Context, id, userId int, photo *model.PhotoRequest) (*model.EditPhotoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoto", ctx, id, userId, photo)
	ret0, _ := ret[0].(*model.EditPhotoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePhoto indicates an expected call of UpdatePhoto.
func (mr *MockPhotoRepositoryMockRecorder) UpdatePhoto(ctx, id, userId, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).UpdatePhoto), ctx, id, userId, photo)
}
