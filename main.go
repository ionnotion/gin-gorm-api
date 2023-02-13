package main

import (
	"gin-gorm-api/controllers"
	"gin-gorm-api/models"

	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()

	models.DbConnect()

	router.GET("/",controllers.Hello)
	router.GET("/todo",controllers.GetTodos)
	router.GET("/todo/:id",controllers.GetTodoById)
	router.POST("/todo",controllers.PostTodo)
	router.PUT("/todo/:id",controllers.PutTodo)
	router.PATCH("/todo/:id",controllers.PatchTodo)
	router.DELETE("/todo/:id",controllers.DeleteTodo)

	router.Run()
}