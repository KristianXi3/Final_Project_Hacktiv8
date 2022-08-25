package mock_service

import (
	context "context"
	model "golang-crud-sql/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSocialMediaService is a mock of SocialMediaService interface.
type MockSocialMediaService struct {
	ctrl     *gomock.Controller
	recorder *MockSocialMediaServiceMockRecorder
}

// MockSocialMediaServiceMockRecorder is the mock recorder for MockSocialMediaService.
type MockSocialMediaServiceMockRecorder struct {
	mock *MockSocialMediaService
}

// NewMockSocialMediaService creates a new mock instance.
func NewMockSocialMediaService(ctrl *gomock.Controller) *MockSocialMediaService {
	mock := &MockSocialMediaService{ctrl: ctrl}
	mock.recorder = &MockSocialMediaServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSocialMediaService) EXPECT() *MockSocialMediaServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockSocialMediaService) Add(ctx context.Context, sosialMedia *model.SocialMediaRequest, userId int) (*model.AddSocialMediaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, sosialMedia, userId)
	ret0, _ := ret[0].(*model.AddSocialMediaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockSocialMediaServiceMockRecorder) Add(ctx, sosialMedia, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSocialMediaService)(nil).Add), ctx, sosialMedia, userId)
}

// Delete mocks base method.
func (m *MockSocialMediaService) Delete(ctx context.Context, id, userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSocialMediaServiceMockRecorder) Delete(ctx, id, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSocialMediaService)(nil).Delete), ctx, id, userId)
}

// Get mocks base method.
func (m *MockSocialMediaService) Get(ctx context.Context) ([]*model.SocialMediaDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]*model.SocialMediaDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSocialMediaServiceMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSocialMediaService)(nil).Get), ctx)
}

// Update mocks base method.
func (m *MockSocialMediaService) Update(ctx context.Context, id, userId int, sosialMedia *model.SocialMediaRequest) (*model.EditSocialMediaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, userId, sosialMedia)
	ret0, _ := ret[0].(*model.EditSocialMediaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSocialMediaServiceMockRecorder) Update(ctx, id, userId, sosialMedia interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSocialMediaService)(nil).Update), ctx, id, userId, sosialMedia)
}

// MockSocialMediaRepository is a mock of SocialMediaRepository interface.
type MockSocialMediaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSocialMediaRepositoryMockRecorder
}

// MockSocialMediaRepositoryMockRecorder is the mock recorder for MockSocialMediaRepository.
type MockSocialMediaRepositoryMockRecorder struct {
	mock *MockSocialMediaRepository
}

// NewMockSocialMediaRepository creates a new mock instance.
func NewMockSocialMediaRepository(ctrl *gomock.Controller) *MockSocialMediaRepository {
	mock := &MockSocialMediaRepository{ctrl: ctrl}
	mock.recorder = &MockSocialMediaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSocialMediaRepository) EXPECT() *MockSocialMediaRepositoryMockRecorder {
	return m.recorder
}

// AddSocialMedia mocks base method.
func (m *MockSocialMediaRepository) AddSocialMedia(ctx context.Context, sosialMedia *model.SocialMediaRequest, userId int) (*model.AddSocialMediaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSocialMedia", ctx, sosialMedia, userId)
	ret0, _ := ret[0].(*model.AddSocialMediaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSocialMedia indicates an expected call of AddSocialMedia.
func (mr *MockSocialMediaRepositoryMockRecorder) AddSocialMedia(ctx, sosialMedia, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSocialMedia", reflect.TypeOf((*MockSocialMediaRepository)(nil).AddSocialMedia), ctx, sosialMedia, userId)
}

// DeleteSocialMedia mocks base method.
func (m *MockSocialMediaRepository) DeleteSocialMedia(ctx context.Context, id, userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSocialMedia", ctx, id, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSocialMedia indicates an expected call of DeleteSocialMedia.
func (mr *MockSocialMediaRepositoryMockRecorder) DeleteSocialMedia(ctx, id, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSocialMedia", reflect.TypeOf((*MockSocialMediaRepository)(nil).DeleteSocialMedia), ctx, id, userId)
}

// GetSocialMedias mocks base method.
func (m *MockSocialMediaRepository) GetSocialMedias(ctx context.Context) ([]*model.SocialMediaDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSocialMedias", ctx)
	ret0, _ := ret[0].([]*model.SocialMediaDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSocialMedias indicates an expected call of GetSocialMedias.
func (mr *MockSocialMediaRepositoryMockRecorder) GetSocialMedias(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSocialMedias", reflect.TypeOf((*MockSocialMediaRepository)(nil).GetSocialMedias), ctx)
}

// UpdateSocialMeadia mocks base method.
func (m *MockSocialMediaRepository) UpdateSocialMedia(ctx context.Context, id, userId int, socialMedia *model.SocialMediaRequest) (*model.EditSocialMediaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSocialMedia", ctx, id, userId, socialMedia)
	ret0, _ := ret[0].(*model.EditSocialMediaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSocialMeadia indicates an expected call of UpdateSocialMeadia.
func (mr *MockSocialMediaRepositoryMockRecorder) UpdateSocialMedia(ctx, id, userId, socialMedia interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSocialMedia", reflect.TypeOf((*MockSocialMediaRepository)(nil).UpdateSocialMedia), ctx, id, userId, socialMedia)
}
