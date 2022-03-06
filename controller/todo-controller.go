package controller

import (
	"belajar-gin/entity"
	"belajar-gin/service"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}
type controller struct {
	service service.TodoService
}

func New(service service.TodoService) TodoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	ctx.JSON(200, c.service.FindAll())
}
func (c *controller) Save(ctx *gin.Context) {
	var todo entity.Todo
	ctx.BindJSON(&todo)
	c.service.Save(todo)
	ctx.JSON(200, todo)
}
