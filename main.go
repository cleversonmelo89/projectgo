package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	_ "log"
	"net/http"
	_ "net/http"

	. "./config"
	. "./config/dao"
	reporouter "./router"
	_ "github.com/gorilla/mux"
)

var dao = RepoDao{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/repo", reporouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/repo/{user}", reporouter.GetReposByUser).Methods("GET")
	r.HandleFunc("/api/v1/repo/{id}", reporouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/repo", reporouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/repo/{id_git}/addTag", reporouter.AddTag).Methods("POST")
	r.HandleFunc("/api/v1/repo/{id_git}/deleteTag", reporouter.DeleteTag).Methods("DELETE")
	r.HandleFunc("/api/v1/repo/{id_git}/editTag/tag/{tag_name}/new/{tag}", reporouter.EditTag).Methods("PATCH")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
