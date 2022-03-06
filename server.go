package main

import (
	"belajar-gin/controller"
	"belajar-gin/service"

	"github.com/gin-gonic/gin"
)

var (
	todoService    service.TodoService       = service.New()
	todoController controller.TodoController = controller.New(todoService)
)

func main() {
	r := gin.Default()
	r.GET("/todos", todoController.FindAll)
	r.POST("/todos", todoController.Save)
	r.Run(":5000")
}
