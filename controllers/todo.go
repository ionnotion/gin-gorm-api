package controllers

import (
	"fmt"
	"gin-gorm-api/models"

	"github.com/gin-gonic/gin"
)

func GetTodos (ctx *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)

	ctx.JSON(200,gin.H{"data":todos})
}

func GetTodoById (ctx *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("ID= ?", ctx.Param("id")).First(&todo).Error; err != nil{
		ctx.JSON(404,gin.H{"error":"Data Not Found"})
		return
	}

	ctx.JSON(200,gin.H{"data":todo})
}

func PostTodo (ctx *gin.Context) {
	var input models.PostTodoForm
	if err := ctx.ShouldBindJSON(&input); err!= nil{
		ctx.JSON(400,gin.H{"error":err.Error()})
	}

	todo := models.Todo{Name: input.Name, Status: false, Description: input.Description, Priority: input.Priority}
	models.DB.Create(&todo)
	ctx.JSON(201,gin.H{"data":todo})
}

func PatchTodo (ctx *gin.Context) {
	var input models.PatchTodoForm
	if err := ctx.ShouldBindJSON(&input); err!= nil{
		ctx.JSON(400,gin.H{"error":err.Error()})
		return
	}

	var todo models.Todo
	if err := models.DB.Where("ID= ?", ctx.Param("id")).First(&todo).Error; err != nil{
		ctx.JSON(404,gin.H{"error":"Data Not Found"})
		return
	}

	models.DB.Model(&todo).Update("Status",input.Status)
	ctx.JSON(201,gin.H{"data":todo})
}

func PutTodo (ctx *gin.Context) {
	var input models.PostTodoForm
	if err := ctx.ShouldBindJSON(&input); err!= nil{
		ctx.JSON(400,gin.H{"error":err.Error()})
		return
	}

	var todo models.Todo
	if err := models.DB.Where("ID= ?", ctx.Param("id")).First(&todo).Error; err != nil{
		ctx.JSON(404,gin.H{"error":"Data Not Found"})
		return
	}

	models.DB.Model(&todo).Updates(models.Todo{Name: input.Name, Description: input.Description, Priority: input.Priority})
	ctx.JSON(201,gin.H{"data":todo})
}

func DeleteTodo (ctx *gin.Context) {
	var todo models.Todo
	todoId := ctx.Param("id")
	if err := models.DB.Where("ID= ?", todoId).First(&todo).Error; err != nil{
		ctx.JSON(404,gin.H{"error":"Data Not Found"})
		return
	}

	models.DB.Delete(&todo)
	message := fmt.Sprintf("Todo with ID %s deleted",todoId)
	ctx.JSON(200,gin.H{"data": message})
}