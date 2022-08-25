package service

import (
	"context"
	"errors"
	"golang-crud-sql/model"
)

type SocialMediaService interface {
	Add(ctx context.Context, sosialMedia *model.SocialMediaRequest, userId int) (*model.AddSocialMediaResponse, error)
	Get(ctx context.Context) ([]*model.SocialMediaDto, error)
	Update(ctx context.Context, id, userId int, sosialMedia *model.SocialMediaRequest) (*model.EditSocialMediaResponse, error)
	Delete(ctx context.Context, id, userId int) (string, error)
}

type SocialMediaRepository interface {
	AddSocialMedia(ctx context.Context, sosialMedia *model.SocialMediaRequest, userId int) (*model.AddSocialMediaResponse, error)
	GetSocialMedias(ctx context.Context) ([]*model.SocialMediaDto, error)
	UpdateSocialMedia(ctx context.Context, id, userId int, socialMedia *model.SocialMediaRequest) (*model.EditSocialMediaResponse, error)
	DeleteSocialMedia(ctx context.Context, id, userId int) (string, error)
}

type SocialMediaSvc struct {
	socialMediaRepo SocialMediaRepository
}

func NewSocialMediaSvc(socialMediaRepo SocialMediaRepository) SocialMediaService {
	return &SocialMediaSvc{socialMediaRepo: socialMediaRepo}
}

func (s *SocialMediaSvc) Add(ctx context.Context, sosialMedia *model.SocialMediaRequest, userId int) (*model.AddSocialMediaResponse, error) {

	if err := validateSocialMedia(sosialMedia); err != nil {
		return nil, err
	}

	socialMediaInDb, err := s.socialMediaRepo.AddSocialMedia(ctx, sosialMedia, userId)
	if err != nil {
		return nil, errors.New("add social media failed")
	}

	return socialMediaInDb, nil
}

func (s *SocialMediaSvc) Get(ctx context.Context) ([]*model.SocialMediaDto, error) {
	return s.socialMediaRepo.GetSocialMedias(ctx)
}

func (s *SocialMediaSvc) Update(ctx context.Context, id, userId int, sosialMedia *model.SocialMediaRequest) (*model.EditSocialMediaResponse, error) {
	if err := validateSocialMedia(sosialMedia); err != nil {
		return nil, err
	}

	socialMediaInDb, err := s.socialMediaRepo.UpdateSocialMedia(ctx, id, userId, sosialMedia)
	if err != nil {
		return nil, err
	}

	if socialMediaInDb.Id == 0 {
		return nil, errors.New("your social media doesn't exist")
	}

	return socialMediaInDb, nil
}

func (s *SocialMediaSvc) Delete(ctx context.Context, id, userId int) (string, error) {
	return s.socialMediaRepo.DeleteSocialMedia(ctx, id, userId)
}

func validateSocialMedia(sosialMedia *model.SocialMediaRequest) error {
	if sosialMedia.Name == "" {
		return errors.New("name cannot be empty")
	}
	if sosialMedia.SocialMediaUrl == "" {
		return errors.New("social media url cannot be empty")
	}
	return nil
}
