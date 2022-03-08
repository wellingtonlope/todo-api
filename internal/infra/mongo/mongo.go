package mongo

import (
	"context"

	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repositories struct {
	UriConnection string
	Database      string
}

func (r *Repositories) GetRepositories() (*repository.AllRepositories, error) {
	collections, err := getCollections(r.UriConnection, r.Database)
	if err != nil {
		return nil, err
	}

	return &repository.AllRepositories{
		TodoRepository: &TodoRepository{
			Collection: collections.Todo,
		},
	}, nil
}

type Collections struct {
	Todo *mongo.Collection
}

func getDatabase(uri, database string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client.Database(database), nil
}

func getCollections(uri, database string) (*Collections, error) {
	db, err := getDatabase(uri, database)
	if err != nil {
		return nil, err
	}

	return &Collections{
		Todo: db.Collection("todo"),
	}, nil
}
