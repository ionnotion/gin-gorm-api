package models

type Todo struct {
	ID uint64 `gorm:"primaryKey; not null" json:"ID"`
	Name string `json:"name"`
	Status bool `json:"status"`
	Description string `json:"description"`
	Priority uint8 `json:"priority"`
}

type PostTodoForm struct {
	Name string `json:"name" binding:"required"`
	Status bool `json:"status"`
	Description string `json:"description" binding:"required"`
	Priority uint8 `json:"priority"`
}

type PatchTodoForm struct {
	Status bool `json:"status"`
}