package routes

import (
	"my-gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func PeopleRoute(route *gin.Engine) {
	c := controllers.PeopleControllers{}
	r := route.Group("people")
	{
		r.GET("/", c.FindAll)
		r.GET("/:id", c.FindOne)
		r.POST("/", c.Create)
		r.PUT("/:id", c.Update)
		r.DELETE("/:id", c.Delete)
	}
}
