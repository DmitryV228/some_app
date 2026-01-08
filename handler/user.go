package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"some_app/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
	userService *service.UserService
	logger      *zap.Logger
}

func NewUserApi(userService service.UserService, logger *zap.Logger) *UserApi {
	return &UserApi{
		userService: &userService,
		logger:      logger,
	}
}

func (api *UserApi) ProcessRequest(c *gin.Context) {
	if c.Request.Method != "POST" {
		api.logger.Info("Request method not POST")
		return
	}

	authUserId := c.GetInt("auth_user_id")
	userIdStr := c.Query("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		api.logger.Info("error conv user id")
		return
	}

	bytes, _ := io.ReadAll(c.Request.Body)
	body := make(map[string]string)
	if err := json.Unmarshal(bytes, &body); err != nil {
		api.logger.Info("fail json unmarshal")
	}

	name := body["name"]
	phone := body["phone"]

	user, err := api.userService.ChangeProfile(authUserId, userId, name, phone)

	if err != nil {
		api.logger.Info("error update user")

		if err.Error() == "fail user param" {
			api.logger.Info("errs list not empty")
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		c.Status(http.StatusForbidden)
		return
	}

	api.logger.Debug("user saved", zap.Int("user_id", user.Id))

	c.JSON(http.StatusOK, user)
}
