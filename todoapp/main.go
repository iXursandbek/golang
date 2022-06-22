package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "todo_db"
	collectionName string = "todo"
	port           string = "9090"
)

type (
	todoModel struct {
		ID        bson.ObjectId `json:"id,omitempty"`
		Title     string        `json:"title"`
		Completed bool          `json:"completed"`
		CreatedAt time.Time     `json:"created_at"`
	}
	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	r := gin.Default()

	r.GET("/", homeHandler)
	r.GET("/todo", todoHandlers())

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Porttni kutyabdi", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen:%s\n", err)
		}
	}()

	r.Run()

	<-stopChan
	log.Println("server to'xtayabdi...")
	ctx, cencel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cencel()
	log.Println("server muvaffaqqiyatli o'chdi")

}

func todoHandlers() http.Handler {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", fetchTodo)
		v1.POST("/", createTodo)
		v1.PUT("/{id}", updateTodo)
		v1.DELETE("/{id}", deleteTodo)
	}
	return r
}

// func homeHandler(c *gin.Context)  {
// 	// err := rnd.Template(http.StatusOK, []string{"static/home.tpl"}, nil)
// 	c.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
// 	checkErr(err)
// }

func createTodo(c *gin.Context)  {
	var t todo
	if err := json.NewDecoder(c.Request.Body).Decode(&t); err != nil {
		rnd.JSON(c, http.StatusProcessing, err)
	}
}

func fetchTodo(c *gin.Context)  {
	
}

func updateTodo(c *gin.Context)  {
	
}

func deleteTodo(c *gin.Context)  {
	
}