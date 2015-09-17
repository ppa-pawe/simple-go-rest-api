package main

import (  
    // Standard library packages
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
	"github.com/ppa-pawe/simple-go-rest-api/controllers"
)

func main() {  
    // Instantiate a new router
    r := httprouter.New()
	
	uc := controllers.NewUserController()

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