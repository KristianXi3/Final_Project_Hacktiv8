package service

import (
	"context"
	"errors"
	model "golang-crud-sql/models"
)

type CommentService interface {
	Add(ctx context.Context, comment *model.CommentRequest, userId int) (*model.CreateCommentResponse, error)
	Get(ctx context.Context) ([]*model.CommentDto, error)
	Update(ctx context.Context, id, userId int, comment *model.EditCommentRequest) (*model.EditCommentResponse, error)
	Delete(ctx context.Context, id, userId int) (string, error)
}

type CommentRepository interface {
	AddComment(ctx context.Context, comment *model.CommentRequest, userId int) (*model.CreateCommentResponse, error)
	GetComments(ctx context.Context) ([]*model.CommentDto, error)
	UpdateComment(ctx context.Context, id, userId int, comment *model.EditCommentRequest) (*model.EditCommentResponse, error)
	DeleteComment(ctx context.Context, id, userId int) (string, error)
}

type CommentSvc struct {
	commentRepo CommentRepository
}

func NewCommentSvc(commentRepo CommentRepository) CommentService {
	return &CommentSvc{commentRepo: commentRepo}
}

func (c *CommentSvc) Add(ctx context.Context, comment *model.CommentRequest, userId int) (*model.CreateCommentResponse, error) {

	if comment.Message == "" {
		return nil, errors.New("message cannot be empty")
	}

	commentInDb, err := c.commentRepo.AddComment(ctx, comment, userId)
	if err != nil {
		return nil, errors.New("add comment failed")
	}

	return commentInDb, nil
}

func (c *CommentSvc) Get(ctx context.Context) ([]*model.CommentDto, error) {
	return c.commentRepo.GetComments(ctx)
}

func (c *CommentSvc) Update(ctx context.Context, id, userId int, comment *model.EditCommentRequest) (*model.EditCommentResponse, error) {
	if comment.Message == "" {
		return nil, errors.New("message cannot be empty")
	}

	commentInDb, err := c.commentRepo.UpdateComment(ctx, id, userId, comment)
	if err != nil {
		return nil, err
	}

	if commentInDb.Id == 0 {
		return nil, errors.New("your comment doesn't exist")
	}

	return commentInDb, nil
}

func (c *CommentSvc) Delete(ctx context.Context, id, userId int) (string, error) {
	return c.commentRepo.DeleteComment(ctx, id, userId)
}
