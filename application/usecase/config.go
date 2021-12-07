package usecase

import "github.com/wellingtonlope/todo-api/application/repositories"

type AllUseCases struct {
	TodoUseCase TodoUseCase
}

func GetUseCases(repositories repositories.Repositories) (*AllUseCases, error) {
	repos, err := repositories.GetRepositories()
	if err != nil {
		return nil, err
	}

	return &AllUseCases{
		TodoUseCase: TodoUseCase{
			TodoRepository: repos.TodoRepository,
		},
	}, nil
}
