package service

import (
	"context"
	"errors"
	"golang-crud-sql/model"
)

type PhotoService interface {
	Add(ctx context.Context, photo *model.PhotoRequest, userId int) (*model.CreatePhotoResponse, error)
	Get(ctx context.Context) ([]*model.PhotoDto, error)
	Update(ctx context.Context, id, userId int, photo *model.PhotoRequest) (*model.EditPhotoResponse, error)
	Delete(ctx context.Context, id, userId int) (string, error)
}

type PhotoRepository interface {
	AddPhoto(ctx context.Context, photo *model.PhotoRequest, userId int) (*model.CreatePhotoResponse, error)
	GetPhotos(ctx context.Context) ([]*model.PhotoDto, error)
	UpdatePhoto(ctx context.Context, id, userId int, photo *model.PhotoRequest) (*model.EditPhotoResponse, error)
	DeletePhoto(ctx context.Context, id, userId int) (string, error)
}

type PhotoSvc struct {
	photoRepo PhotoRepository
}

func NewPhotoSvc(photoRepo PhotoRepository) PhotoService {
	return &PhotoSvc{photoRepo: photoRepo}
}

func (p *PhotoSvc) Add(ctx context.Context, photo *model.PhotoRequest, userId int) (*model.CreatePhotoResponse, error) {

	if err := validatePhoto(photo); err != nil {
		return nil, err
	}

	photoInDb, err := p.photoRepo.AddPhoto(ctx, photo, userId)
	if err != nil {
		return nil, errors.New("add photo failed")
	}

	return photoInDb, nil
}

func (p *PhotoSvc) Get(ctx context.Context) ([]*model.PhotoDto, error) {
	return p.photoRepo.GetPhotos(ctx)
}

func (p *PhotoSvc) Update(ctx context.Context, id, userId int, photo *model.PhotoRequest) (*model.EditPhotoResponse, error) {
	if err := validatePhoto(photo); err != nil {
		return nil, err
	}

	photoInDb, err := p.photoRepo.UpdatePhoto(ctx, id, userId, photo)
	if err != nil {
		return nil, err
	}

	if photoInDb.Id == 0 {
		return nil, errors.New("your photo doesn't exist")
	}

	return photoInDb, nil

}

func (p *PhotoSvc) Delete(ctx context.Context, id, userId int) (string, error) {
	return p.photoRepo.DeletePhoto(ctx, id, userId)
}

func validatePhoto(photo *model.PhotoRequest) error {
	if photo.Title == "" {
		return errors.New("title cannot be empty")
	}
	if photo.PhotoUrl == "" {
		return errors.New("photo url cannot be empty")
	}
	return nil
}
