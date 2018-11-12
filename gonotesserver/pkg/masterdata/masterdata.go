// Package masterdata is the source of constant masterdata
package masterdata

const (
	AtlasConnectionString = "mongodb://dbUser:LetMe1n!@cluster0-shard-00-00-thawu.mongodb.net:27017,cluster0-shard-00-01-thawu.mongodb.net:27017,cluster0-shard-00-02-thawu.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true"

	DBName = "gonotes"

	CollectionNameCategories = "categories"
	CollectionNameTags       = "tags"
	CollectionNameNotes      = "notes"
)
