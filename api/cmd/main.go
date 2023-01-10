package main

import (
	"log"
	"rest/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// server
	sv := gin.Default()

	// router and groups
	products := sv.Group("/products")
	products.POST("", handlers.CreateProduct)
	products.GET("", handlers.GetProducts)
	products.GET("/:id", handlers.GetProductById)


	// start
	if err := sv.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
