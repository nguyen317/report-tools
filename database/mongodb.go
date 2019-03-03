package database

import (
	"context"
	"time"

	"../modules"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func init() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://roger:roger123@ds213255.mlab.com:13255/report-tools"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	collection = client.Database("report-tools").Collection("cards")
	if err != nil {

	}
}

func InsertData(data interface{}, fn func(*mongo.InsertOneResult, error)) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		fn(nil, err)
	}
	fn(res, nil)
}

func FindOne(id string, fn func(interface{}, error)) {
	var card modules.MyCard
	filter := bson.M{"id": id}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&card)
	if err != nil {
		fn(nil, err)
	}
	fn(card, nil)
}
