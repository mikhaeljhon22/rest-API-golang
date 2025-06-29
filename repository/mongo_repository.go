package repository

import (
	"context"
	"restGolang/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type AboutRepository interface {
	Create(about *model.Mongos) error
}

type aboutRepositoryMongo struct {
	collection *mongo.Collection
}

func NewAboutRepositoryMongo(col *mongo.Collection) AboutRepository {
	return &aboutRepositoryMongo{collection: col}
}

func (r *aboutRepositoryMongo) Create(about *model.Mongos) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, about)
	return err
}
