package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gihub.com/ixursandbek/mongo-rest-golang/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		c.Writer.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		c.Writer.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// c.Header().Set("Content-type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	fmt.Fprintf("%s\n", uj)
}
func CreateUser() {}
func DeleteUser() {}
