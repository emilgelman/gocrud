package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/emilg02/gocrud/pkg/db"
	"github.com/emilg02/gocrud/pkg/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	db        *db.Db
	validator *validator.Validate
	router    *mux.Router
}

func NewHTTPServer(db *db.Db) *server {
	router := mux.NewRouter().StrictSlash(true)
	server := &server{db, validator.New(), router}
	router.HandleFunc("/articles", HandleWithLog(server.getAll)).Methods("GET")
	router.HandleFunc("/articles/{id}", HandleWithLog(server.get)).Methods("GET")
	router.HandleFunc("/articles/{id}", HandleWithLog(server.delete)).Methods("DELETE")
	router.HandleFunc("/articles", HandleWithLog(server.create)).Methods("POST")
	router.HandleFunc("/articles/{id}", HandleWithLog(server.update)).Methods("PUT")
	return server
}

func (s server) Serve() {
	log.Println("http listening on port 80")
	log.Fatal(http.ListenAndServe(":80", s.router))
}

func (s server) getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode((*s.db).GetAll())
}

func (s server) get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	res, err := (*s.db).Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s server) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := (*s.db).Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Article with id %s was deleted successfuly", id)
}

func (s server) create(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var article domain.Article
	if err := json.Unmarshal(body, &article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.validate(article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	(*s.db).Create(article.Id, article)
	json.NewEncoder(w).Encode(article)
}

func (s server) update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	body, _ := ioutil.ReadAll(r.Body)
	var article domain.Article
	json.Unmarshal(body, &article)
	if err := (*s.db).Update(id, article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Article with id %s was updated successfuly", id)
}

func (s server) validate(article domain.Article) error {
	if err := s.validator.Struct(article); err != nil {
		var result string
		for _, err := range err.(validator.ValidationErrors) {
			result += fmt.Sprintf("Field %s is %s\n", err.Field(), err.Tag())
		}
		return errors.New(result)
	}
	return nil
}
