package routers

import (
	"yyphan-pw/backend/internal/controllers"
	"yyphan-pw/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	public := r.Group("/api")
	{
		public.GET("/series", controllers.GetSeriesList)
	}

	admin := r.Group("/api/admin")
	admin.Use(middleware.AdminAuth())
	{
		admin.POST("/posts", controllers.CreatePost)
	}
}
