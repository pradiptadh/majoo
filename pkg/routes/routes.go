package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pradiptadh/majoo/pkg/auth"
	"github.com/pradiptadh/majoo/pkg/middleware"
	"github.com/pradiptadh/majoo/pkg/user"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	route := r.Group("api/v1/")

	{
		authRoute := route.Group("auth")
		authRoute.POST("/register", auth.AuthRegister)
		authRoute.POST("/login", auth.AuthLogin)
		//	authRoute.GET("/validate", middleware.RequireAuth, auth.Validate)
	}
	{
		userRoute := route.Group("user")
		userRoute.GET("", middleware.RequireAuth, user.GetUser)
		userRoute.GET("/merchant", middleware.RequireAuth, user.GetMerchant)
		userRoute.GET("/detail-merchant", middleware.RequireAuth, user.GetDetailMerchant)
	}
	return r
}
