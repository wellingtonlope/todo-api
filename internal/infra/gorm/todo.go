package gorm

import (
	"context"

	"github.com/google/uuid"
	todoUC "github.com/wellingtonlope/todo-api/internal/app/usecase/todo"
	"github.com/wellingtonlope/todo-api/internal/domain"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(ctx context.Context, t domain.Todo) (domain.Todo, error) {
	t.ID = uuid.New().String()
	model := fromDomain(t)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return domain.Todo{}, err
	}
	return toDomain(model), nil
}

func (r *todoRepository) GetAll(ctx context.Context) ([]domain.Todo, error) {
	var models []TodoModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	todos := make([]domain.Todo, len(models))
	for i, m := range models {
		todos[i] = toDomain(m)
	}
	return todos, nil
}

func (r *todoRepository) GetByID(ctx context.Context, id string) (domain.Todo, error) {
	var model TodoModel
	if err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Todo{}, todoUC.ErrGetByIDStoreNotFound
		}
		return domain.Todo{}, err
	}
	return toDomain(model), nil
}

func (r *todoRepository) DeleteByID(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&TodoModel{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return todoUC.ErrDeleteByIDStoreNotFound
	}
	return nil
}

func (r *todoRepository) Update(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	model := fromDomain(todo)
	result := r.db.WithContext(ctx).Model(&model).Where("id = ?", todo.ID).Updates(&model)
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}
	if result.RowsAffected == 0 {
		return domain.Todo{}, todoUC.ErrUpdateStoreNotFound
	}
	return toDomain(model), nil
}
