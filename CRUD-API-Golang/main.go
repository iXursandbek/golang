package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := gin.Default()

	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "Titanik", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "434543", Title: "Uzuklar hukmdori", Director: &Director{Firstname: "Stive", Lastname: "Gerrard"}})

	r.GET("/movies", getMovies)
	r.GET("/movies/{id}", getMovie)
	r.POST("/movies", createMovie)
	r.PUT("/movies/{id}", updateMovie)
	// r.DELETE("/movies/{id}", deleteMovie)

	fmt.Printf("Server 8000 portida ishini boshladi")

	r.Run()
}

func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func createMovie(c *gin.Context) {
	var newMovie Movie
	if err := c.BindJSON(&newMovie); err != nil {
		return
	}
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

func getMovie(c *gin.Context) {
	id := c.Param("id")
	for _, v := range movies {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "Kino topilmadi")
}

func updateMovie(c *gin.Context) {
	id := c.Param("id")
	
}