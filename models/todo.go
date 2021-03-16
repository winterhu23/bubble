package models

import "bubble/dao"

//Todo 结构体 数据表
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//Todo 的增删改查

//CreateTodo 创建todo函数
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

//GetTodoList 查询所有todo函数
func GetTodoList() (todo []*Todo, err error) {
	todo = make([]*Todo, 0) //make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
	//而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
	if err = dao.DB.Find(&todo).Error; err != nil {
		return nil, err
	}
	return

}

//GetTodoByID 获取一个todo的函数
func GetTodoByID(ID string) (todo *Todo, err error) {
	todo = new(Todo) //new函数得到的是一个类型的指针
	if err = dao.DB.Where("id = ?", ID).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

//UpdateTodo 更新todo函数
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

//DeleteTodo 删除todo函数
func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}
