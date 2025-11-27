package router

import (
	"net/http"
	"stocksapi/pkg/controller"
)

func Router(router *http.ServeMux, controllers controller.StockController){
	router.Handle("POST /api/stocks", http.HandlerFunc(controllers.CreateStock))
	router.Handle("GET /api/stocks", http.HandlerFunc(controllers.GetStocks))
    router.Handle("GET /api/stocks/{id}", http.HandlerFunc(controllers.GetStockByID))
    router.Handle("PUT /api/stocks/{id}", http.HandlerFunc(controllers.UpdateStock))
    router.Handle("DELETE /api/stocks/{id}", http.HandlerFunc(controllers.DeleteStock))
}