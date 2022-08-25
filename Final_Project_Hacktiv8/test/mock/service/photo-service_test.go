package mock_service_test

import (
	context "context"
	"errors"
	"golang-crud-sql/model"
	"golang-crud-sql/service"
	mock_service "golang-crud-sql/test/mock/service"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPhotoSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Initiate New Photo Service", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		require.NotNil(t, photoService)
	})
}

func TestAddPhoto(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty title", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		photo := &model.PhotoRequest{
			Title:    "",
			Caption:  "random picture",
			PhotoUrl: "https://i.picsum.photos/id/744/200/200.jpg?hmac=8T0b9ya-1hs9mQn1Sosud4eldJZ6haGcupAiLTJHe2o",
		}

		res, err := photoService.Add(context.Background(), photo, 1)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("title cannot be empty"), err)
	})

	t.Run("Insert to database failed", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		photo := &model.PhotoRequest{
			Title:    "Capture",
			Caption:  "random picture",
			PhotoUrl: "https://i.picsum.photos/id/744/200/200.jpg?hmac=8T0b9ya-1hs9mQn1Sosud4eldJZ6haGcupAiLTJHe2o",
		}
		mockPhotoRepo.EXPECT().AddPhoto(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("add photo failed"))

		res, err := photoService.Add(context.Background(), photo, 1)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, "add photo failed", err.Error())
	})

	t.Run("Successfullly insert to db", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		photo := &model.PhotoRequest{
			Title:    "Capture",
			Caption:  "random picture",
			PhotoUrl: "https://i.picsum.photos/id/744/200/200.jpg?hmac=8T0b9ya-1hs9mQn1Sosud4eldJZ6haGcupAiLTJHe2o",
		}

		photoRes := &model.CreatePhotoResponse{
			Id:        1,
			Title:     "Capture",
			Caption:   "random picture",
			PhotoUrl:  "https://i.picsum.photos/id/744/200/200.jpg?hmac=8T0b9ya-1hs9mQn1Sosud4eldJZ6haGcupAiLTJHe2o",
			UserId:    1,
			CreatedAt: time.Now(),
		}
		mockPhotoRepo.EXPECT().AddPhoto(gomock.Any(), gomock.Any(), gomock.Any()).Return(photoRes, nil)

		res, err := photoService.Add(context.Background(), photo, 1)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}

func TestUpdatePhotos(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty photo url", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		photo := &model.PhotoRequest{
			Title:    "test",
			Caption:  "random picture",
			PhotoUrl: "",
		}

		res, err := photoService.Update(context.Background(), 1, 1, photo)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("photo url cannot be empty"), err)
	})

	t.Run("Insert to database failed", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		photo := &model.PhotoRequest{
			Title:    "update photo",
			Caption:  "random picture",
			PhotoUrl: "https://i.picsum.photos/id/744/200/200.jpg?hmac=8T0b9ya-1hs9mQn1Sosud4eldJZ6haGcupAiLTJHe2o",
		}
		mockPhotoRepo.EXPECT().UpdatePhoto(gomock.Any(), gomock.Any(), gomock.Any(), photo).Return(nil, errors.New("update photo failed"))

		res, err := photoService.Update(context.Background(), 1, 1, photo)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, "update photo failed", err.Error())
	})

	t.Run("Successfullly insert to db", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		photo := &model.PhotoRequest{
			Title:    "Capture",
			Caption:  "random picture",
			PhotoUrl: "https://i.picsum.photos/id/744/200/200.jpg?hmac=8T0b9ya-1hs9mQn1Sosud4eldJZ6haGcupAiLTJHe2o",
		}

		photoRes := &model.EditPhotoResponse{
			Id:        1,
			Title:     "Capture",
			Caption:   "random picture",
			PhotoUrl:  "https://i.picsum.photos/id/744/200/200.jpg?hmac=8T0b9ya-1hs9mQn1Sosud4eldJZ6haGcupAiLTJHe2o",
			UserId:    1,
			UpdatedAt: time.Now(),
		}
		mockPhotoRepo.EXPECT().UpdatePhoto(gomock.Any(), gomock.Any(), gomock.Any(), photo).Return(photoRes, nil)

		res, err := photoService.Update(context.Background(), 1, 1, photo)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}

func TestDeletePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Delete failed", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)

		mockPhotoRepo.EXPECT().DeletePhoto(gomock.Any(), 1, 1).Return("", errors.New("Delete photo failed"))

		res, err := photoService.Delete(context.Background(), 1, 1)
		require.Error(t, err)
		require.NotNil(t, res)
		require.Equal(t, "Delete photo failed", err.Error())
	})

	t.Run("Successfully delete", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)

		mockPhotoRepo.EXPECT().DeletePhoto(gomock.Any(), 1, 1).Return("your photo has been sucessfully deleted", nil)

		res, err := photoService.Delete(context.Background(), 1, 1)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, "your photo has been sucessfully deleted", res)
	})
}
