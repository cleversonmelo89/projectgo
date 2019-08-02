package reporouter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

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

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := dao.GetByID(params["id"])
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
	repo.ID = bson.NewObjectId()
	if err := dao.Create(repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, repo)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var repo Repo
	if err := json.NewDecoder(r.Body).Decode(&repo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": repo.URL + " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
