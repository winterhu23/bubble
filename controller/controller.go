package controller

import (
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	url -----> controller----->logic -------> model
	请求来了 ---->控制器 --------> 业务逻辑----->模型层的增删改查
*/

//IndexHandler 处理首页的中间件
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//CreateTodo 创建事项中间件
func CreateTodo(c *gin.Context) {
	//前端页面填写待办事项，点击提交，会发请求到这里
	//1.获取数据
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	err := models.CreateTodo(&todo)
	//3.返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

//GetTodoList 查看所有待办事项中间件
func GetTodoList(c *gin.Context) {
	//查询数据库里的内容
	//if err = DB.Find(&todoList).Error;
	todoList, err := models.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}

	//显示
}

//UpdateTodo 更新事项中间件
func UpdateTodo(c *gin.Context) {
	//获取修改id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "获取ID失败",
		})
		return
	}
	//在数据库中找到对应的ID
	todo, err := models.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	//更新对应ID的修改后的值
	err = models.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

//DeleteTodo 删除todo
func DeleteTodo(c *gin.Context) {
	//获取要删除的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "获取ID失败",
		})
		return
	}
	//在数据库中找到对应的ID
	err := models.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})

}
