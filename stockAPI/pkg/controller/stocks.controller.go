package controller

import (
	"net/http"
	"stocksapi/pkg/models"
	"stocksapi/pkg/services"
	"stocksapi/pkg/utils"
	"strconv"
)

type StockController interface {
	CreateStock(w http.ResponseWriter, r *http.Request)
	GetStocks(w http.ResponseWriter, r *http.Request)
	GetStockByID(w http.ResponseWriter, r *http.Request)
	UpdateStock(w http.ResponseWriter, r *http.Request)
	DeleteStock(w http.ResponseWriter, r *http.Request)
}

type stockController struct {
	service services.StocksService
}

func NewStocksController(service services.StocksService) StockController {
	return &stockController{service: service}
}

func (c *stockController) CreateStock(w http.ResponseWriter, r *http.Request) {
	stocks := &models.Stock{}
	utils.BodyParser(r, stocks)

	newStock, err := c.service.Create(stocks)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.Response(w, newStock, 201)
}

func (c *stockController) GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := c.service.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, stocks, 200)
}

func (c *stockController) GetStockByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stock, err := c.service.GetById(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, stock, 200)
}

func (c *stockController) UpdateStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stock := &models.Stock{}
	utils.BodyParser(r, stock)

	newStock, err := c.service.Update(int64(id), stock)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.Response(w, newStock, 200)
}

func (c *stockController) DeleteStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.service.Delete(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Stock is Deleted", 200)
}
