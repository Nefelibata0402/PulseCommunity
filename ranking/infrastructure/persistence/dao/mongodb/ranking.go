package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"pulseCommunity/article/infrastructure/persistence/database/mongodb"
)

type TopMongoDB struct {
	col    *mongo.Collection
	pubCol *mongo.Collection
}

func NewArticleMongoDB() *TopMongoDB {
	return &TopMongoDB{
		col:    mongodb.New().Collection("articles"),
		pubCol: mongodb.New().Collection("published_articles"),
	}
}
