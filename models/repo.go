package models

import "gopkg.in/mgo.v2/bson"

type Repo struct {
	BsonID      bson.ObjectId `bson:"_id" json:"bson_id"`
	ID          int64         `bson:"id_git" json:"id"`
	HtmlUrl     string        `bson:"html_url" json:"html_url"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Language    string        `bson:"language" json:"language"`
	Tag         []Tags        `json:"tags"`
}

type Tags struct {
	TagName string `bson:"tag_name" json:"tag_name"`
}

type Suggestions struct {
	SuggestionName string `bson:"suggestion_name" json:"suggestion_name"`
}

type TotalRepo struct {
	Total int `bson: "total" json:"total"`
	Repo []Repo `bson: "repo" json:"repo"`
}