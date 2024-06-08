package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/nk-hung/go-ecommerce-backend-api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1/2024")
	{
		v1.GET("/users", c.NewUserController().GetUserInfo)
	}

	return r
}

func once(x [3]int) [3]int {
	for i := range x {
		x[i] *= 2
	}
	return x
}

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}
