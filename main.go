package main

import (
	"log"

	"github.com/CadisRaziel/CRUD-GO/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Default x New -> New não estancia nenhum handler(middleware) o Default inicia com os handler(middleware) e loggers
	router := gin.Default()
	//& -> ponteiro para nós pegarmos o objeto original e não uma copia
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
