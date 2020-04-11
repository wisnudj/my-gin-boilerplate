package routes

import (
	"my-gin-boilerplate/config"

	"github.com/gin-gonic/gin"
)

func Init() {
	config := config.GetConfig()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	PeopleRoute(r)

	r.Run(config.GetString("server.port"))
}
