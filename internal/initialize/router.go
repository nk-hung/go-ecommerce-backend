package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/nk-hung/go-ecommerce-backend-api/internal/controller"
	"github.com/nk-hung/go-ecommerce-backend-api/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")

		c.Next()

		fmt.Println("After --> AA")
	}
}

func BB(c *gin.Context) {
	fmt.Println("Before --> BB")

	c.Next()

	fmt.Println("After --> BB")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	// use middlewares
	r.Use(middlewares.AuthMiddleware(), AA(), BB)

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
