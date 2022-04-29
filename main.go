package main

import (
	"fmt"
	"go-mongo/router"
	"os"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	fmt.Println(22)
}
