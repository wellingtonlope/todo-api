package memory

import "github.com/wellingtonlope/todo-api/internal/app/repository"

type Repositories struct{}

func (r *Repositories) GetRepositories() (*repository.AllRepositories, error) {
	return &repository.AllRepositories{
		TodoRepository: &TodoRepository{},
	}, nil
}
