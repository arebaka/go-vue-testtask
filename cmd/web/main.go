package main

import (
	"fmt"
	"log"

	"are.moe/testtask/internal/config"
	"are.moe/testtask/internal/db"
	"are.moe/testtask/internal/handler"
	"are.moe/testtask/internal/repository"
	"are.moe/testtask/internal/router"
	"are.moe/testtask/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	// configure
	err := config.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	config := config.Get()
	dbInfo := config.DBInfo

	// init database
	db, err := db.Connect(dbInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// create handler for products
	repo := repository.NewProductRepo(db)
	productService := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(productService)

	// run server
	router := router.New(productHandler)
	if err = router.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		log.Fatal(err)
	}
}
