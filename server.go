package main

import (  
    // Standard library packages
    "encoding/json"
    "fmt"
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "github.com/ppa-pawe/simple-go-rest-api/models"
)

func main() {  
    // Instantiate a new router
    r := httprouter.New()

    // Get a user resource
    r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        // Stub an example user
        u := models.User{
            Name:   "Bob Smith",
            Gender: "male",
            Age:    50,
            Id:     p.ByName("id"),
        }

        // Marshal provided interface into JSON structure
        uj, _ := json.Marshal(u)

        // Write content-type, statuscode, payload
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        fmt.Fprintf(w, "%s", uj)
    })
	
	r.POST("/user", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		u := models.User{}
		json.NewDecoder(r.Body).Decode(&u)
		u.Id = "foo"
		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", uj)
	})

    // Fire up the server
    http.ListenAndServe("localhost:3000", r)
}