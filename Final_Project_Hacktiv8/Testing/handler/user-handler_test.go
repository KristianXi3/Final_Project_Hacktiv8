package mock_handler_test

import (
	"golang-crud-sql/handler"
	mock_service "golang-crud-sql/test/mock/service"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Initiate New User Handler", func(t *testing.T) {
		mockUserService := mock_service.NewMockUserService(ctrl)
		userHandler := handler.NewUserHandler(mockUserService)
		require.NotNil(t, userHandler)
	})
}
