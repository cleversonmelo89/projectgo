package reporouter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"../client"
	. "../config/dao"
	. "../models"
)

var dao = RepoDao{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	routes, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, routes)
}

func GetReposByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var url = "https://api.github.com/users/" + params["user"] + "/starred"

	resp, err := client.GetReposGitByUser(url)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//var repository []Repository
	var repo []Repo

	json.Unmarshal(resp.Body(), &repo)

	for index, _ := range repo {
		repo[index].BsonID = bson.NewObjectId()

		if err := dao.Create(repo[index]); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
	respondWithJson(w, http.StatusCreated, repo)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := dao.GetByIDGit(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}
	respondWithJson(w, http.StatusOK, repo)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var repo Repo
	if err := json.NewDecoder(r.Body).Decode(&repo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	repo.BsonID = bson.NewObjectId()

	if err := dao.Create(repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, repo)
}

func AddTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag := BuildTag(w, r)

	repo, err := dao.GetByIDGit(params["id_git"])

	for _, tagName := range tag {
		for _, repoTagName := range repo.Tag {
			if repoTagName.TagName == tagName.TagName {
				respondWithError(w, http.StatusConflict, "Tag Name: '"+tagName.TagName+"' duplicate")
				return
			}
		}
		repo.Tag = append(repo.Tag, tagName)
	}

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}

	if err := dao.Update(params["id_git"], repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": params["id_git"] + " atualizado com sucesso!"})
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag := BuildTag(w, r)

	repo, err := dao.GetByIDGit(params["id_git"])

	for _, tagName := range tag {
		for index, repoTagName := range repo.Tag {
			if repoTagName.TagName == tagName.TagName {
				repo.Tag = repo.Tag[:index+copy(repo.Tag[index:], repo.Tag[index+1:])]
			}
		}
	}

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}

	if err := dao.Update(params["id_git"], repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": params["id_git"] + " atualizado com sucesso!"})
}

func EditTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	repo, err := dao.GetByIDGit(params["id_git"])

	for index, repoTagName := range repo.Tag {
		if repoTagName.TagName == params["tag_name"] {
			repo.Tag[index].TagName = params["tag"]
		}
	}

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}

	if err := dao.Update(params["id_git"], repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": params["id_git"] + " atualizado com sucesso!"})
}

func BuildTag(w http.ResponseWriter, r *http.Request) []Tags {
	defer r.Body.Close()
	var tag []Tags
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return tag
	}
	return tag
}
