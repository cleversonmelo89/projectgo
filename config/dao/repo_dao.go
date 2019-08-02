package dao

import (
	"log"

	. "../../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RepoDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "repo"
)

func (r *RepoDao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

func (r *RepoDao) GetAll() ([]Repo, error) {
	var repos []Repo
	err := db.C(COLLECTION).Find(bson.M{}).All(&repos)
	return repos, err
}

func (r *RepoDao) GetByID(id string) (Repo, error) {
	var repo Repo
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&repo)
	return repo, err
}

func (r *RepoDao) Create(repo Repo) error {
	err := db.C(COLLECTION).Insert(&repo)
	return err
}

func (r *RepoDao) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (r *RepoDao) Update(id string, repo Repo) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &repo)
	return err
}
