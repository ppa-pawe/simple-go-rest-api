package main

import (  
    // Standard library packages
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
	"github.com/ppa-pawe/simple-go-rest-api/controllers"
	"gopkg.in/mgo.v2"
)

func main() {  
    // Instantiate a new router
    r := httprouter.New()
	
	uc := controllers.NewUserController(getSession())

    // Get a user resource
    r.GET("/user/:id", uc.GetUser)
	
	// Create a user resource
	r.POST("/user", uc.CreateUser)
	
	// Update a user resource
	r.PUT("/user/:id", uc.UpdateUser)
	
	// Delete a user resource
	r.DELETE("/user/:id", uc.DeleteUser)

    // Fire up the server
    http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb://rnd-win2012-1")
	
	if err != nil{
		panic(err)
	}
	
	return s
}