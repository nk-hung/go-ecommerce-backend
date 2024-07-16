package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ur *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := Router.Group("/admin")
	{
		adminRouterPrivate.POST("/login")
	}
}
