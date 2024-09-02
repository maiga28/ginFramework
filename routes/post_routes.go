package routes

import (
	"ginFramework/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {

	// Post routes
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.GET("/users/:user_id/posts", controllers.GetUserPosts)
}
