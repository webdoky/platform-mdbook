package main

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Record struct {
	ObjectID string `json:"objectID"`
	Slug     string `json:"slug"`
	Text     string `json:"text"`
	Title    string `json:"title"`
}

type Algolia struct {
	Client *search.Client
	Index  *search.Index
}

func NewAlgolia() *Algolia {
	client := search.NewClient(os.Getenv("ALGOLIA_APP_ID"), os.Getenv("ALGOLIA_ADMIN_KEY"))
	return &Algolia{
		Client: search.NewClient(os.Getenv("ALGOLIA_APP_ID"), os.Getenv("ALGOLIA_ADMIN_KEY")),
		Index:  client.InitIndex("articles"),
	}
}
