package routes

import (
	"ginFramework/controllers"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.Engine) {
	// Comment routes
	r.GET("/comments", controllers.GetComments)
	r.GET("/comments/:id", controllers.GetComment)
	r.POST("/comments", controllers.CreateComment)
	r.PUT("/comments/:id", controllers.UpdateComment)
	r.DELETE("/comments/:id", controllers.DeleteComment)
	r.GET("/posts/:post_id/comments", controllers.GetPostComments)
}
