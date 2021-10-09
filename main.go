package main

import (
	"net/http"

	"Appointy/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.GET("/posts/:id", uc.GetPost)
	r.POST("/user", uc.CreateUser)
	r.POST("/posts", uc.CreatePost)
	r.GET("/posts/:id/user/:id", uc.GetAllPost)
	http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
