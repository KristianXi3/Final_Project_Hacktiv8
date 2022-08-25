package mock_service_test

import (
	context "context"
	"errors"
	"golang-crud-sql/helper"
	"golang-crud-sql/model"
	"golang-crud-sql/service"
	mock_service "golang-crud-sql/test/mock/service"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewUserSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Initiate New User Service", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		require.NotNil(t, userService)
	})
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty Username", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := &model.UserRegistration{
			Username: "",
			Password: "password",
			Email:    "email@email.com",
			Age:      25,
		}

		res, err := userService.Register(context.Background(), user)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("username cannot be empty"), err)
	})

	t.Run("Database down", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := &model.UserRegistration{
			Username: "abc 123",
			Password: "password",
			Email:    "email@email.com",
			Age:      25,
		}
		mockUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("registration failed"))

		res, err := userService.Register(context.Background(), user)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, "registration failed", err.Error())
	})

	t.Run("Successfullly insert to db", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := &model.UserRegistration{
			Username: "abc 123",
			Password: "password",
			Email:    "email@email.com",
			Age:      25,
		}
		userRes := &model.User{
			Id:       1,
			Username: "abc 123",
			Email:    "email@email.com",
			Age:      25,
		}
		mockUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(userRes, nil)

		res, err := userService.Register(context.Background(), user)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty Password", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := model.UserLogin{
			Email:    "email@email.com",
			Password: "",
		}

		res, err := userService.Login(context.Background(), user)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("password cannot be empty"), err)
	})

	t.Run("Successfullly Login", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := model.UserLogin{
			Email:    "email@email.com",
			Password: "password",
		}
		hashPassword, _ := helper.GeneratehashPassword(user.Password)
		userRes := &model.UserDto{
			Id:       1,
			Username: "abc 123",
			Email:    "email@email.com",
			Password: hashPassword,
			Age:      25,
		}
		mockUserRepo.EXPECT().LoginUser(gomock.Any(), user).Return(userRes, nil)

		res, err := userService.Login(context.Background(), user)
		require.Nil(t, err)
		require.NotNil(t, res)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty username", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := model.UserRequest{
			Email:    "email@email.com",
			Username: "",
		}

		res, err := userService.UpdateUser(context.Background(), 1, &user)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, errors.New("username cannot be empty"), err)
	})

	t.Run("Update Successfullly", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := model.UserRequest{
			Email:    "email@email.com",
			Username: "username",
		}

		userRes := &model.EditUserResponse{
			Id:        1,
			Username:  "username",
			Email:     "email@email.com",
			Age:       25,
			UpdatedAt: time.Now(),
		}
		mockUserRepo.EXPECT().UpdateUser(gomock.Any(), 1, &user).Return(userRes, nil)

		res, err := userService.UpdateUser(context.Background(), 1, &user)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, 1, res.Id)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Delete failed", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)

		mockUserRepo.EXPECT().DeleteUser(gomock.Any(), 1).Return("", errors.New("Delete user failed"))

		res, err := userService.DeleteUser(context.Background(), 1)
		require.Error(t, err)
		require.NotNil(t, res)
		require.Equal(t, "Delete user failed", err.Error())
	})

	t.Run("Successfully delete", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)

		mockUserRepo.EXPECT().DeleteUser(gomock.Any(), 1).Return("your account has been sucessfully deleted", nil)

		res, err := userService.DeleteUser(context.Background(), 1)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, "your account has been sucessfully deleted", res)
	})
}
