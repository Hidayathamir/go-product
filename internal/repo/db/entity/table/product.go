package table

import "github.com/sirupsen/logrus"

// Product is table `product`. Use this to get table name and column name when query to database.
// Got panic? did you run Init which run initTableProduct?
var Product *product

type product struct {
	tableName  string
	Dot        *product
	Constraint productConstraint

	ID          string
	SKU         string
	Slug        string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
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
		tableName: "product",
		Dot:       &product{},
		Constraint: productConstraint{
			ProductPk:      "product_pk",
			ProductUnique:  "product_unique",
			ProductUnique1: "product_unique_1",
		},
		ID:          "id",
		SKU:         "sku",
		Slug:        "slug",
		Name:        "\"name\"",
		Description: "description",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
	}

	Product.Dot = &product{
		tableName: Product.tableName,
		Dot:       &product{},
		Constraint: productConstraint{
			ProductPk:      Product.Constraint.ProductPk,
			ProductUnique:  Product.Constraint.ProductUnique,
			ProductUnique1: Product.Constraint.ProductUnique1,
		},
		ID:          Product.tableName + "." + Product.ID,
		SKU:         Product.tableName + "." + Product.SKU,
		Slug:        Product.tableName + "." + Product.Slug,
		Name:        Product.tableName + "." + Product.Name,
		Description: Product.tableName + "." + Product.Description,
		CreatedAt:   Product.tableName + "." + Product.CreatedAt,
		UpdatedAt:   Product.tableName + "." + Product.UpdatedAt,
	}
}
