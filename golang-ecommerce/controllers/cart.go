package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ixursandbek/golang-ecommerce/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Panicln("maxsulot id si bo'sh")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("maxsulot id si bo'sh"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("foydalanuvchi id si bo'sh")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("foydalanuvchi id si bo'sh"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Savatchaga Muvaffaqqiyatli q'shildi")
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Panicln("maxsulot id si bo'sh")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("maxsulot id si bo'sh"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("foydalanuvchi id si bo'sh")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("foydalanuvchi id si bo'sh"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(200, "Savatchadan o'chirildi")
	}
}

func GetItemFromCart() gin.HandlerFunc {

}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" {
			log.Panicln("foydalanuvchi id si bo'sh")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID bo'sh"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "tartibga qo'shildi")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Panicln("maxsulot id si bo'sh")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("maxsulot id si bo'sh"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("foydalanuvchi id si bo'sh")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("foydalanuvchi id si bo'sh"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "tartibga qo'shildi")
	}
}
