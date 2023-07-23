package main

import (
	"os"
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {


	port := os.Getenv("PORT")
	if port == "" {	port = "8080" }


	router := gin.New()
	router.Use(gin.Logger())