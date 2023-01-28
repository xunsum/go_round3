package models

type TodoItem struct {
	OwnerId    string
	Id         string
	Title      string
	Content    string
	CreateTime int64
	EndTime    int64
	Status     int //笔记状态，1 - 待办 0 - 已办
}

func (TodoItem) TableName() string {
	return "todos"
}
