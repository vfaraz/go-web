package main

import (
	"github.com/gin-gonic/gin"
)

type person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {
	router := gin.Default()
	router.POST("/saludo", func(c *gin.Context) {
		var persona person
		//Guardamos el json en la variable de tipo person persona
		c.BindJSON(&persona)
		response := "Hola" + " " + persona.Name + " " + persona.LastName
		//response as string
		//c.String(200, response)
		//response as json
		c.JSON(200, gin.H{"message": response})

	})

	err := router.Run(":8000")
	if err != nil {
		panic(err)

	}
}
