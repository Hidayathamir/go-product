package table

import "github.com/sirupsen/logrus"

// Stock is table `stock`. Use this to get table name and column name when query to database.
// Got panic? did you run Init which run initTableStock?
var Stock *stock

type stock struct {
	tableName string

	Dot *stock

	ID        string
	ProductID string
	Stock     string
	CreatedAt string
	UpdatedAt string

	Constraint stockConstraint
}

type stockConstraint struct {
	StockPk        string
	StockProductFk string
}

func (s *stock) String() string {
	return s.tableName
}

func initTableStock() {
	if Stock != nil {
		logrus.Warn("table Stock already initialized")
		return
	}

	Stock = &stock{
		tableName: "stock",
		ID:        "id",
		ProductID: "product_id",
		Stock:     "stock",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		Constraint: stockConstraint{
			StockPk:        "stock_pk",
			StockProductFk: "stock_product_fk",
		},
	}

	Stock.Dot = &stock{
		tableName: Stock.tableName,
		ID:        Stock.tableName + "." + Stock.ID,
		ProductID: Stock.tableName + "." + Stock.ProductID,
		Stock:     Stock.tableName + "." + Stock.Stock,
		CreatedAt: Stock.tableName + "." + Stock.CreatedAt,
		UpdatedAt: Stock.tableName + "." + Stock.UpdatedAt,
		Constraint: stockConstraint{
			StockPk:        Stock.Constraint.StockPk,
			StockProductFk: Stock.Constraint.StockProductFk,
		},
	}
}
