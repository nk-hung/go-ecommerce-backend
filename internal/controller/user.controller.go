package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nk-hung/go-ecommerce-backend-api/internal/services"
	"github.com/nk-hung/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, response.ResponseData{
		Code:    20011,
		Message: "success",
		Data: []string{
			"hungnk", "007",
		},
	})
}
