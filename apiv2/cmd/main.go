package main

import (
	"encoding/json"
	"log"
	"os"
	"rest/cmd/routes"
	"rest/internal/domain"

	"github.com/gin-gonic/gin"
)

var db []domain.Product

// api con db en memoria
func main() {
	// instances
	//db := []domain.Product{
	//	{ID: 1, Name: "pollo", Quantity: 10, CodeValue: 1, IsPublished: true, Expiration: "10/10/2022", Price: 500},
	//	{ID: 2, Name: "manzana", Quantity: 10, CodeValue: 2, IsPublished: true, Expiration: "10/10/2022", Price: 500},
	//	{ID: 3, Name: "pera", Quantity: 10, CodeValue: 3, IsPublished: true, Expiration: "10/10-2022", Price: 500},
	//}

	err := ReadJson("./products.json")
	if err != nil {
		log.Fatal(err)
	}
	// server
	en := gin.Default()
	rt := routes.NewRouter(en, &db)
	rt.SetRoutes()

	// start
	if err := en.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}

func ReadJson(path string) (err error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return
	}
	json.Unmarshal(raw, &db)
	return
}
