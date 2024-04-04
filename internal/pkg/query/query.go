// Package query contains func related to sql query string.
package query

import "fmt"

// JoinOnEqual return sql query "table on column1 = column2".
func JoinOnEqual(table, column1, column2 string) string {
	return fmt.Sprintf("%s ON %s = %s", table, column1, column2)
}

// Returning return sql query string RETURNING column.
func Returning(column string) string {
	return "RETURNING " + column
}
