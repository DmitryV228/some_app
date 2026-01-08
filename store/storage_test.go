package store_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"some_app/handler"
	"some_app/model"
	"some_app/service"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Find(sql string) *model.User {
	return m.Called(sql).Get(0).(*model.User)
}

func (m *UserRepositoryMock) Save(user *model.User) {
	m.Called(user)
}

func TestUserApi_ProcessRequest(t *testing.T) {
	g := gin.Default()
	g.Use(func(c *gin.Context) {
		c.Set("auth_user_id", 999)
	})

	user := &model.User{
		Id:           999,
		Name:         "Alex",
		Phone:        "79222222222",
		IsAdmin:      false,
		LastViewedAt: time.Now().Add(-time.Hour * 48),
	}

	repo := &UserRepositoryMock{}
	repo.On("Find", mock.Anything).Return(user).Twice()
	repo.On("Save", user).Once()

	userService := service.NewUserService(repo)
	logger := zap.NewNop()
	userApi := handler.NewUserApi(*userService, logger)

	g.POST("/user", userApi.ProcessRequest)

	req := httptest.NewRequest(
		http.MethodPost,
		"/user?id=999",
		strings.NewReader(`{"name":"Petr","phone":"79111111111"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	g.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	var body model.User
	err := json.Unmarshal(resp.Body.Bytes(), &body)
	require.NoError(t, err)

	require.Equal(t, "Petr", body.Name)
	require.Equal(t, "79111111111", body.Phone)

	require.Equal(t, "Petr", user.Name)
	require.Equal(t, "79111111111", user.Phone)

	require.WithinDuration(t, time.Now(), user.LastViewedAt, time.Second)
}
