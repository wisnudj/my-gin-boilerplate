package main

import (
	"my-gin-boilerplate/db"
	"my-gin-boilerplate/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Init()

	r := gin.Default()
	routes.PeopleRoute(r)
	r.Run()

	defer db.CloseDB()
}
