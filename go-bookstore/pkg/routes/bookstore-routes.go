package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ixursandbek/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func (c *gin.Context) {
	// router := gin.Default()
	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/", controllers.GetBook)
	router.GET("/book/{bookId}", controllers.GetBookByID)
	router.PUT("/book/{bookId}", controllers.UpdateBook)
	router.DELETE("/book/{bookId}", controllers.DeleteBook)

}

// islom universiteti 