package models

import "gopkg.in/mgo.v2/bson"

type Repo struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	URL     string        `bson:"html_url" json:"html_url"`
	Starred bool          `bson:"starred" json:"starred"`
}
