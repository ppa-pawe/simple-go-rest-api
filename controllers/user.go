package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/julienschmidt/httprouter"
    "github.com/ppa-pawe/simple-go-rest-api/models"
	"gopkg.in/mgo.v2"
)

type (
	UserController struct{
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	u := models.User{
		Name: "Bob Smith",
		Gender: "male",
		Age: 50,
		Id: p.ByName("id"),
	}
	
	uj, _ := json.Marshal(u)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	u := models.User{}
	
	json.NewDecoder(r.Body).Decode(&u)
	
	u.Id = "foo"
	uj, _ := json.Marshal(u)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	w.WriteHeader(200)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	w.WriteHeader(200)
}