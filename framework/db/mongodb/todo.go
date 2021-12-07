package mongodb

import (
	"context"
	"errors"

	"github.com/wellingtonlope/todo-api/application/myerrors"
	"github.com/wellingtonlope/todo-api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryMongo struct {
	Collection *mongo.Collection
}

func (r *TodoRepositoryMongo) GetAll() (*[]domain.Todo, *myerrors.Error) {
	todos := []domain.Todo{}
	cur, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, myerrors.NewError(err, myerrors.REPOSITORY)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		todo := domain.Todo{}
		cur.Decode(&todo)
		todos = append(todos, todo)
	}
	return &todos, nil
}

func (r *TodoRepositoryMongo) GetById(id string) (*domain.Todo, *myerrors.Error) {
	todo := domain.Todo{}

	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, myerrors.NewError(errors.New("todo not found"), myerrors.REGISTER_NOT_FOUND)
		}
		return nil, myerrors.NewError(err, myerrors.REPOSITORY)
	}

	return &todo, nil
}

func (r *TodoRepositoryMongo) Insert(todo *domain.Todo) (*domain.Todo, *myerrors.Error) {
	todoGet, _ := r.GetById(todo.ID)
	if todoGet != nil {
		return nil, myerrors.NewError(errors.New("todo is already exists"), myerrors.REGISTER_ALREADY_EXISTS)
	}

	_, err := r.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, myerrors.NewError(err, myerrors.REPOSITORY)
	}

	return todo, nil
}

func (r *TodoRepositoryMongo) Update(todo *domain.Todo) (*domain.Todo, *myerrors.Error) {
	_, myerr := r.GetById(todo.ID)
	if myerr != nil {
		return nil, myerr
	}

	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": todo.ID}, bson.M{"$set": todo})
	if err != nil {
		return nil, myerrors.NewError(err, myerrors.REPOSITORY)
	}

	return todo, nil
}

func (r *TodoRepositoryMongo) Delete(id string) *myerrors.Error {
	_, myerr := r.GetById(id)
	if myerr != nil {
		return myerr
	}

	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return myerrors.NewError(err, myerrors.REPOSITORY)
	}

	return nil
}
