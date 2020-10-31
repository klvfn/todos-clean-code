package entity

// Todo is a struct contain single todo item
type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ModifyTodoReq is a struct contain request payload when modify or create todo
type ModifyTodoReq struct {
	Todo Todo `json:"todo"`
}
