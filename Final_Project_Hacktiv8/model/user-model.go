package model

import "time"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}
type UserDto struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserRegistration struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type EditUserResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCommentDetail struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserSosialMediaDetail struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}
