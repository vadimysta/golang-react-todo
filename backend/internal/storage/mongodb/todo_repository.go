package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/vadimysta/golang-react-todo/backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	collections *mongo.Collection
}

func NewRepository(client *mongo.Client) *TodoRepository {

	return &TodoRepository{
		collections: client.Database("todoapp").Collection("todos"),
	}

}

func (r *TodoRepository) Create(ctx context.Context, todo *model.Todo) error {

	todo.ID = primitive.NewObjectID()
	_, err := r.collections.InsertOne(ctx, todo)

	return err

}

func (r *TodoRepository) Update(ctx context.Context, id string, completed bool) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to conversion to string: %v", err)
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = r.collections.UpdateOne(ctx, filter, update)

	return err

}

func (r *TodoRepository) Delete(ctx context.Context, id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to conversion to string: %v", err)
	}

	filter := bson.M{"_id": objectID}

	result, err := r.collections.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("todo not found")
	}

	return nil

}

func (r *TodoRepository) GetAll(ctx context.Context) ([]model.Todo, error) {

	cursor, err := r.collections.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get documents is MongoDB: %v", err)
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Printf("failed to close documents is MongoDB: %v", err)
			return
		}
	}()

	var todos []model.Todo

	if err := cursor.All(ctx, &todos); err != nil {
		return nil, fmt.Errorf("error in deserializing documents in slice: %v", err)
	}

	return todos, nil

}