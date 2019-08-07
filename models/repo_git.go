package models

type Repository struct {
	ID int64 `bson:"id" json:"id"`
	HtmlUrl string `bson:"html_url" json:"html_url"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Language string `bson:"language" json:"language"`
}