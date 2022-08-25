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

type UserRepo struct {
	sql *sql.DB
}

func NewUserRepo(context *sql.DB) service.UserRepository {
	return &UserRepo{sql: context}
}

func (repo *UserRepo) LoginUser(ctx context.Context, loginUser model.UserLogin) (*model.UserDto, error) {
	var user model.UserDto

	data, err := repo.sql.QueryContext(ctx, "SELECT id, username, password, email, age FROM USERS WHERE email=@email",
		sql.Named("email", loginUser.Email))
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		err := data.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Age)
		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
	}
	return &user, nil
}

func (repo *UserRepo) GetUsers(ctx context.Context) ([]*model.User, error) {
	users := []*model.User{}

	data, err := repo.sql.QueryContext(ctx, "SELECT id, username, email, age FROM USERS")
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var user model.User
		err := data.Scan(&user.Id, &user.Username, &user.Email, &user.Age)
		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserRepo) GetUserById(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	data, err := repo.sql.QueryContext(ctx, "SELECT id, username, email, age FROM USERS WHERE Id = @Id",
		sql.Named("Id", id))
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		err := data.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Age)

		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
	}
	return &user, nil
}

func (repo *UserRepo) CreateUser(ctx context.Context, user *model.UserRegistration) (*model.User, error) {

	data, err := repo.sql.QueryContext(ctx, "INSERT into USERS (username, email, password, age, created_at, profile_image_url) values (@username, @email, @password, @age, @createdAt, @profileUrl); select id = convert(bigint, SCOPE_IDENTITY())",
		sql.Named("username", user.Username),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
		sql.Named("age", user.Age),
		sql.Named("createdAt", time.Now()),
		sql.Named("profileUrl", " "))

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

	result := model.User{
		Id:       id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age}

	return &result, nil
}

func (repo *UserRepo) UpdateUser(ctx context.Context, id int, user *model.UserRequest) (*model.EditUserResponse, error) {
	var result model.EditUserResponse

	_, err := repo.sql.ExecContext(ctx, "UPDATE USERS set username = @username, email = @email, updated_at = @updatedAt where id = @id",
		sql.Named("id", id),
		sql.Named("username", user.Username),
		sql.Named("email", user.Email),
		sql.Named("updatedAt", time.Now()))
	if err != nil {
		return nil, err
	}

	data, err := repo.sql.QueryContext(ctx, "SELECT id, username, email, age, updated_at FROM USERS WHERE Id = @id",
		sql.Named("id", id))
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		err := data.Scan(
			&result.Id,
			&result.Username,
			&result.Email,
			&result.Age,
			&result.UpdatedAt)

		if err != nil {
			return nil, errors.New("error in retrieving data")
		}
	}
	return &result, nil
}

func (repo *UserRepo) DeleteUser(ctx context.Context, id int) (string, error) {
	_, err := repo.sql.ExecContext(ctx, "DELETE from USERS where id=@id",
		sql.Named("id", id))
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("your account has been sucessfully deleted")
	return result, nil
}
