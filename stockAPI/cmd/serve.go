package cmd

import (
	"fmt"
	"log"
	"net/http"
	"stocksapi/pkg/connection"
	"stocksapi/pkg/controller"
	"stocksapi/pkg/repositories"
	"stocksapi/pkg/router"
	"stocksapi/pkg/services"
	"time"
)

func Serve() {
	db, err := connection.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mux := http.NewServeMux()

	createRepo := repositories.NewStocksRepository(db)
	stockService := services.NewStocksService(createRepo)
	stockController := controller.NewStocksController(stockService)

	router.Router(mux, stockController)

	
	svr := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: time.Second *30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Minute,
	}

	

	fmt.Println("Server Running on Port 8080")
	log.Fatal(svr.ListenAndServe())
}
