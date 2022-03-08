package mongo

import (
	"context"
	"time"

	"github.com/wellingtonlope/todo-api/internal/app/repository"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedDate *time.Time         `bson:"created_date"`
	UpdatedDate *time.Time         `bson:"updated_date,omiempty"`
}

func idToObjectID(id string) (primitive.ObjectID, error) {
	if id == "" {
		return primitive.NewObjectID(), nil
	}
	return primitive.ObjectIDFromHex(id)
}

func fromDomainTodo(todo domain.Todo) (*Todo, error) {
	id, err := idToObjectID(todo.ID)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:          id,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedDate: todo.CreatedDate,
		UpdatedDate: todo.UpdatedDate,
	}, nil
}

func (t Todo) toDomainTodo() *domain.Todo {
	return &domain.Todo{
		ID:          t.ID.Hex(),
		Title:       t.Title,
		Description: t.Description,
		CreatedDate: t.CreatedDate,
		UpdatedDate: t.UpdatedDate,
	}
}

type TodoRepository struct {
	Collection *mongo.Collection
}

func (r *TodoRepository) GetAll() (*[]domain.Todo, error) {
	todos := []domain.Todo{}
	cur, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		todo := Todo{}
		cur.Decode(&todo)
		todos = append(todos, *todo.toDomainTodo())
	}
	return &todos, nil
}

func (r *TodoRepository) GetByID(id string) (*domain.Todo, error) {
	todo := Todo{}
	searchId, err := idToObjectID(id)
	if err != nil {
		return nil, repository.ErrTodoNotFound
	}

	err = r.Collection.FindOne(context.Background(), bson.M{"_id": searchId}).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrTodoNotFound
		}
		return nil, err
	}

	return todo.toDomainTodo(), nil
}

func (r *TodoRepository) Insert(todo domain.Todo) (*domain.Todo, error) {
	todoGet, _ := r.GetByID(todo.ID)
	if todoGet != nil {
		return nil, repository.ErrTodoAlreadyExists
	}

	todoInsert, err := fromDomainTodo(todo)
	if err != nil {
		return nil, err
	}

	insertOne, err := r.Collection.InsertOne(context.Background(), todoInsert)
	if err != nil {
		return nil, err
	}

	todoGet, err = r.GetByID(insertOne.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}

	return todoGet, nil
}

func (r *TodoRepository) Update(todo domain.Todo) (*domain.Todo, error) {
	_, err := r.GetByID(todo.ID)
	if err != nil {
		return nil, err
	}

	todoUpdate, err := fromDomainTodo(todo)
	if err != nil {
		return nil, err
	}

	_, err = r.Collection.UpdateOne(context.Background(), bson.M{"_id": todoUpdate.ID}, bson.M{"$set": todoUpdate})
	if err != nil {
		return nil, err
	}

	todoGet, err := r.GetByID(todo.ID)
	if err != nil {
		return nil, err
	}

	return todoGet, nil
}

func (r *TodoRepository) DeleteByID(id string) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}

	deleteId, err := idToObjectID(id)
	if err != nil {
		return repository.ErrTodoNotFound
	}

	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": deleteId})
	if err != nil {
		return err
	}

	return nil
}
