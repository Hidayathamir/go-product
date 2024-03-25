package table

import "github.com/sirupsen/logrus"

// Product is table `product`. Use this to get table name and column name when query to database.
// Got panic? did you run Init which run initTableProduct?
var Product *product

type product struct {
	tableName string

	ID          string
	SKU         string
	Slug        string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string

	Constraint productConstraint
}

type productConstraint struct {
	ProductPk      string
	ProductUnique  string
	ProductUnique1 string
}

func (p *product) String() string {
	return p.tableName
}

func initTableProduct() {
	if Product != nil {
		logrus.Warn("table Product already initialized")
		return
	}

	Product = &product{
		tableName:   "product",
		ID:          "id",
		SKU:         "sku",
		Slug:        "slug",
		Name:        "\"name\"",
		Description: "description",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
		Constraint: productConstraint{
			ProductPk:      "product_pk",
			ProductUnique:  "product_unique",
			ProductUnique1: "product_unique_1",
		},
	}
}
