package controllers

import (
	"my-gin-boilerplate/db"
	"my-gin-boilerplate/models"

	"github.com/gin-gonic/gin"
)

type PeopleControllers struct{}

func (u PeopleControllers) FindAll(c *gin.Context) {
	var people []models.People
	var getDB = db.GetDB()

	getDB.Find(&people)

	c.JSON(200, gin.H{
		"result": people,
	})
}

func (u PeopleControllers) Create(c *gin.Context) {
	name := c.PostForm("name")

	var people = models.People{Name: name}
	var getDB = db.GetDB()

	getDB.Create(&people)
	c.JSON(200, people)
}

func (u PeopleControllers) FindOne(c *gin.Context) {
	var getDB = db.GetDB()
	var people models.People

	id := c.Param("id")
	getDB.Find(&people, "id=?", id)

	c.JSON(200, people)
}

func (u PeopleControllers) Update(c *gin.Context) {
	var getDB = db.GetDB()
	var people models.People

	id := c.Param("id")

	getDB.Where("id = ?", id).First(&people)

	people.Name = c.PostForm("name")
	getDB.Save(&people)

	c.JSON(200, people)
}

func (u PeopleControllers) Delete(c *gin.Context) {
	id := c.Param("id")
	var people models.People
	var getDB = db.GetDB()

	getDB.Where("id = ?", id).First(&people).Delete(&people)

	c.JSON(200, people)
}
