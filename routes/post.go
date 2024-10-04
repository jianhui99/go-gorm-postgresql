package routes

import (
	controller "jd99/controller"
	"jd99/middleware"

	"github.com/gin-gonic/gin"
)

type PostRouteController struct {
	postController controller.PostController
}

func NewRoutePostController(postController controller.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.postController.CreatePost)
	router.GET("/", pc.postController.FindPosts)
	router.PUT("/:postId", pc.postController.UpdatePost)
	router.GET("/:postId", pc.postController.FindPostById)
	router.DELETE("/:postId", pc.postController.DeletePost)
}
