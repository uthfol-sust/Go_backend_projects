package repositories

import (
	"database/sql"
	"stocksapi/pkg/models"
)

type StocksRepository interface {
	Create(stock *models.Stock) (*models.Stock, error)
	GetAll() ([]models.Stock, error)
	GetById(id int64) (* models.Stock, error)
	Update(stock *models.Stock) (*models.Stock ,error)
	Delete(id int64) error
}

type stocksRepository struct{
	db *sql.DB
}

func NewStocksRepository(db *sql.DB) StocksRepository{
	return &stocksRepository{db: db}
}

func (re *stocksRepository) Create(stock *models.Stock) (*models.Stock, error){

}
func (re *stocksRepository)	GetAll() ([]models.Stock, error){

}
func (re *stocksRepository)	GetById(id int64) (* models.Stock, error){

}
func (re *stocksRepository)	Update(stock *models.Stock) (*models.Stock ,error){

}
func (re *stocksRepository)	Delete(id int64) error{
	
}