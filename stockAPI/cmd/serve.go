package cmd

import (
	"fmt"
	"stocksapi/pkg/connection"
	"stocksapi/pkg/repositories"
)

func Serve() {
    db, err := connection.Connection()
	if err != nil {
	fmt.Println(err.Error())
	  return
	}
    
	createRepo := repositories.NewStocksRepository(db)
}
