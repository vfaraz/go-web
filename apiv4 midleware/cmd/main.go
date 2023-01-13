package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"os"
	"rest/cmd/middlewares"
	"rest/cmd/routes"
	"rest/pkg/store"

	//"rest/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//var db []domain.Product

// api aplicando DDD
// @title Meli Bootcamp API
// @version 1.0
// @description This is a sample API
// @termsOfService http://swagger.io/terms/
// @contact.name Valentin Faraz
// @contact.url http://www.swagger.io/support
// @contact.email meli@meli
// @host localhost:8080
// @BasePath
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

	//engine por defaul con logger default
	//en := gin.Default()

	// engine con logger manual
	// La función gin.New() retorna un motor sin ningún middleware adicional, luego r.Use(gin.Recovery())
	en := gin.New()
	en.Use(gin.Recovery(), middlewares.LoggerMiddleware())

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
