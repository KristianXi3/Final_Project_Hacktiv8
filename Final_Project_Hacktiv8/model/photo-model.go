package model

import (
	"time"
)

type PhotoDto struct {
	Id        int         `json:"id"`
	Title     string      `json:"title"`
	Caption   string      `json:"caption"`
	PhotoUrl  string      `json:"photo_url"`
	UserId    int         `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	User      UserRequest `json:"user"`
}

type PhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type PhotoRequestWithUserId struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

type CreatePhotoResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type EditPhotoResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoCommentDetail struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}
