package todo

import (
	"context"
	"time"
)

type TodoUsecase interface{
	GetAllTodo(c context.Context, activity_group_id string) (todos []Todo, err error)
	GetTodoById(c context.Context, id int) (todo Todo, err error)
	CreateTodo(c context.Context, payload Payload) (todo Todo, err error)
	UpdateTodo(c context.Context, id int, payload Payload) (todo Todo, err error)
	DeleteTodo(c context.Context, id int) (status string, err error)
}

type todoUsecase struct{
	todoRepository TodoRepository
}

func NewTodoUsecase(todoRepository TodoRepository) TodoUsecase{
	return &todoUsecase{todoRepository}
}

func (uc *todoUsecase) GetAllTodo(c context.Context, activity_group_id string) (todos []Todo, err error){
	todos, err = uc.todoRepository.GetAll(c, activity_group_id)

	return todos, err
}

func (uc *todoUsecase) GetTodoById(c context.Context, id int) (todo Todo, err error){
	todo, err = uc.todoRepository.GetById(c, id)

	return todo, err
}

func (uc *todoUsecase) CreateTodo(c context.Context, payload Payload) (todo Todo, err error){
	todo.Title = payload.Title
	todo.ActivityGroupID = payload.ActivityGroupID
	todo.IsActive = payload.IsActive
	todo.Priority = "very-high"
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	todo, err = uc.todoRepository.Create(c, todo)

	return todo, err
}

func (uc *todoUsecase) UpdateTodo(c context.Context, id int, payload Payload) (todo Todo, err error){
	todo.Title = payload.Title
	todo.Priority = payload.Priority
	todo.UpdatedAt = time.Now()

	todo, err = uc.todoRepository.Update(c, id, todo)

	return todo, err	
}

func (uc *todoUsecase) DeleteTodo(c context.Context, id int) (status string, err error){
	status, err = uc.todoRepository.Delete(c, id)

	return status, err
}