package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-crud-sql/model"
	"golang-crud-sql/service"
	"time"
)

type CommentRepo struct {
	sql *sql.DB
}

func NewCommentRepo(context *sql.DB) service.CommentRepository {
	return &CommentRepo{sql: context}
}

func (repo *CommentRepo) AddComment(ctx context.Context, comment *model.CommentRequest, userId int) (*model.CreateCommentResponse, error) {

	var createdDate = time.Now()
	data, err := repo.sql.QueryContext(ctx, "INSERT into Comments (message, photo_id, user_id, created_at) values (@message, @photoId, @userId, @createdAt); select id = convert(bigint, SCOPE_IDENTITY())",
		sql.Named("message", comment.Message),
		sql.Named("photoId", comment.PhotoId),
		sql.Named("userId", userId),
		sql.Named("createdAt", createdDate))

	if err != nil {
		return nil, err
	}
	defer data.Close()

	var id int
	for data.Next() {
		if err = data.Scan(&id); err != nil {
			return nil, errors.New("error in retrieving data")
		}
	}

	result := model.CreateCommentResponse{
		Id:        id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    userId,
		CreatedAt: createdDate}

	return &result, nil
}

func (repo *CommentRepo) GetComments(ctx context.Context) ([]*model.CommentDto, error) {
	comments := []*model.CommentDto{}

	data, err := repo.sql.QueryContext(ctx, "SELECT Comments.id, message, photo_id, Comments.user_id, Comments.created_at, Comments.updated_at, Users.id, Users.username, Users.email, Photos.id, Photos.title, Photos.caption, Photos.photo_url, Photos.user_id FROM Comments INNER JOIN Users on Users.id = Comments.user_id INNER JOIN Photos on Photos.id = Comments.photo_id")
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var comment model.CommentDto
		err := data.Scan(&comment.Id, &comment.Message, &comment.PhotoId, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt, &comment.User.Id, &comment.User.Username, &comment.User.Email,
			&comment.Photo.Id, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.PhotoUrl, &comment.Photo.UserId)
		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (repo *CommentRepo) UpdateComment(ctx context.Context, id, userId int, comment *model.EditCommentRequest) (*model.EditCommentResponse, error) {
	var result model.EditCommentResponse

	_, err := repo.sql.ExecContext(ctx, "UPDATE Comments set message = @message, updated_at = @updatedAt where id = @id AND user_id = @userId",
		sql.Named("id", id),
		sql.Named("message", comment.Message),
		sql.Named("updatedAt", time.Now()),
		sql.Named("userId", userId))
	if err != nil {
		return nil, err
	}

	data, err := repo.sql.QueryContext(ctx, "SELECT id, message, photo_id, user_id, updated_at FROM Comments WHERE id = @id AND user_id = @userId",
		sql.Named("id", id),
		sql.Named("userId", userId))
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		err := data.Scan(
			&result.Id,
			&result.Message,
			&result.PhotoId,
			&result.UserId,
			&result.UpdatedAt)

		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
	}
	return &result, nil
}

func (repo *CommentRepo) DeleteComment(ctx context.Context, id, userId int) (string, error) {
	var result string
	res, err := repo.sql.ExecContext(ctx, "DELETE from Comments where id=@id AND user_id = @userId",
		sql.Named("id", id),
		sql.Named("userId", userId))
	if err != nil {
		return "", err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return "", err
	}

	if count == 0 {
		result = fmt.Sprintf("your comment doesn't exist")
	} else {
		result = fmt.Sprintf("your comment has been sucessfully deleted")
	}
	return result, nil
}
