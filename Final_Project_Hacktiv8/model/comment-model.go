package model

import "time"

type CommentRequest struct {
	Message string `json:"message"`
	PhotoId int    `json:"photo_id"`
}

type CreateCommentResponse struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type EditCommentRequest struct {
	Message string `json:"message"`
}

type EditCommentResponse struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentDto struct {
	Id        int                `json:"id"`
	Message   string             `json:"message"`
	PhotoId   int                `json:"photo_id"`
	UserId    int                `json:"user_id"`
	UpdatedAt time.Time          `json:"updated_at"`
	CreatedAt time.Time          `json:"created_at"`
	User      UserCommentDetail  `json:"user"`
	Photo     PhotoCommentDetail `json:"photo"`
}
