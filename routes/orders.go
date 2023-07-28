package routes

import (
	"context"
	"log"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

//"server/models"
var validate = validator.New()

var orderCollection *mongo.Collection = OpenCollection(Client, "order")
//add an order 
func AddOrder(c *gin.Context) {
	var order models.Order
	c.BindJSON(&order)
	err := orderCollection.InsertOne(context.Background(), order)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"order":  order,
	})
}


