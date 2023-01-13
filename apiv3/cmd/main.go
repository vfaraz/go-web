package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"os"
	"rest/cmd/routes"
	"rest/pkg/store"

	//"rest/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//var db []domain.Product

// api aplicando DDD
func main() {
	// instances
	// db en memoria--->
	//db := []domain.Product{
	//	{ID: 1, Name: "pollo", Quantity: 10, CodeValue: 1, IsPublished: true, Expiration: "10/10/2022", Price: 500},
	//	{ID: 2, Name: "manzana", Quantity: 10, CodeValue: 2, IsPublished: true, Expiration: "10/10/2022", Price: 500},
	//	{ID: 3, Name: "pera", Quantity: 10, CodeValue: 3, IsPublished: true, Expiration: "10/10-2022", Price: 500},
	//}
	// cargando db en memoria desde un archivo--->
	//if err := ReadJson("./products.json"); err != nil {
	//	log.Fatal(err)
	//}
	//config env
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("TOKEN"))

	// server
	// implementando un paquete storage que controla la informacion en un json
	storage := store.NewStorage("./products.json")
	en := gin.Default()
	rt := routes.NewRouter(en, storage)
	rt.SetRoutes()

	// start
	if err := en.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}

//func ReadJson(path string) (err error) {
//	raw, err := os.ReadFile(path)
//	if err != nil {
//		return
//	}
//	json.Unmarshal(raw, &db)
//	return
//}
