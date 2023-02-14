// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/anuraaga/sqlc/internal/sql/ast"
	"github.com/anuraaga/sqlc/internal/sql/catalog"
)

var funcsPgBuffercache = []*catalog.Function{
	{
		Name:       "pg_buffercache_pages",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "record"},
	},
}

func PgBuffercache() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsPgBuffercache
	return s
}
