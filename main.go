package main

import (
	"github.com/gin-gonic/gin"
	"os" 
	"log"
)

func main() {
	route := gin.Default()

	StartRouter(route)

	port := os.Getenv("PORT")
		if port == "" {
			port = "3000" 
		}

	err := route.Run(":" + port)
		if err != nil {
			log.Fatal("Error starting server: ", err)
		}
}
