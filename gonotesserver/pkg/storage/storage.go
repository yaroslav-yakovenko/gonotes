// Package storage implemets MongoDB backend API
package storage

import (
	"context"
	"errors"
	"gonotes/gonotesserver/pkg/model"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo" // MongoDB official driver
)

// MongoDB - MongoDB server representation
type MongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

var (
	err error
)

// NewMongoDBClient - constructor fabric
func NewMongoDBClient() (storage *MongoDB, err error) {

	var db MongoDB
	client, err := mongo.NewClient("mongodb://localhost")
	if err != nil {
		err = errors.New("Storage.NewMongoDBClient NewClient error: " + err.Error())
		return &db, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		err = errors.New("Storage.NewMongoDBClient Connect error: " + err.Error())
		return &db, err
	}

	database := client.Database("gonotes")

	db.Client = client
	db.Database = database

	return &db, err
}

//**************
// 	CATEGORIES
//**************

// GetCategories - fetch all categories
func (db *MongoDB) GetCategories() (categories []model.Category, err error) {

	collection := db.Database.Collection("categories")

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		err = errors.New("Storage.GetCategories Find error: " + err.Error())
		return categories, err
	}

	for cursor.Next(context.Background()) {

		var category model.Category

		err = cursor.Decode(&category)
		if err != nil {
			err = errors.New("Storage.GetCategories Decode error: " + err.Error())
		}

		categories = append(categories, category)

	}

	return categories, err

}

// AddCategory - adds category record
func (db *MongoDB) AddCategory(document model.Category) (err error) {

	collection := db.Database.Collection("categories")
	result := bson.NewDocument()
	filter := bson.NewDocument(bson.EC.String("name", document.Name))

	err = collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil && err != mongo.ErrNoDocuments {
		err = errors.New("Storage.AddCategory FindOne error: " + err.Error())
		return err
	}

	if err == mongo.ErrNoDocuments {
		_, err = collection.InsertOne(context.Background(), document)
		if err != nil {
			err = errors.New("Storage.AddCategory InserOne error: " + err.Error())
		}

		return err
	}

	err = errors.New("Storage.AddCategory error: Document exists")
	return err

}

//**************
// 	TAGS
//**************

//**************
// 	NOTES
//**************
