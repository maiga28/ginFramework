package routes

import (
	"ginFramework/controllers"
	"github.com/gin-gonic/gin"
)

func LikeRoutes(r *gin.Engine) {

	// Like routes
	r.GET("/likes", controllers.GetLikes)
	r.POST("/likes", controllers.LikePost)
	r.DELETE("/likes/:id", controllers.UnlikePost)
	r.GET("/posts/:post_id/likes", controllers.GetPostLikes)

}
