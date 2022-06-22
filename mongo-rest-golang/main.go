package main

import (
	"gihub.com/ixursandbek/mongo-rest-golang/controllers"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {
	r := gin.Default()
	uc := controllers.NewUserController(getsession())

	r.GET("/user/:id", uc.GetUser)
	// r.POST("/user/", uc.CreateUser)
	// r.DELETE("/user/:id", uc.DeleteUser)

	r.Run(":9090")

}

func getsession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s
}
