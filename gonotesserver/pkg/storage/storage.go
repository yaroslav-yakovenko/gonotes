// Package storage implemets MongoDB backend API
// https://godoc.org/github.com/mongodb/mongo-go-driver/mongo
package storage

import (
	"context"
	"errors"
	"gonotes/gonotesserver/pkg/masterdata"
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

// ErrDocumentExists - document exists!
var ErrDocumentExists = errors.New("Error. Document exists")

// NewMongoDBClient - constructor fabric
func NewMongoDBClient() (storage *MongoDB, err error) {

	var db MongoDB
	client, err := mongo.NewClient(masterdata.AtlasConnectionString)
	if err != nil {
		err = errors.New("Storage.NewMongoDBClient NewClient error: " + err.Error())
		return &db, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		err = errors.New("Storage.NewMongoDBClient Connect error: " + err.Error())
		return &db, err
	}

	database := client.Database(masterdata.DBName)

	db.Client = client
	db.Database = database

	return &db, err
}

//**************
// 	CATEGORIES
//**************

// GetCategories - fetch all categories
func (db *MongoDB) GetCategories() (categories []model.Category, err error) {

	collection := db.Database.Collection(masterdata.CollectionNameCategories)

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

// AddCategory - adds category to collection
func (db *MongoDB) AddCategory(document model.Category) (err error) {

	collection := db.Database.Collection(masterdata.CollectionNameCategories)

	var result bson.D
	var filter = bson.D{
		{
			Key:   "name",
			Value: document.Name,
		},
	}

	// bson.NewDocument(bson.EC.String("name", document.Name))

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

	return ErrDocumentExists

}

//**************
// 	TAGS
//**************

// GetTags - fetch all tags
func (db *MongoDB) GetTags() (tags []model.Tag, err error) {

	collection := db.Database.Collection(masterdata.CollectionNameTags)

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		err = errors.New("Storage.GetTags Find error: " + err.Error())
		return tags, err
	}

	for cursor.Next(context.Background()) {

		var tag model.Tag

		err = cursor.Decode(&tag)
		if err != nil {
			err = errors.New("Storage.GetTags Decode error: " + err.Error())
		}

		tags = append(tags, tag)

	}

	return tags, err

}

// AddTag - adds tag to collection
func (db *MongoDB) AddTag(document model.Tag) (err error) {

	collection := db.Database.Collection(masterdata.CollectionNameTags)

	var result bson.D
	var filter = bson.D{
		{
			Key:   "name",
			Value: document.Name,
		},
	}

	err = collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil && err != mongo.ErrNoDocuments {
		err = errors.New("Storage.AddTag FindOne error: " + err.Error())
		return err
	}

	if err == mongo.ErrNoDocuments {
		_, err = collection.InsertOne(context.Background(), document)
		if err != nil {
			err = errors.New("Storage.AddTag InserOne error: " + err.Error())
		}

		return err
	}

	return ErrDocumentExists

}

//**************
// 	NOTES
//**************

// GetNotes - fetch all notes
func (db *MongoDB) GetNotes() (notes []model.Note, err error) {

	collection := db.Database.Collection(masterdata.CollectionNameNotes)

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		err = errors.New("Storage.GetNotes Find error: " + err.Error())
		return notes, err
	}

	for cursor.Next(context.Background()) {

		var note model.Note

		err = cursor.Decode(&note)
		if err != nil {
			err = errors.New("Storage.GetNotes Decode error: " + err.Error())
		}

		notes = append(notes, note)

	}

	return notes, err

}

// AddNote - adds note to collection
func (db *MongoDB) AddNote(document model.Note) (err error) {

	collection := db.Database.Collection(masterdata.CollectionNameNotes)

	var result bson.D
	var filter = bson.D{
		{
			Key:   "title",
			Value: document.Title,
		},
	}

	err = collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil && err != mongo.ErrNoDocuments {
		err = errors.New("Storage.AddNote FindOne error: " + err.Error())
		return err
	}

	if err == mongo.ErrNoDocuments {
		_, err = collection.InsertOne(context.Background(), document)
		if err != nil {
			err = errors.New("Storage.AddNote InserOne error: " + err.Error())
		}

		return err
	}

	return ErrDocumentExists

}

// UpdateNote - updates note
func (db *MongoDB) UpdateNote(document model.Note) (err error) {

	collection := db.Database.Collection(masterdata.CollectionNameNotes)

	var filter = bson.D{
		{
			Key:   "_id",
			Value: document.ID,
		},
	}

	res := collection.FindOneAndReplace(context.Background(), filter, document, nil)
	if res == nil {
		err = errors.New("Storage.UpdateNote UpdateNote error: " + err.Error())
		return err
	}

	return err

}
