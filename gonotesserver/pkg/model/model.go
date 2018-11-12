// Package model implements data models structs and custom  types
package model

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Category - category of a note
type Category struct {
	ID          objectid.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string
	Description string
}

// Tag - description tag
type Tag struct {
	ID          objectid.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string
	Description string
}

// Note - code snippet, pattern, tutorial and so on
type Note struct {
	ID         objectid.ObjectID `bson:"_id,omitempty" json:"id"`
	Title      string
	Body       string
	CategoryID string
	Tags       []Tag
	Rating     int // 1 to 5 stars
}
