package table

import "github.com/sirupsen/logrus"

// Stock is table `stock`. Use this to get table name and column name when query to database.
// Got panic? did you run Init which run initTableStock?
var Stock *stock

type stock struct {
	tableName string

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
}
