package mock_service_test

import (
	context "context"
	"errors"
	model "golang-crud-sql/model"
	"golang-crud-sql/service"
	mock_service "golang-crud-sql/test/mock/service"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewSocialMediaSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Initiate New Social Media Service", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)
		require.NotNil(t, socialMediaService)
	})
}

func TestAddSocialMedia(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty name", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)
		socialMedia := &model.SocialMediaRequest{
			Name:           "",
			SocialMediaUrl: "https://www.linkedin.com/",
		}

		res, err := socialMediaService.Add(context.Background(), socialMedia, 1)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("name cannot be empty"), err)
	})

	t.Run("Insert to database failed", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)
		socialMedia := &model.SocialMediaRequest{
			Name:           "Media",
			SocialMediaUrl: "https://www.linkedin.com/",
		}
		mockSocialMediaRepo.EXPECT().AddSocialMedia(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("add social media failed"))

		res, err := socialMediaService.Add(context.Background(), socialMedia, 1)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, "add social media failed", err.Error())
	})

	t.Run("Successfullly insert to db", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)
		socialMedia := &model.SocialMediaRequest{
			Name:           "Media",
			SocialMediaUrl: "https://www.linkedin.com/",
		}
		commentRes := &model.AddSocialMediaResponse{
			Id:             1,
			Name:           "test comment",
			SocialMediaUrl: "https://www.linkedin.com/",
			UserId:         1,
			CreatedAt:      time.Now(),
		}
		mockSocialMediaRepo.EXPECT().AddSocialMedia(gomock.Any(), gomock.Any(), gomock.Any()).Return(commentRes, nil)

		res, err := socialMediaService.Add(context.Background(), socialMedia, 1)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}

func TestUpdateSocialMedia(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty url", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)
		socialMedia := &model.SocialMediaRequest{
			Name:           "Edit",
			SocialMediaUrl: "",
		}

		res, err := socialMediaService.Update(context.Background(), 1, 1, socialMedia)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("social media url cannot be empty"), err)
	})

	t.Run("Update failed", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)
		socialMedia := &model.SocialMediaRequest{
			Name:           "Media",
			SocialMediaUrl: "https://www.linkedin.com/",
		}
		mockSocialMediaRepo.EXPECT().UpdateSocialMedia(gomock.Any(), gomock.Any(), gomock.Any(), socialMedia).Return(nil, errors.New("update social media failed"))

		res, err := socialMediaService.Update(context.Background(), 1, 1, socialMedia)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, "update social media failed", err.Error())
	})

	t.Run("Successfullly update", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)
		socialMedia := &model.SocialMediaRequest{
			Name:           "Media",
			SocialMediaUrl: "https://www.linkedin.com/",
		}

		socialMediaRes := &model.EditSocialMediaResponse{
			Id:             1,
			Name:           "Media",
			SocialMediaUrl: "https://www.linkedin.com/",
			UserId:         1,
			UpdatedAt:      time.Now(),
		}
		mockSocialMediaRepo.EXPECT().UpdateSocialMedia(gomock.Any(), gomock.Any(), gomock.Any(), socialMedia).Return(socialMediaRes, nil)

		res, err := socialMediaService.Update(context.Background(), 1, 1, socialMedia)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}

func TestDeleteSocialMedia(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Delete failed", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)

		mockSocialMediaRepo.EXPECT().DeleteSocialMedia(gomock.Any(), 1, 1).Return("", errors.New("Delete social media failed"))

		res, err := socialMediaService.Delete(context.Background(), 1, 1)
		require.Error(t, err)
		require.NotNil(t, res)
		require.Equal(t, "Delete social media failed", err.Error())
	})

	t.Run("Successfully delete", func(t *testing.T) {
		mockSocialMediaRepo := mock_service.NewMockSocialMediaRepository(ctrl)
		socialMediaService := service.NewSocialMediaSvc(mockSocialMediaRepo)

		mockSocialMediaRepo.EXPECT().DeleteSocialMedia(gomock.Any(), 1, 1).Return("your social media has been sucessfully deleted", nil)

		res, err := socialMediaService.Delete(context.Background(), 1, 1)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, "your social media has been sucessfully deleted", res)
	})
}
