package services

import (
	"errors"
	"stocksapi/pkg/models"
	"stocksapi/pkg/repositories"
)

type StocksService interface {
	Create(stock *models.Stock) (*models.Stock, error)
	GetAll() ([]models.Stock, error)
	GetById(id int64) (*models.Stock, error)
	Update(id int64, stock *models.Stock) (*models.Stock, error)
	Delete(id int64) error
}

type stocksService struct{
	stockRepo repositories.StocksRepository
}


func NewStocksService(repo repositories.StocksRepository) StocksService {
	return &stocksService{stockRepo: repo}
}

func (s *stocksService) Create(stock *models.Stock) (*models.Stock, error){
 return s.stockRepo.Create(stock)
}

func (s *stocksService)	GetAll() ([]models.Stock, error){
 return s.stockRepo.GetAll()
}

func (s *stocksService)	GetById(id int64) (*models.Stock, error){
 return s.stockRepo.GetById(id)
}

func (s *stocksService)	Update(id int64, stock *models.Stock) (*models.Stock, error){
  updatedStocks, err := s.stockRepo.GetById(id)
  if err!=nil{
	return nil, errors.New("stock is not found to updated by id")
  }

  if stock.Name!=""{
	updatedStocks.Name = stock.Name
  }
  if stock.Price >=0 {
	updatedStocks.Price = stock.Price
  }
  if updatedStocks.Company!=""{
	updatedStocks.Company=stock.Company
  }

  return s.stockRepo.Update(updatedStocks)
}

func (s *stocksService)	Delete(id int64) error{
 _ , err := s.stockRepo.GetById(id)
  if err!=nil{
	return errors.New("stock is not found to delete by id")
  }

 return s.stockRepo.Delete(id)
}
