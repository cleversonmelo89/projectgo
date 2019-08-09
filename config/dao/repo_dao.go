package dao

import (
	. "../../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
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

func (r *RepoDao) GetAll() (TotalRepo, error) {
	var repos []Repo
	var total TotalRepo
	err := db.C(COLLECTION).Find(bson.M{}).All(&repos)
	total.Total = len(repos)
	total.Repo = repos
	return total, err
}

func (r *RepoDao) GetByIDGit(id string) (Repo, error) {
	var repo Repo
	err := db.C(COLLECTION).Find(buildQueryIdGit(id)).One(&repo)
	return repo, err
}

func (r *RepoDao) GetRepoByTag(tagsArray []bson.RegEx) ([]Repo, error) {
	var repo []Repo
	err := db.C(COLLECTION).Find(buildQueryByRepoTag(tagsArray)).All(&repo)
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
	err := db.C(COLLECTION).Update(buildQueryIdGit(id), &repo)
	return err
}

func buildQueryIdGit(id string) bson.M {
	idGit, _ := strconv.Atoi(id)
	query := bson.M{"id_git": idGit}
	return query
}

func buildQueryByRepoTag(tagsArray []bson.RegEx) bson.M {
	query := bson.M{"tag.tag_name": bson.M{"$in": tagsArray}}
	return query
}
