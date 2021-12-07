package repositories

type AllRepositories struct {
	TodoRepository TodoRepository
}

type Repositories interface {
	GetRepositories() (*AllRepositories, error)
}
