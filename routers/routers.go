package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

//SetRouters 设置路由
func SetRouters() *gin.Engine {
	//创建默认的路由引擎
	r := gin.Default()
	//加载静态资源
	r.Static("/static", "./static")
	//渲染HTML
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", controller.IndexHandler)
	//代办事项
	//路由组v1
	v1group := r.Group("v1")
	{
		//添加
		v1group.POST("/todo", controller.CreateTodo)
		//查看所有待办事项
		v1group.GET("/todo", controller.GetTodoList)
		//查看某个待办事项
		v1group.POST("/todo/:id", func(c *gin.Context) {

		})
		//修改某一个待办事项
		v1group.PUT("/todo/:id", controller.UpdateTodo)
		//删除某一个待办事项
		v1group.DELETE("/todo/:id", controller.DeleteTodo)

	}

	return r
}
