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

type SocialMediaRepo struct {
	sql *sql.DB
}

func NewSocialMediaRepo(context *sql.DB) service.SocialMediaRepository {
	return &SocialMediaRepo{sql: context}
}

func (repo *SocialMediaRepo) AddSocialMedia(ctx context.Context, sosialMedia *model.SocialMediaRequest, userId int) (*model.AddSocialMediaResponse, error) {

	var createdDate = time.Now()
	data, err := repo.sql.QueryContext(ctx, "INSERT into Social_Medias (name, social_media_url, user_id, created_at) values (@name, @social_media_url, @userId, @createdAt); select id = convert(bigint, SCOPE_IDENTITY())",
		sql.Named("name", sosialMedia.Name),
		sql.Named("social_media_url", sosialMedia.SocialMediaUrl),
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

	result := model.AddSocialMediaResponse{
		Id:             id,
		Name:           sosialMedia.Name,
		SocialMediaUrl: sosialMedia.SocialMediaUrl,
		UserId:         userId,
		CreatedAt:      createdDate}

	return &result, nil
}

func (repo *SocialMediaRepo) GetSocialMedias(ctx context.Context) ([]*model.SocialMediaDto, error) {
	socialMedias := []*model.SocialMediaDto{}

	data, err := repo.sql.QueryContext(ctx, "SELECT Social_Medias.id, name, social_media_url, Social_Medias.user_id, Social_Medias.created_at, Social_Medias.updated_at, Users.id, Users.username, Users.profile_image_url FROM Social_Medias INNER JOIN Users on Users.id = Social_Medias.user_id")
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var socialMedia model.SocialMediaDto
		err := data.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt, &socialMedia.User.Id, &socialMedia.User.Username, &socialMedia.User.ProfileImageUrl)
		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
		socialMedias = append(socialMedias, &socialMedia)
	}
	return socialMedias, nil
}

func (repo *SocialMediaRepo) UpdateSocialMedia(ctx context.Context, id, userId int, socialMedia *model.SocialMediaRequest) (*model.EditSocialMediaResponse, error) {
	var result model.EditSocialMediaResponse

	_, err := repo.sql.ExecContext(ctx, "UPDATE Social_Medias set name = @name, social_media_url = @socialMediaUrl, updated_at = @updatedAt where id = @id AND user_id = @userId",
		sql.Named("id", id),
		sql.Named("userId", userId),
		sql.Named("name", socialMedia.Name),
		sql.Named("socialMediaUrl", socialMedia.SocialMediaUrl),
		sql.Named("updatedAt", time.Now()))
	if err != nil {
		return nil, err
	}

	data, err := repo.sql.QueryContext(ctx, "SELECT id, name, social_media_url, user_id, updated_at FROM Social_Medias WHERE id = @id AND user_id = @userId",
		sql.Named("id", id),
		sql.Named("userId", userId))
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		err := data.Scan(
			&result.Id,
			&result.Name,
			&result.SocialMediaUrl,
			&result.UserId,
			&result.UpdatedAt)

		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
	}
	return &result, nil
}

func (repo *SocialMediaRepo) DeleteSocialMedia(ctx context.Context, id, userId int) (string, error) {
	var result string

	res, err := repo.sql.ExecContext(ctx, "DELETE from Social_Medias where id=@id AND user_id = @userId",
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
		result = fmt.Sprintf("your social media doesn't exist")
	} else {
		result = fmt.Sprintf("your social media has been sucessfully deleted")
	}
	return result, nil
}
