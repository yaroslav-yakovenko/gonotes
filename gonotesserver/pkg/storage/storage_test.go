package storage

import (
	"fmt"
	"os"
	"testing"
)

var db *MongoDB
var err error

// unicode symbols
const check = "\u2714"
const cross = "\u2716"

func TestMain(m *testing.M) {

	db, err = NewMongoDBClient()
	if err != nil {
		fmt.Println(cross, "Cannot connect to DB")
	}

	os.Exit(m.Run())

}

func TestGetCategories(t *testing.T) {

	data, err := db.GetCategories()
	if err != nil {
		t.Error(cross, "Error: ", err)
	}

	if data == nil {
		t.Error(cross, "Result is nil")
	}

	t.Log(check, "Successful. Records found", len(data))

}

func TestGetTags(t *testing.T) {

	data, err := db.GetTags()
	if err != nil {
		t.Error(cross, "Error: ", err)
	}

	if data == nil {
		t.Error(cross, "Result is nil")
	}

	t.Log(check, "Successful. Records found", len(data))

}

func TestGetNotes(t *testing.T) {

	data, err := db.GetNotes()
	if err != nil {
		t.Error(cross, "Error: ", err)
	}

	if data == nil {
		t.Error(cross, "Result is nil")
	}

	t.Log(check, "Successful. Records found", len(data))

}
