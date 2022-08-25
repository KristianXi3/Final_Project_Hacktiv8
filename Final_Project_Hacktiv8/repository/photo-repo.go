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

type PhotoRepo struct {
	sql *sql.DB
}

func NewPhotoRepo(context *sql.DB) service.PhotoRepository {
	return &PhotoRepo{sql: context}
}

func (repo *PhotoRepo) AddPhoto(ctx context.Context, photo *model.PhotoRequest, userId int) (*model.CreatePhotoResponse, error) {

	var createdDate = time.Now()
	data, err := repo.sql.QueryContext(ctx, "INSERT into PHOTOS (title, caption, photo_url, user_id, created_at) values (@title, @caption, @photoUrl, @userId, @createdAt); select id = convert(bigint, SCOPE_IDENTITY())",
		sql.Named("title", photo.Title),
		sql.Named("caption", photo.Caption),
		sql.Named("photoUrl", photo.PhotoUrl),
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

	result := model.CreatePhotoResponse{
		Id:        id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    userId,
		CreatedAt: createdDate}

	return &result, nil
}

func (repo *PhotoRepo) GetPhotos(ctx context.Context) ([]*model.PhotoDto, error) {
	photos := []*model.PhotoDto{}

	data, err := repo.sql.QueryContext(ctx, "SELECT Photos.id, title, caption, photo_url, user_id, Photos.created_at, Photos.updated_at, Users.username, Users.email FROM Photos INNER JOIN Users on Photos.user_id = Users.id")
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var photo model.PhotoDto
		err := data.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt, &photo.User.Username, &photo.User.Email)
		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
		photos = append(photos, &photo)
	}
	return photos, nil
}

func (repo *PhotoRepo) UpdatePhoto(ctx context.Context, id, userId int, photo *model.PhotoRequest) (*model.EditPhotoResponse, error) {
	var result model.EditPhotoResponse

	_, err := repo.sql.ExecContext(ctx, "UPDATE Photos set title = @title, caption = @caption, photo_url = @photoUrl, updated_at = @updatedAt where id = @id AND user_id = @userId",
		sql.Named("id", id),
		sql.Named("title", photo.Title),
		sql.Named("caption", photo.Caption),
		sql.Named("photoUrl", photo.PhotoUrl),
		sql.Named("updatedAt", time.Now()),
		sql.Named("userId", userId))
	if err != nil {
		return nil, err
	}

	data, err := repo.sql.QueryContext(ctx, "SELECT id, title, caption, photo_url, user_id, updated_at FROM Photos WHERE id = @id AND user_id = @userId",
		sql.Named("id", id),
		sql.Named("userId", userId))
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		err := data.Scan(
			&result.Id,
			&result.Title,
			&result.Caption,
			&result.PhotoUrl,
			&result.UserId,
			&result.UpdatedAt)

		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
	}
	return &result, nil
}

func (repo *PhotoRepo) DeletePhoto(ctx context.Context, id, userId int) (string, error) {
	var result string
	res, err := repo.sql.ExecContext(ctx, "DELETE from Photos where id=@id AND user_id = @userId",
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
		result = fmt.Sprintf("your photo doesn't exist")
	} else {
		result = fmt.Sprintf("your photo has been sucessfully deleted")
	}
	return result, nil
}
