package models

type TodoItemResponse struct {
	Data   TodoData `json:"data"`
	Error  string   `json:"error"`
	Msg    string   `json:"msg"`
	Status int      `json:"Status"` // 状态码
}

type TodoData struct {
	Item  []TodoItem `json:"item"`
	Total int        `json:"total"`
}

func (TodoItemResponse) TableName() string {
	return "todos"
}
