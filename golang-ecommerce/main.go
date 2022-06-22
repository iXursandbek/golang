package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ixursandbek/golang-ecommerce/controllers"
	"github.com/ixursandbek/golang-ecommerce/database"
	"github.com/ixursandbek/golang-ecommerce/middleware"
	// "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":", port))
}
