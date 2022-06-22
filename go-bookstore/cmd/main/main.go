package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ixursandbek/go-bookstore/pkg/routes"
)


func main() {
	r := gin.Default()
	
	routes.RegisterBookStoreRoutes(r)

}