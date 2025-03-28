package main

import (
	config2 "api-calculator/internal/config"
	"api-calculator/internal/handlers/calculate"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	config := config2.MustLoad()
	fmt.Print(config)

	r := gin.Default()
	r.POST("/calculate", calculate.CalculateHandler)
	err := r.Run(config.HTTPServer.Addr)

	if err != nil {
		log.Fatal("Server Error", err)
	}
}
