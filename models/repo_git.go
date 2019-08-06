package models

type Repository struct {
	HtmlUrl string `bson:"html_url" json:"html_url"`
}