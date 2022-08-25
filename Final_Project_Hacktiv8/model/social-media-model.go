package model

import "time"

type SocialMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type AddSocialMediaResponse struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type EditSocialMediaResponse struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"update_at"`
}

type SocialMediaDto struct {
	Id             int                   `json:"id"`
	Name           string                `json:"name"`
	SocialMediaUrl string                `json:"social_media_url"`
	UserId         int                   `json:"user_id"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	User           UserSosialMediaDetail `json:"user"`
}
