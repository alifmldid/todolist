package todo

import (
	"context"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type TodoRepository interface{
	GetAll(c context.Context, activity_group_id string) (todos []Todo, err error)
	GetById(c context.Context, id int) (activity Todo, err error)
	Create(c context.Context, payload Todo) (todo Todo, err error)
	Update(c context.Context, id int, payload Todo) (activty Todo, err error)
	Delete(c context.Context, id int) (todo string, err error)
}

type todoRepository struct{
	Conn *gorm.DB
}

func NewTodoRepository(Conn *gorm.DB) TodoRepository{
	return &todoRepository{Conn}
}

func (repo *todoRepository) GetAll(c context.Context, activity_group_id string) (todos []Todo, err error){
	err = repo.Conn.Where("activity_group_id = ?", activity_group_id).Find(&todos).Error

	if (err != nil){
		return []Todo{}, err
	}
	
	return todos, nil
}

func (repo *todoRepository) GetById(c context.Context, id int) (todo Todo, err error){
	err = repo.Conn.Where("id = ?", id).First(&todo).Error

	if (err != nil){
		return Todo{}, err
	}

	return todo, err
}

func (repo *todoRepository) Create(c context.Context, payload Todo) (todo Todo, err error){
	todo = payload

	err = repo.Conn.Create(&todo).Error

	if err != nil{
		return Todo{}, err
	}

	return todo, nil
}

func (repo *todoRepository) Update(c context.Context, id int, payload Todo) (todo Todo, err error){
	err = repo.Conn.First(&todo, id).Error

	if err != nil{
		return Todo{}, err
	}

	todo.Title = payload.Title
	todo.Priority = payload.Priority
	todo.IsActive = payload.IsActive
	
	err = repo.Conn.Save(&todo).Error

	if err != nil{
		return Todo{}, err
	}

	return todo, nil
}

func (repo *todoRepository) Delete(c context.Context, id int) (status string, err error){
	var todo Todo
	err = repo.Conn.First(&todo, id).Error

	idString := strconv.Itoa(id)

	if err != nil{
		return "Not Found", errors.New("Todo with ID "+idString+" Not Found")
	}

	err = repo.Conn.Delete(&todo, id).Error

	if err != nil{
		return "", err
	}

	return "", nil
}
