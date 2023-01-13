package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

var AllProducts []Product

func main() {
	err := ReadJson("./products.json")
	if err != nil {
		panic(err)
	}

	sv := gin.Default()
	sv.GET("/ping", Ping)
	sv.GET("/products", GetProducts)
	sv.GET("/products/:id", GetProductById)
	sv.GET("/products/search", GetProductWithFilter)

	err = sv.Run(":8000")
	if err != nil {
		panic(err)
	}
}
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "succeed to get all products", "data": AllProducts})
}
func GetProductById(c *gin.Context) {
	// request
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error to get product", "data": nil})
		return
	}

	// process
	var response Product
	var flag bool
	for _, prod := range AllProducts {
		if prod.ID == id {
			response = prod
			flag = true
		}
	}
	// response
	if flag {
		c.JSON(http.StatusOK, gin.H{"message": "succeed to get product", "data": response})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "product not found", "data": nil})
	}
}
func GetProductWithFilter(c *gin.Context) {
	// request
	priceGT, err := strconv.ParseFloat(c.Query("priceGT"), 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "succeed to get products", "data": AllProducts})
		return
	}

	// process
	var response []Product
	var flag bool
	for _, prod := range AllProducts {
		if prod.Price > priceGT {
			response = append(response, prod)
			flag = true
		}
	}
	// response
	if flag {
		c.JSON(http.StatusOK, gin.H{"message": "succeed to get products", "data": response})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "products not found", "data": nil})
	}
}
func ReadJson(path string) (err error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return
	}
	json.Unmarshal(raw, &AllProducts)
	return
}
