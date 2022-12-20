package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return TodoRepository{db}
}

func (u *TodoRepository) AddTodo(todo model.Todo) error {
	result := u.db.Create(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *TodoRepository) ReadTodo() ([]model.Todo, error) {
	result := []model.Todo{}
	resp := u.db.Raw("SELECT * FROM users WHERE deleted_at is null").Scan(&result)
	if resp.Error != nil {
		return []model.Todo{}, resp.Error
	}
	return result, nil // TODO: replace this
}

func (u *TodoRepository) UpdateDone(id uint, status bool) error {
	result := u.db.Model(&model.Todo{}).Where("id = ?", id).Update("done", status)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *TodoRepository) DeleteTodo(id uint) error {
	result := u.db.Where("id = ?", id).Delete(&model.Todo{})
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
