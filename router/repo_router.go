package reporouter

import (
	"../client"
	. "../config/dao"
	. "../models"
	"strconv"
	"strings"

	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
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
	total, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, total)
}

func GetReposByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var url = "https://api.github.com/users/" + params["user"] + "/starred"

	resp, err := client.GetReposGitByUser(url)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var repo []Repo

	json.Unmarshal(resp.Body(), &repo)

	var getRepo Repo

	for index, repository := range repo {
		getRepo, err = dao.GetByIDGit(strconv.FormatInt(repository.ID, 10))

		if getRepo.ID == 0{
			repo[index].BsonID = bson.NewObjectId()

			if err := dao.Create(repo[index]); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
			}

			respondWithJson(w, http.StatusOK, repo)

			return
		}

		repo[index].BsonID = getRepo.BsonID
	}
	respondWithJson(w, http.StatusOK, repo)
}

func GetRepoByTag(w http.ResponseWriter, r *http.Request) {
	tags := BuildTag(w, r)

	var tagsArray []bson.RegEx
	var repo []Repo

	for _, tag := range tags {
		tagsArray = append(tagsArray, bson.RegEx{Pattern: ".*" + tag.TagName + ".*"})
	}

	repo, err := dao.GetRepoByTag(tagsArray)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}
	respondWithJson(w, http.StatusOK, repo)
}

func GetSuggestions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var repo Repo
	var suggestions []Suggestions

	repo, err := dao.GetByIDGit(params["id_git"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}

	if repo.Name != "" {
		for index, suggestion := range strings.Split(repo.Name, `-`) {
			var suggestionName Suggestions
			suggestionName.SuggestionName = suggestion
			suggestions = append(suggestions, suggestionName)
			if index == len(strings.Split(repo.Name, `-`))-1 {
				if repo.Language != "" {
					suggestionName.SuggestionName = repo.Language
					suggestions = append(suggestions, suggestionName)
				}
			}
		}
	}

	respondWithJson(w, http.StatusOK, suggestions)
}

func GetRepoByIdGit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo, err := dao.GetByIDGit(params["id_git"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}
	respondWithJson(w, http.StatusOK, repo)
}

func AddTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag := BuildTag(w, r)

	repo, err := dao.GetByIDGit(params["id_git"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}

	for _, tagName := range tag {
		for _, repoTagName := range repo.Tag {
			if repoTagName.TagName == tagName.TagName {
				respondWithError(w, http.StatusConflict, "Tag Name: '"+tagName.TagName+"' duplicate")
				return
			}
		}
		repo.Tag = append(repo.Tag, tagName)
	}

	if err := dao.Update(params["id_git"], repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, map[string]string{"result": "Tag Name adicionada com sucesso!"})
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag := BuildTag(w, r)

	repo, err := dao.GetByIDGit(params["id_git"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}

	for _, tagName := range tag {
		for index, repoTagName := range repo.Tag {
			if repoTagName.TagName == tagName.TagName {
				repo.Tag = repo.Tag[:index+copy(repo.Tag[index:], repo.Tag[index+1:])]
			}
		}
	}

	if err := dao.Update(params["id_git"], repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "Tag Name deletada com sucesso!"})
}

func EditTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	repo, err := dao.GetByIDGit(params["id_git"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Repo ID")
		return
	}

	for index, repoTagName := range repo.Tag {
		if repoTagName.TagName == params["tag_name"] {
			repo.Tag[index].TagName = params["tag"]
		}
	}

	if err := dao.Update(params["id_git"], repo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusAccepted, map[string]string{"result": " Tag Name atualizada com sucesso!"})
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
