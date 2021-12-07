package local

import "github.com/wellingtonlope/todo-api/application/repositories"

func GetRepositories() repositories.AllRepositories {
	return repositories.AllRepositories{
		TodoRepository: &TodoRepositoryLocal{},
	}
}
