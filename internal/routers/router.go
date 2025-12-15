package routers

import (
	"yyphan-pw/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/series", controllers.GetSeriesList)
	}
}
