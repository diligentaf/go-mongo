package main

import (
	"fmt"
	"go-mongo/db"
	"go-mongo/router"
	"go-mongo/store"
	"log"
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
	mongoClient, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	_ = mongoClient
	projectDb := db.SetupProjectDb(mongoClient)
	project := store.NewProjectStore(projectDb)
	_ = project
	fmt.Println(111)
}
