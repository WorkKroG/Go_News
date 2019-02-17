package app

import (
	"github.com/gin-gonic/gin"

	"./controllers"
)

func MakeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Add("access-control-allow-origin", "*")
		c.Writer.Header().Add("access-control-allow-headers", "accept, content-type")
		c.Writer.Header().Add("access-control-allow-methods", "GET,HEAD,POST,DELETE,OPTIONS,PUT,PATCH")
	}

	r.Use(cors)

	todosRouter := r.Group("/news")
	{
		NewsController := new(controllers.NewsController)

		todosRouter.GET("/", NewsController.Index)
		todosRouter.GET("/:id", NewsController.Show)
		todosRouter.POST("/", NewsController.Create)
		todosRouter.PUT("/:id", NewsController.Update)
		todosRouter.DELETE("/:id", NewsController.Destroy)
	}
}

func Connect(r *gin.Engine) {
	controllers.Connect(r)
}
