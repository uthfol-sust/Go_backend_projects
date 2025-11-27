package repositories

import (
    "database/sql"
    "stocksapi/pkg/models"
)

type StocksRepository interface {
    Create(stock *models.Stock) (*models.Stock, error)
    GetAll() ([]models.Stock, error)
    GetById(id int64) (*models.Stock, error)
    Update(stock *models.Stock) (*models.Stock, error)
    Delete(id int64) error
}

type stocksRepository struct {
    db *sql.DB
}

func NewStocksRepository(db *sql.DB) StocksRepository {
    return &stocksRepository{db: db}
}

func (re *stocksRepository) Create(stock *models.Stock) (*models.Stock, error) {
    sqlStatement := `INSERT INTO stocks (name, price, company)
                     VALUES ($1, $2, $3)
                     RETURNING stockID`

    err := re.db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&stock.StockID)
    if err != nil {
        return nil, err
    }

    return stock, nil
}


func (re *stocksRepository) GetAll() ([]models.Stock, error) {
    sqlStatement := `SELECT stockid, name, price, company FROM stocks`

    rows, err := re.db.Query(sqlStatement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    stocks := []models.Stock{}
    for rows.Next() {
        var stck models.Stock
        err := rows.Scan(&stck.StockID, &stck.Name, &stck.Price, &stck.Company)
        if err != nil {
            return nil, err
        }
        stocks = append(stocks, stck)
    }
    return stocks, nil
}

func (re *stocksRepository) GetById(id int64) (*models.Stock, error) {
    sqlStatement := `SELECT stockid, name, price, company FROM stocks WHERE stockid=$1`
    newStock := &models.Stock{}

    err := re.db.QueryRow(sqlStatement, id).Scan(&newStock.StockID, &newStock.Name, &newStock.Price, &newStock.Company)
    if err != nil {
        return nil, err
    }
    return newStock, nil
}

func (re *stocksRepository) Update(stock *models.Stock) (*models.Stock, error) {
    sqlStatement := `UPDATE stocks SET name=$1, price=$2, company=$3 WHERE stockid=$4`

    _, err := re.db.Exec(sqlStatement, stock.Name, stock.Price, stock.Company, stock.StockID)
    if err != nil {
        return nil, err
    }
    return stock, nil
}

func (re *stocksRepository) Delete(id int64) error {
    sqlStatement := `DELETE FROM stocks WHERE stockid=$1`

    _, err := re.db.Exec(sqlStatement, id)
    return err
}
