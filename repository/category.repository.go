package repository

import (
	"context"
	"log"

	"github.com/go-bp-ikm-api/dto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// type CategoryRepository struct {
// 	c      *mongo.Collection
// 	ctx    context.Context
// 	logger log.Logger
// }

type CategoryRepository struct {
	c      *mongo.Collection
	ctx    context.Context
	logger log.Logger
}

func NewCategoryRepository(collection *mongo.Collection, ctx context.Context, logger log.Logger) CategoryRepository {
	return CategoryRepository{
		c:      collection,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *CategoryRepository) CreateCategory(category *dto.CreateCategoryDto) (*mongo.InsertOneResult, error) {
	id := uuid.New().String()

	category.Id = id
	result, err := r.c.InsertOne(r.ctx, category)
	return result, err
}

func (r *CategoryRepository) UpdateCategory(category *dto.FilterCategoryDto) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "name", Value: category.Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: category.Name}, primitive.E{Key: "name", Value: category.Status}}}}
	result, err := r.c.UpdateOne(r.ctx, filter, update)
	return result, err
}

func (r *CategoryRepository) DeleteCategory(id *string) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	result, err := r.c.DeleteOne(r.ctx, filter)
	return result, err
}

func (r *CategoryRepository) GetAll() (*mongo.Cursor, error) {
	cursor, err := r.c.Find(r.ctx, bson.D{{}})
	return cursor, err
}

func (r *CategoryRepository) FindFilter() (*mongo.Cursor, error) {
	cursor, err := r.c.Find(r.ctx, bson.D{{}})
	return cursor, err
}

func (r *CategoryRepository) FindById(id *string) *mongo.SingleResult {
	query := bson.D{bson.E{Key: "id", Value: id}}
	cursor := r.c.FindOne(r.ctx, query)
	return cursor
}
