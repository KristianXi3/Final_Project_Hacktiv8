package service

import (
	"context"
	"errors"
	"golang-crud-sql/helper"
	"golang-crud-sql/model"
	"net/mail"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserById(ctx context.Context, id int) (*model.User, error)
	Register(ctx context.Context, user *model.UserRegistration) (*model.User, error)
	UpdateUser(ctx context.Context, id int, user *model.UserRequest) (*model.EditUserResponse, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	Login(ctx context.Context, login model.UserLogin) (*model.User, error)
}

type UserRepository interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserById(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.UserRegistration) (*model.User, error)
	UpdateUser(ctx context.Context, id int, user *model.UserRequest) (*model.EditUserResponse, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	LoginUser(ctx context.Context, login model.UserLogin) (*model.UserDto, error)
}

type UserSvc struct {
	userRepo UserRepository
}

func NewUserSvc(userRepo UserRepository) UserService {
	return &UserSvc{userRepo: userRepo}
}

func (u *UserSvc) Login(ctx context.Context, login model.UserLogin) (*model.User, error) {
	if err := validateLoginUser(&login); err != nil {
		return nil, err
	}

	userInDb, err := u.userRepo.LoginUser(ctx, login)
	if err != nil {
		return nil, err
	}

	if userInDb.Id == 0 {
		return nil, errors.New("user doesn't exist")
	}

	check := helper.CheckPasswordHash(login.Password, userInDb.Password)
	if !check {
		return nil, errors.New("invalid password")
	}

	user := model.User{
		Id:       userInDb.Id,
		Username: userInDb.Username,
		Email:    userInDb.Email,
		Age:      userInDb.Age}
	return &user, nil
}

func (u *UserSvc) GetUsers(ctx context.Context) ([]*model.User, error) {
	return u.userRepo.GetUsers(ctx)
}

func (u *UserSvc) GetUserById(ctx context.Context, id int) (*model.User, error) {
	return u.userRepo.GetUserById(ctx, id)
}

func (u *UserSvc) Register(ctx context.Context, user *model.UserRegistration) (*model.User, error) {

	var userInDb *model.User

	if err := validateUser(user); err != nil {
		return nil, err
	}

	hashPassword, err := helper.GeneratehashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashPassword

	userInDb, err = u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New("registration failed")
	}

	return userInDb, nil
}

func (u *UserSvc) UpdateUser(ctx context.Context, id int, user *model.UserRequest) (*model.EditUserResponse, error) {
	if err := validateEditUser(user); err != nil {
		return nil, err
	}

	userInDb, err := u.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, err
	}

	if userInDb.Id == 0 {
		return nil, errors.New("user doesn't exist")
	}

	return userInDb, nil
}

func (u *UserSvc) DeleteUser(ctx context.Context, id int) (string, error) {
	return u.userRepo.DeleteUser(ctx, id)
}

func validateUser(user *model.UserRegistration) error {

	if user.Username == "" {
		return errors.New("username cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if isValidEmail(user.Email) == false {
		return errors.New("invalid email address")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be minimum 6 characters")
	}
	if user.Age < 8 {
		return errors.New("age must be greater than 8")
	}
	return nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func validateEditUser(user *model.UserRequest) error {

	if user.Username == "" {
		return errors.New("username cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if isValidEmail(user.Email) == false {
		return errors.New("invalid email address")
	}
	return nil
}

func validateLoginUser(user *model.UserLogin) error {

	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if isValidEmail(user.Email) == false {
		return errors.New("invalid email address")
	}

	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be minimum 6 characters")
	}
	return nil
}
