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

func TestNewCommentSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Initiate New Comment Service", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		require.NotNil(t, commentService)
	})
}

func TestAddComment(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty message", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		comment := &model.CommentRequest{
			Message: "",
			PhotoId: 1,
		}

		res, err := commentService.Add(context.Background(), comment, 1)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("message cannot be empty"), err)
	})

	t.Run("Insert to database failed", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		comment := &model.CommentRequest{
			Message: "test comment",
			PhotoId: 1,
		}
		mockCommentRepo.EXPECT().AddComment(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("add comment failed"))

		res, err := commentService.Add(context.Background(), comment, 1)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, "add comment failed", err.Error())
	})

	t.Run("Successfullly insert to db", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		comment := &model.CommentRequest{
			Message: "test comment",
			PhotoId: 1,
		}
		commentRes := &model.CreateCommentResponse{
			Id:        1,
			Message:   "test comment",
			PhotoId:   1,
			UserId:    1,
			CreatedAt: time.Now(),
		}
		mockCommentRepo.EXPECT().AddComment(gomock.Any(), gomock.Any(), gomock.Any()).Return(commentRes, nil)

		res, err := commentService.Add(context.Background(), comment, 1)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}

func TestUpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty message", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		comment := &model.EditCommentRequest{
			Message: "",
		}

		res, err := commentService.Update(context.Background(), 1, 1, comment)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("message cannot be empty"), err)
	})

	t.Run("Update failed", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		comment := &model.EditCommentRequest{
			Message: "update comment",
		}
		mockCommentRepo.EXPECT().UpdateComment(gomock.Any(), gomock.Any(), gomock.Any(), comment).Return(nil, errors.New("update comment failed"))

		res, err := commentService.Update(context.Background(), 1, 1, comment)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, "update comment failed", err.Error())
	})

	t.Run("Successfullly update", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		comment := &model.EditCommentRequest{
			Message: "update comment",
		}

		commentRes := &model.EditCommentResponse{
			Id:        1,
			Message:   "update comment",
			PhotoId:   1,
			UserId:    1,
			UpdatedAt: time.Now(),
		}
		mockCommentRepo.EXPECT().UpdateComment(gomock.Any(), gomock.Any(), gomock.Any(), comment).Return(commentRes, nil)

		res, err := commentService.Update(context.Background(), 1, 1, comment)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}

func TestDeleteComment(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Delete failed", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)

		mockCommentRepo.EXPECT().DeleteComment(gomock.Any(), 1, 1).Return("", errors.New("Delete comment failed"))

		res, err := commentService.Delete(context.Background(), 1, 1)
		require.Error(t, err)
		require.NotNil(t, res)
		require.Equal(t, "Delete comment failed", err.Error())
	})

	t.Run("Successfully delete", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)

		mockCommentRepo.EXPECT().DeleteComment(gomock.Any(), 1, 1).Return("your comment has been sucessfully deleted", nil)

		res, err := commentService.Delete(context.Background(), 1, 1)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, "your comment has been sucessfully deleted", res)
	})
}
