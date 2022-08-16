package services

import (
	"context"
	"log"

	"github.com/go-bp-ikm-api/dto"
	"github.com/go-bp-ikm-api/mapping"
	"github.com/go-bp-ikm-api/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryService struct {
	r   repository.CategoryRepository
	ctx context.Context
	c   *mongo.Collection
	log log.Logger
}

func NewCategoryService(ctx context.Context, collection *mongo.Collection, log log.Logger) ICategoryService {
	return &CategoryService{
		ctx: ctx,
		c:   collection,
		log: log,
	}
}

func (u *CategoryService) CreateCategory(categoryDto *dto.CreateCategoryDto) error {
	repo := repository.NewCategoryRepository(u.c, u.ctx, u.log)
	_, err := repo.CreateCategory(categoryDto)

	return err
}

func (u *CategoryService) GetAll() ([]*mapping.CategoryVm, error) {
	var categorys []*mapping.CategoryVm
	repo := repository.NewCategoryRepository(u.c, u.ctx, u.log)
	cursor, err := repo.GetAll()

	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		var category mapping.CategoryVm
		err := cursor.Decode(&category)
		if err != nil {
			return nil, err
		}

		categorys = append(categorys, &category)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)

	return categorys, nil
}

func (u *CategoryService) GetCategoryById(id *string) *mapping.CategoryVm {
	var category *mapping.CategoryVm
	repo := repository.NewCategoryRepository(u.c, u.ctx, u.log)
	err := repo.FindById(id).Decode(&category)
	if err != nil {
		return nil
	}
	return category
}
