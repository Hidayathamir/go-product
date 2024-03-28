package repopostgres

import "github.com/Masterminds/squirrel"

var builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
