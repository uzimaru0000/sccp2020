package repo

import (
	"errors"

	"github.com/uzimaru0000/kapro/todo"
)

type TodoRepo interface {
	Get(*todo.Todo) (*todo.Todo, error)
	GetAll() ([]*todo.Todo, error)
	Store(*todo.Todo) error
	Update(*todo.Todo) error
	Delete(*todo.Todo) error
}

type inMemTodoRepo struct {
	table map[string]*todo.Todo
}

func NewInMemTodoRepo() *inMemTodoRepo {
	return &inMemTodoRepo{table: make(map[string]*todo.Todo)}
}

func (repo *inMemTodoRepo) Get(todo *todo.Todo) (*todo.Todo, error) {
	t, ok := repo.table[todo.Id]
	if !ok {
		return nil, errors.New("Please specify a valid ID")
	}

	return t, nil
}

func (repo *inMemTodoRepo) GetAll() ([]*todo.Todo, error) {
	todoList := []*todo.Todo{}

	for _, todo := range repo.table {
		todoList = append(todoList, todo)
	}

	return todoList, nil
}

func (repo *inMemTodoRepo) Store(todo *todo.Todo) error {
	_, ok := repo.table[todo.Id]
	if ok {
		return errors.New("Always store")
	}

	repo.table[todo.Id] = todo
	return nil
}

func (repo *inMemTodoRepo) Update(todo *todo.Todo) error {
	_, ok := repo.table[todo.Id]
	if !ok {
		return errors.New("Please specify a valid ID")
	}

	repo.table[todo.Id] = todo
	return nil
}

func (repo *inMemTodoRepo) Delete(todo *todo.Todo) error {
	_, ok := repo.table[todo.Id]
	if !ok {
		return errors.New("Please specify a valid ID")
	}

	delete(repo.table, todo.Id)
	return nil
}
